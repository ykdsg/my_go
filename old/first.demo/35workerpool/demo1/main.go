package main

import (
	"com.yk/first/workpool"
	"time"
)

func main() {
	p := workerpool.New(5)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(3 * time.Second)
		})
		if err != nil {
			println("task:", i, "err:", err)
		}
	}

	p.Free()

}
