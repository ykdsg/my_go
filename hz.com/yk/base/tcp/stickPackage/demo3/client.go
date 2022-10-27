package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	data := []byte("{@@@@@@@@@@@一个完整的包@@@@@@@@@@@@@}")
	conn, err := net.DialTimeout("tcp", "localhost:8888", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
