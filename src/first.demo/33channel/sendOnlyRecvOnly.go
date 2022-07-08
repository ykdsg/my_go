package main

import (
	"sync"
	"time"
)

//send only channel
func produce(ch chan<- int) {
	//for range 会阻塞在对 channel 的接收操作上，直到 channel 中有数据可接收或 channel 被关闭循环，才会继续向下执行。
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

//recv only channel
func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}
