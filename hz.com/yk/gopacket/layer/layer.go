package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
	"time"
)

var (
	//指定网络要捕获的网络设备名称，可以是FindAllDevs返回的设备的Name
	device       string = "ens33"
	snapshot_len int32  = 65536
	//是否将网口设置为混杂模式,即是否接收目的地址不为本机的包
	promiscuous bool = true
	err         error
	// timeout：设置抓到包返回的超时。如果设置成30s，那么每30s才会刷新一次数据包；设置成负数，会立刻刷新数据包，即不做等待
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle

	// Will reuse these for each packet
	ethLayer layers.Ethernet
	ipLayer  layers.IPv4
	tcpLayer layers.TCP
)

func main() {
	// 打开一个实时设备进行实时捕获（timeout为负）
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	//打开一个pcap文件
	//handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	//设置过滤器
	filter := "tcp and port 8000"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 8000 packets.")

	// 构造一个数据包源，NewPacketSource6只是构造一个PacketSource5对象
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	// 读取包，调用Packet函数时会启动一个协程来进行读取数据包，并将其写入到返回的管道中
	for packet := range packetSource.Packets() {
		// 打印包
		//fmt.Println(packet)
		fmt.Println("-------------------------------------------")
		printPacketInfo(packet)
		//quickDecodLayer(packet)
	}
}

//如果我们知道需要什么层，我们可以使用已有的结构来存储packet信息，
//而不是为每个packet创建新的结构,既浪费内存又浪费时间。使用DecodingLayerParser可以更快一点
func quickDecodLayer(packet gopacket.Packet) {
	parser := gopacket.NewDecodingLayerParser(
		layers.LayerTypeEthernet,
		&ethLayer,
		&ipLayer,
		&tcpLayer,
	)
	foundLayerTypes := []gopacket.LayerType{}

	err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
	if err != nil {
		fmt.Println("Trouble decoding layers: ", err)
	}

	for _, layerType := range foundLayerTypes {
		if layerType == layers.LayerTypeIPv4 {
			fmt.Println("IPv4: ", ipLayer.SrcIP, "->", ipLayer.DstIP)
		}
		if layerType == layers.LayerTypeTCP {
			fmt.Println("TCP Port: ", tcpLayer.SrcPort, "->", tcpLayer.DstPort)
			fmt.Println("TCP SYN:", tcpLayer.SYN, " | ACK:", tcpLayer.ACK)
		}
	}
}

func printPacketInfo(packet gopacket.Packet) {
	//看下是否是ethernet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)
		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
	}

	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)
		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}

	// Iterate over all layers, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}

	// When iterating through packet.Layers() above,
	// if it lists Payload layer then that is the same as
	// this applicationLayer. applicationLayer contains the payload
	//如果存在应用层，应用层存在Payload（有效载荷）
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("--------Application layer/Payload found.")
		fmt.Printf("%s\n", applicationLayer.Payload())
		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("--------HTTP found!")
		}
	}
	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}
