package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"crypto/tls"
	"fmt"
	"hash"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/Shopify/sarama"
	"github.com/xdg-go/scram"
)

func init() {
	go http.ListenAndServe(":6060", nil)
}

func main() {
	c, err := newClient([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	sp, err := sarama.NewSyncProducerFromClient(*c)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {

		kmsg := new(sarama.ProducerMessage)
		kmsg.Topic = "yestopic"
		kmsg.Key = sarama.ByteEncoder(nil)
		kmsg.Value = sarama.ByteEncoder([]byte("hello world"))

		p, o, e := sp.SendMessage(kmsg)
		fmt.Println(p, o, e)
		time.Sleep(time.Second * 5)
	}
}

func newClient(brokerAddrs []string, secret *Secret) (*sarama.Client, error) {
	saramaConf := newConfig()
	if secret != nil {
		if secret.TLS.Enable {
			// connect by tls
			saramaConf.Net.TLS.Enable = true
			saramaConf.Net.TLS.Config = secret.TLS.Config
		}
		if secret.SASL.Enable {
			// connect by SASL
			saramaConf.Net.SASL.Enable = true
			saramaConf.Net.SASL.User = secret.SASL.User
			saramaConf.Net.SASL.Password = secret.SASL.Password
			if secret.SASL.Mechanism == sarama.SASLTypeSCRAMSHA256 {
				saramaConf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
				saramaConf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
					return &XDGSCRAMClient{HashGeneratorFcn: SHA256}
				}
			} else {
				saramaConf.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
				saramaConf.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
					return &XDGSCRAMClient{HashGeneratorFcn: SHA512}
				}
			}
		}
	}

	client, err := sarama.NewClient(brokerAddrs, saramaConf)
	if err != nil {
		fmt.Println("err", saramaConf)
		return nil, err
	}

	return &client, nil
}

func newConfig() *sarama.Config {
	saramaConfig := sarama.NewConfig()
	//
	saramaConfig.Net.MaxOpenRequests = 1
	saramaConfig.Net.ReadTimeout = 120 * time.Second
	//
	saramaConfig.Producer.MaxMessageBytes = 524288000
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Idempotent = true
	//
	saramaConfig.Producer.Flush.Bytes = 40960000                 // batch大小 >= 1w*4k (xreq平均:176 aopz_result平均:3k)
	saramaConfig.Producer.Flush.Messages = 10000                 // msg数量 >= 1w
	saramaConfig.Producer.Flush.Frequency = 1 * time.Millisecond // 频次 >= 1ms
	saramaConfig.Producer.Flush.MaxMessages = 100000             // batch的最大msg数量 <= 10w (当到达这个限制后 会阻塞Input()<- 甚至导致死锁)
	saramaConfig.Producer.Return.Successes = true                // 同步生产模式下必须开启此参数, 如业务仓库中不需要则额外实现异步代码并忽略此配置
	saramaConfig.Producer.Partitioner = sarama.NewHashPartitioner

	saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	saramaConfig.ChannelBufferSize = 10240 // 10k: 有利于增加throughput
	saramaConfig.Version = sarama.V2_0_0_0

	return saramaConfig
}

type Secret struct {
	TLS struct {
		// Whether or not to use TLS when connecting to the broker
		// (defaults to false).
		Enable bool
		// The TLS configuration to use for secure connections if
		// enabled (defaults to nil).
		Config *tls.Config
	}
	SASL struct {
		// enable
		Enable bool
		// User is the authentication identity (authcid) to present for
		// SASL/PLAIN or SASL/SCRAM authentication
		User string
		// Password for SASL/PLAIN authentication
		Password string
		//default SASL will using SHA512 in this package
		Mechanism string
	}
}

var SHA256 scram.HashGeneratorFcn = func() hash.Hash { return sha256.New() }
var SHA512 scram.HashGeneratorFcn = func() hash.Hash { return sha512.New() }

type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	x.Client, err = x.HashGeneratorFcn.NewClient(userName, password, authzID)
	if err != nil {
		return err
	}
	x.ClientConversation = x.Client.NewConversation()
	return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	response, err = x.ClientConversation.Step(challenge)
	return
}

func (x *XDGSCRAMClient) Done() bool {
	return x.ClientConversation.Done()
}
