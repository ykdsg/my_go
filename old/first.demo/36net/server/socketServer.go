package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			println("accept error:", err)
			break
		}
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	time.Sleep(time.Second * 10)

	for {
		time.Sleep(5 * time.Second)
		var buf = make([]byte, 60000)
		log.Println("strat to read from conn")
		//设置超时时间，注意是绝对时间
		//err := c.SetReadDeadline(time.Now().Add(time.Second))
		//if err != nil {
		//	return
		//}
		n, err := c.Read(buf)
		if err != nil {
			fmt.Printf("conn read %d bytes,error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				//如果是超时进行其他业务处理
				continue
			}
		}
		fmt.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
