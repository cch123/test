package main


import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "time"
)

var (
    device       string = "lo0"
    snapshot_len int32  = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 30 * time.Second
    handle       *pcap.Handle
)

func main() {
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer handle.Close()

	packetSrc := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSrc.Packets() {
		// print packet
		fmt.Println(packet)
	}
}