package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct {
}

func worker(i int) {
	fmt.Printf("worker %d: is working ...\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d:works done\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d:start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	//用于通知main goroutine 所有任务执行结束
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	//相当于一个控制阀，通过下面的close函数告诉goroutine 什么时候统一行动
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work..")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done!")
}
