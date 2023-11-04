package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	for {
		data, _ := Encode("123456789")
		_, err := conn.Write(data)
		data, _ = Encode("888888888")
		_, err = conn.Write(data)
		time.Sleep(time.Second * 1)
		data, _ = Encode("777777777")
		_, err = conn.Write(data)
		data, _ = Encode("123456789")
		_, err = conn.Write(data)
		time.Sleep(time.Second * 1)
		fmt.Println(err)
	}
}

//这里使用了Go的binary库，先写入4个字节的头，再写入消息主体
func Encode(message string) ([]byte, error) {
	// 读取消息的长度
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
