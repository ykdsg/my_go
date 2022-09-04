package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello from goroutine")
		ch <- 0
	}()
	//这样能保证在主程序结束前goroutine 能够执行完成。
	i := <-ch
	fmt.Printf("main receive chan = %d", i)

}
