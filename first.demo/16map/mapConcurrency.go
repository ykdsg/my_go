package main

import (
	"fmt"
	"time"
)

func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d,%d", k, v)
	}
}

func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}

func main() {
	m := map[int]int{
		1: 11, 2: 12, 3: 13,
	}

	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()

	time.Sleep(5 * time.Second)
	//	运行之后会出现：fatal error: concurrent map iteration and map write
}
