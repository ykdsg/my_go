package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"time"
)

func getLocalIP(dstip net.IP) (net.IP, int) {
	serverAddr, err := net.ResolveUDPAddr("udp", dstip.String()+":23456")
	if err != nil {
		log.Fatal(err)
	}
	if con, err := net.DialUDP("udp", nil, serverAddr); err == nil {
		if udpaddr, ok := con.LocalAddr().(*net.UDPAddr); ok {
			return udpaddr.IP, udpaddr.Port
		}
	}
	log.Fatal("could not get local ip: " + err.Error())
	return nil, -1
}

func main() {
	srcMac, _ := net.ParseMAC("a6:5e:60:cd:0b:65")
	distMac, _ := net.ParseMAC("00:0c:29:d5:22:98")

	srcIP := net.ParseIP("192.168.119.1")
	dstIP, _ := net.ResolveIPAddr("ip4", "192.168.119.31")

	_, srcPort := getLocalIP(dstIP.IP)

	// 链路层
	eth := layers.Ethernet{
		SrcMAC:       srcMac,  //发送端的 mac
		DstMAC:       distMac, //发送端的 mac
		EthernetType: layers.EthernetTypeIPv4,
	}

	//IP层
	ipLayer := layers.IPv4{
		SrcIP:    srcIP,
		DstIP:    dstIP.IP,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}

	// 四层 tcp
	tcpLayer := layers.TCP{
		SrcPort: layers.TCPPort(srcPort),
		DstPort: layers.TCPPort(8000),
	}

	data := []byte(`abc`)
	payload := gopacket.Payload(data)

	err := tcpLayer.SetNetworkLayerForChecksum(&ipLayer)
	if err != nil {
		panic(err)
	}
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err = gopacket.SerializeLayers(buf, opts, &eth, &ipLayer, &tcpLayer, payload)
	if err != nil {
		panic(err)
	}
	//这里网卡需要选对，否则流量并不能发送出去
	handle, err := pcap.OpenLive("bridge101", 2048, false, 30*time.Second)
	if err != nil {
		log.Fatal("pcap打开网络设备失败:", err)
	}
	defer handle.Close()
	//向 我们的网络设备发包
	err = handle.WritePacketData(buf.Bytes())
	if err != nil {
		log.Fatal("发送数据失败:", err)
	}
	log.Print("数据包已经发送\n")

}
