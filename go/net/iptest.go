package main

import "net"
import "fmt"

func main() {
	var host = "www.google.com"
		var ip = net.ParseIP(host)
		var ip2, _ = net.ResolveIPAddr("ip", host)
		fmt.Println(ip)
		fmt.Println(ip2.Zone)
		fmt.Println(ip2.IP)
		var ip3, _ = net.ResolveIPAddr("ip", "127.0.0.1")
		fmt.Println(ip3.IP)
}
