package main

import (
	workpool2 "com.yk/demo/workpool2"
	"time"
)

func main() {
	pool := workpool2.New(5)
	for i := 0; i < 10; i++ {
		err := pool.Schedule(func() {
			time.Sleep(3 * time.Second)
		})
		if err != nil {
			println("task:", i, "err:", err)
		}
	}

}
