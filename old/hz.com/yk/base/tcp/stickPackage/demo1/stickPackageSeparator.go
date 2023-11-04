package main

import (
	"bufio"
	"fmt"
	"net"
)

//采用边界符解决粘包问题，存在的缺点：
//1.如果需要传输的消息包含分隔符，那就需要提前做转义处理。
//2.性能问题，如果消息体特别大，每次查找分隔符的位置的话肯定会有一点消耗。
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
		//使用bufio库里面的NewReader把conn包装一下，然后使用ReadSlice方法读取内容，该方法会一直读直到遇到分隔符
		reader := bufio.NewReader(conn)
		for {
			slice, err := reader.ReadSlice('\n')
			if err != nil {
				continue
			}
			fmt.Printf("%s", slice)
		}
	}

}
