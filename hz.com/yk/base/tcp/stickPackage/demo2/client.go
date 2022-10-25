package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn) {
	//拼命的发，从server 打印出来的内容很容易看到后面的包越来越大，存在粘包的情况。
	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write([]byte(words))
	}
}

func main() {
	server := "127.0.0.1:8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")

	go sender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}
