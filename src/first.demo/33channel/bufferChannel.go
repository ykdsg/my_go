package main

import (
	"log"
	"sync"
	"time"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs)
		println("close jobs-------------------------")
	}()

	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			// active 缓冲大小是3，因此这里同时活跃的最大是3个
			active <- struct{}{}
			log.Printf("handle job:%d\n", j)
			time.Sleep(2 * time.Second)
			//释放1个缓冲区
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}
