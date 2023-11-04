package main

import (
	"fmt"
	"time"
)

type signal1 struct {
}

func worker1() {
	println("worker1 is working...")
	time.Sleep(time.Second)
}

//spawn 函数返回的 channel，被用于承载新 Goroutine 退出的“通知信 号”，这个信号专门用作通知 main goroutine。
//main goroutine 在调用 spawn 函数后 一直阻塞在对这个“通知信号”的接收动作上。
func spawn(f func()) <-chan signal1 {
	c := make(chan signal1)
	go func() {
		println("worker1 start to work...")
		f()
		c <- signal1{}
	}()
	return c
}

func main() {
	println("start a worker1...")
	c := spawn(worker1)
	<-c
	fmt.Println("worker1 work done!")
}
