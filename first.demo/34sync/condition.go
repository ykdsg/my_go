package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

var ready bool

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			//使用轮询的方式等待信号量
			//轮询的方式开销大，轮询间隔设置的不同，条件检查的及时性也会受到 影响。
			for {
				mu.Lock()
				if !ready {
					mu.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mu.Unlock()
				fmt.Printf("worker %d: start to work...\n", i)
				f(i)
				wg.Done()
				return
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func spawnGroupChan(f func(i int), num int, groupSignal *sync.Cond) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			groupSignal.L.Lock()
			for !ready {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}
	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func condition() {

	fmt.Println("start a group of workers...")
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroupChan(worker, 5, groupSignal)
	time.Sleep(2 * time.Second) // 模拟ready前的准备工作

	fmt.Println("the group of workers start to work...")
	groupSignal.L.Lock()
	ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}

func main() {
	//roll()
	condition()
}

func roll() {
	fmt.Println("start a group of workers...")
	mu := &sync.Mutex{}
	c := spawnGroup(worker, 5, mu)
	time.Sleep(2 * time.Second) // 模拟ready前的准备工作
	fmt.Println("the group of workers start to work...")

	mu.Lock()
	ready = true
	mu.Unlock()
	<-c
	fmt.Println("the group of workers work done!")
}
