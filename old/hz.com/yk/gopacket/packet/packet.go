package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
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
		fmt.Println(packet)
		fmt.Println("--------------------------------------------")
	}
}
