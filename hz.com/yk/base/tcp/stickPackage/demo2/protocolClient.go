package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func newSender(conn net.Conn) {
	for i := 0; i < 1000; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(Packet([]byte(words)))
	}
	fmt.Println("send over")
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
	go newSender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}
