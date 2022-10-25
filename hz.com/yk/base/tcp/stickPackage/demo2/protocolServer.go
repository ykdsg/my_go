package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	netListen, err := net.Listen("tcp", ":8888")
	CheckError2(err)

	defer netListen.Close()

	Log2("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log2(conn.RemoteAddr().String(), " tcp connect success")
		go newHandleConnection(conn)
	}
}

func newHandleConnection(conn net.Conn) {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log2(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		tmpBuffer = Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			Log2(string(data))
		}
	}
}

func Log2(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError2(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
