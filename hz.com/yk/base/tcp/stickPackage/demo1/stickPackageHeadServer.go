package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

//这个解决方案存在一定问题，可以参考demo2的实现，

//在消息的头部添加数据包长度，接收方根据消息长度进行接收
//这种解决方案就需要双方能够按照约定的方式来编码
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(conn)
		for {
			//前4个字节表示数据长度
			peek, err := reader.Peek(4)
			if err != nil {
				continue
			}
			buffer := bytes.NewBuffer(peek)
			//读取数据长度
			var length int32
			err = binary.Read(buffer, binary.BigEndian, &length)
			if err != nil {
				continue
			}
			//Buffered 返回缓存中未读取的数据的长度,如果缓存区的数据小于总长度，则意味着数据不完整
			// 这里continue 个人理解是有点问题的，应该找个临时缓存 存下来，再往下读。如果length 非常大，这样是处理不了的。
			if int32(reader.Buffered()) < length+4 {
				continue
			}
			//从缓存区读取大小为数据长度的数据
			data := make([]byte, length+4)
			_, err = reader.Read(data)
			if err != nil {
				continue
			}
			fmt.Printf("receive data: %s\n", data[4:])
		}
	}
}
