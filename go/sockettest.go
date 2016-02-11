package main

import (
	"fmt"
	"net"
)

func main() {
	//exit := make(chan bool)
	ip := net.ParseIP("127.0.0.1")
	port := 81
	tcp_addr := net.TCPAddr{ip, port, ""}
	udp_addr := net.UDPAddr{ip, port, ""}
	fmt.Println(tcp_addr)
	fmt.Println(udp_addr)
	fmt.Println(tcp_addr.Network())
	fmt.Println(udp_addr.Network())
}
