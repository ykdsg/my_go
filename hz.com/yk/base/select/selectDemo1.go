package main

import (
	"fmt"
	"time"
)

//能够看到ch写入的部分数字不见了
//当tick.C和ch同时满足读写条件时，select随机选择了一个执行
//这个例子是比较极端的，因为向ch写入的数据本身就与外部for循环计数耦合了，导致依赖于select的随机结果(本次没随机到，
//放到下次，但此时写入的数据已经变更了)，因此实际不是数据丢了，而是代码设计时没有考虑到每次select只会执行一条读写语句(并且是随机选取的)，导致结果不如预期。
func main() {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val:%d\n", val)
		}

	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case <-tick.C:
			fmt.Printf("%d:case<-tick.C\n", i)
		}
		time.Sleep(500 * time.Millisecond)
	}

	close(ch)
	tick.Stop()

}
