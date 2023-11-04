package main

import (
	workerpool "com.yk/first/workpool"
	"fmt"
	"time"
)

func main() {
	p := workerpool.New(5, workerpool.WithPreAllocWorkers(false), workerpool.WithBlock(false))
	//time.Sleep(10 * time.Minute)
	time.Sleep(5 * time.Second)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(3 * time.Second)
		})
		if err != nil {
			fmt.Printf("task[%d]: error: %s\n", i, err.Error())
		}
	}
	p.Free()
}
