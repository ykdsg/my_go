package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"os"
	"time"
)

var (
	device      string = "ens33"
	snaplen     int32  = 65536
	promiscuous bool   = true
	err         error
	timeout     time.Duration = -1 * time.Second
	handle      *pcap.Handle
	packetCount = 0
)

//保存到文件
func main() {
	f, _ := os.Create("test.pcap")
	w := pcapgo.NewWriter(f)
	// 写入文件头，必须在调用前调用
	w.WriteFileHeader(uint32(snaplen), layers.LinkTypeEthernet)
	defer f.Close()

	handle, err := pcap.OpenLive(device, snaplen, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", device, err)
		os.Exit(1)
	}
	defer handle.Close()

	// source：源，需要实现ReadPacketData接口
	// decorder：解码器需要实现Decodeer接口
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++
		if packetCount > 500 {
			break

		}
	}

}
