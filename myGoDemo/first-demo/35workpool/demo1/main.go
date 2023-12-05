package main

import (
	workpool1 "com.yk/demo/workpool1"
	"time"
)

func main() {
	p := workpool1.New(5)
	for i := 0; i < 30; i++ {
		err := p.Schedule(func() {
			time.Sleep(3 * time.Second)
		})
		if err != nil {
			println("task:", i, "err:", err)
		}
	}
	p.Free()
}
