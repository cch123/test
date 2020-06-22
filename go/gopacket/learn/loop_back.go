/*
If you look at this diagram, you'll notice tcpdump acts on the Ethernet layer. Then comes, IP then TCP/UDP, then Sockets. nc operates at TCP/UDP layer.

In the IP level, the packets might be getting dropped. Very often the case is Reverse Path Filtering.

So, you are able to see the packets arriving at the ethernet layer which can be seen by tcpdump, but packets are not arriving for nc because they might be routed somewhere else, or dropped.

So, it's better to check if disabling RP filtering and checking iptable rules, helps!

Update:

As you're operating on the loopback interface:

MAC's are used at the lowest level of ethernet traffic, and only within one LAN and help direct traffic around within it. It isn't needed on a local network interface (lo) because packets are being handled internally.

The loopback address connects to the same computer directly in the IP layer without using any physical hardware. So, it lets you bypass Ethernet, PPP, other drivers.
*/

package main

import (
    "fmt"
    "net"

    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)

func main() {
    handle, err := pcap.OpenLive("lo0", 1500, false, pcap.BlockForever)
    if err != nil {
        fmt.Printf("%s\n", err.Error())
        return
    }

    eth := layers.Ethernet{
        EthernetType: layers.EthernetTypeIPv4,
        SrcMAC:       net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
        DstMAC:       net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
    }

    _ = eth // Ignore. Use where ethernet interface is used

    // Used for loopback interface
    lo := layers.Loopback{
        Family: layers.ProtocolFamilyIPv4,
    }

    ip := layers.IPv4{
        Version:  4,
        TTL:      64,
        SrcIP:    net.IP{127, 0, 0, 1},
        DstIP:    net.IP{127, 0, 0, 1},
        Protocol: layers.IPProtocolUDP,
    }

    udp := layers.UDP{
        SrcPort: 62003,
        DstPort: 9000,
    }
    udp.SetNetworkLayerForChecksum(&ip)

    payload := []byte{}

    options := gopacket.SerializeOptions{
        ComputeChecksums: true,
        FixLengths:       true,
    }

    buffer := gopacket.NewSerializeBuffer()

    // if err = gopacket.SerializeLayers(buffer, options,
    //  &eth,
    //  &ip,
    //  &udp,
    //  gopacket.Payload(payload),
    // ); err != nil {
    //  fmt.Printf("[-] Serialize error: %s\n", err.Error())
    //  return
    // }
    if err = gopacket.SerializeLayers(buffer, options,
        &lo,
        &ip,
        &udp,
        gopacket.Payload(payload),
    ); err != nil {
        fmt.Printf("[-] Serialize error: %s\n", err.Error())
        return
    }
    outgoingPacket := buffer.Bytes()

    if err = handle.WritePacketData(outgoingPacket); err != nil {
        fmt.Printf("[-] Error while sending: %s\n", err.Error())
        return
    }

}
