package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

type EchoServer struct {
}

func (e *EchoServer) Echo(ctx context.Context, req *EchoReq) (*EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())

	res := &EchoRes{
		Msg: "success",
	}

	return res, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		panic(err)
	}

	handler := &EchoServer{}
	processor := NewEchoProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
