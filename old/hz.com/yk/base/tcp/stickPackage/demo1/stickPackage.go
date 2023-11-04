package main

import (
	"log"
	"net"
	"strings"
)

//通过telnet 127.0.0.1 8888 ，来发送相应的数据，如果超过10个字符，从日志中可以看到被截断
func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		for {
			//在接受消息的时候我们每次读取10个字节的内容输出并返回，如果输入的消息小于等于8（减去换行符）个
			//字符的时候没有问题，但是当我们在telnet里面输入大于10个字符的内容的时候，这些数据的时候会被强行拆开处理。
			data := make([]byte, 10)
			_, err := conn.Read(data)
			if err != nil {
				log.Printf("%s\n", err.Error())
				break
			}
			receive := string(data)
			log.Printf("receive msg: %s\n", receive)
			send := []byte(strings.ToUpper(receive))
			_, err = conn.Write(send)
			if err != nil {
				log.Printf("send msg failed, error: %s\n", err.Error())
			}

			log.Printf("send msg: %s\n", receive)
		}
	}

}
