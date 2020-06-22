package main

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, device := range devices {
		fmt.Println("name", device.Name)
		fmt.Println("description", device.Description)
		fmt.Println("addr", device.Addresses)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}
