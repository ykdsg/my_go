package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	i := 0
	for {
		var err error
		_, err = conn.Write([]byte(strconv.Itoa(i) + " => 77777\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + " => 88888\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + " => 555555555555555555555555555555555555555555\n"))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
		//如果报文中包含分隔符，就需要做转义处理。
		_, err = conn.Write([]byte(strconv.Itoa(i) + " => 123456\n6666"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + " => 123456\n"))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
		_, err = conn.Write([]byte(strconv.Itoa(i) + " => 9999999\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + " => 0000000000000000000000000000000000000000000\n"))
		if err != nil {
			panic(err)
		}
		i++
	}
}
