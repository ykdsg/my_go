package main

import (
	"fmt"
	"runtime"
	"time"
)

func deadloop() {
	for {

	}
}

func main() {
	//go1.13的话加上runtime.GOMAXPROCS(1) ，在执行deadloop 之后main goroutine就无法继续。
	runtime.GOMAXPROCS(1)
	//在1.14 之后，go 支持了非协作的抢占式调度，所以main goroutine 还能继续调度
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
