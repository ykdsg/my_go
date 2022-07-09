package main

import (
	"fmt"
	"sync"
	"time"
)

//使用互斥锁保护的全局变量作为计数器
type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	//使用了一个带有互斥锁保护的全局变量作为计数器
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func lockCounter() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine-%d:current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

type counterNew struct {
	c chan int
	i int
}

//将计数器操作全部交给一个独立的 Goroutine 去处理，并通过无缓冲 channel 的同步阻塞特性，实现了计数器的控制。
func NewCounter() *counterNew {
	cter := &counterNew{
		c: make(chan int),
	}

	go func() {
		for {
			cter.i++
			//因为是无缓冲channel，在send之后，没有receive的时候，下一次调用send的时候会阻塞
			cter.c <- cter.i
		}
	}()

	return cter
}

//通过无缓冲 channel 的接收动作，相当于释放了NewCounter 中无缓冲channel send 的阻塞。
func (cter *counterNew) IncreaseNew() int {
	return <-cter.c
}

func channelCount() {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.IncreaseNew()
			fmt.Printf("goroutine-%d:current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	//传统的通过锁的方式，是通过共享内存来通信
	lockCounter()
	time.Sleep(2 * time.Second)

	println("--------------------next count------------------------")
	//通过channel的方式，是通过通信来共享内存
	channelCount()
}
