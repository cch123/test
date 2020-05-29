package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9898"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, err := transportFactory.GetTransport(transport)
	client := NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9898", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	req := &EchoReq{Msg: "You are welcome."}
	res, err := client.Echo(context.TODO(), req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}

	log.Println("response:", res.Msg)
	fmt.Println("well done")

}
