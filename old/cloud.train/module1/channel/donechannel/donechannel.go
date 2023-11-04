package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	fmt.Println("sleep......")
	//这里sleep 的时间有讲究，如果sleep 超过了consumer 处理完message 的时间，就会导致consumer 阻塞在<-messages，就不能响应<-done
	//所以缩小上面i的上限，或者增加sleep 的时间，就看不到interrupt 的打印
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
