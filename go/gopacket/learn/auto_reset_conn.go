package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
    "log"
    "time"
)

var (
    device       string = "en0"
    snapshot_len int32  = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 30 * time.Second
    handle       *pcap.Handle
)

func main() {
    // Open device
    handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    // Set filter
    var filter string = "tcp and port 8088 and tcp[tcpflags] & (tcp-syn) != 0 and tcp[tcpflags] & (tcp-ack) != 0"
    err = handle.SetBPFFilter(filter)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Only capturing TCP port 80 packets.")

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        // Do something with a packet here.
        fmt.Println(packet)
        handlePacket(handle, packet)
    }
}

func handlePacket(handle *pcap.Handle, packet gopacket.Packet) {
    ethLayer := packet.LinkLayer()
    if ethLayer == nil {
        println("oh noz")
        return
    }
    eth1, ok := ethLayer.(*layers.Ethernet)
    if !ok {
        println("oh nox")
        return
    }

    ipLayer := packet.NetworkLayer()
    if ipLayer == nil {
        println("oh no0")
        return
    }

    ip, ok := ipLayer.(*layers.IPv4)
    if !ok {
        println("oh no1")
        return
    }

    println("yes")

    buf := gopacket.NewSerializeBuffer()
    opts := gopacket.SerializeOptions{
        FixLengths:       true,
        ComputeChecksums: true,
    }

    eth := layers.Ethernet{
        SrcMAC:       eth1.SrcMAC,
        DstMAC:       eth1.DstMAC,
        EthernetType: layers.EthernetTypeIPv4,
    }

    ip4 := layers.IPv4{
        SrcIP:    ip.SrcIP,
        DstIP:    ip.DstIP,
        Id:       ip.Id,
        Flags:    layers.IPv4DontFragment,
        Version:  4,
        TTL:      ip.TTL,
        Protocol: layers.IPProtocolTCP,
    }

    tcpLayer := packet.Layer(layers.LayerTypeTCP)
    if tcpLayer == nil {
        println("oh no2")
        return
    }
    tcp, ok := tcpLayer.(*layers.TCP)
    if !ok {
        println("oh no3")
        return
    }

    tcp1 := layers.TCP{
        SrcPort: tcp.SrcPort,
        DstPort: tcp.DstPort,
        Seq:     tcp.Seq + 1,
        RST: true,
        Window:  tcp.Window,
       // PSH:     true, // 立刻处理
    }

    tcp1.SetNetworkLayerForChecksum(&ip4)
    payload := gopacket.Payload([]byte{})

    if err := gopacket.SerializeLayers(buf, opts, &eth, &ip4, &tcp1, payload); err != nil {
        //return err
        fmt.Println(err)
        return
    }

    err = handle.WritePacketData(buf.Bytes())
    fmt.Println(ip4.DstIP, ip4.SrcIP)
    fmt.Println(tcp1.DstPort, tcp1.SrcPort)
    fmt.Println(tcp1.RST)
    fmt.Println(err)
}
