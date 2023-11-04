package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(c chan<- int) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := trySend(c, i)
		if ok {
			fmt.Printf("[producer]:send [%d] to channel \n", i)
			i++
			continue
		}
		fmt.Printf("[producer]:try send [%d],but channel is full\n", i)
	}
}

//由于用到了 select 原语的 default 分支语义，当 channel 空的时候，tryRecv 不会阻塞
func tryRecv(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true

	default:
		return 0, false

	}
}

//当 channel 满的时候，trySend 也不会阻塞。
func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := tryRecv(c)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel is empty")
			time.Sleep(time.Second)
			continue
		}
		fmt.Printf("[consumer]: recv [%d] from channel\n", i)
		if i >= 3 {
			fmt.Println("[consumer]: exit")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()
	wg.Wait()
}
