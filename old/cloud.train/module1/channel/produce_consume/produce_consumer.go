package main

import (
	"fmt"
	"time"
)

/*
队列长度10，队列元素类型为 int
• 生产者：
每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
• 消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
func main() {
	ch := make(chan int, 10)
	defer close(ch)
	done := make(chan int)
	//consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case _, ok := <-done:
				fmt.Printf("consumer receive done,ok=%t...\n", ok)
				return
			//这样可以工作，但是不会阻塞for循环，如果要阻塞for循环，那么done就不能及时响应
			case receive := <-ch:
				fmt.Printf("consumer receive message :%d \n", receive)
			default:
				fmt.Println("consumer time ticker")
				//fmt.Printf("receive message :%d \n", <-ch)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	//produce
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for _ = range ticker.C {
		fmt.Printf("produce i=%d \n", i)
		ch <- i
		i++
		if i > 2 {
			break

		}
	}

	//sleep 之后，consumer会被阻塞在<-ch
	time.Sleep(5 * time.Second)
	//done <- 1
	close(done)
	time.Sleep(5 * time.Second)
	fmt.Println("main process done")
}
