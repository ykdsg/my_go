package main

import (
	"fmt"
	"sync"
)

func main() {
	//mistakeDefer()
	rightDefer()
}

func rightDefer() {
	locker := sync.Mutex{}
	for i := 0; i < 10; i++ {
		//通过闭包（也就是函数），来实现函数退出时执行defer
		func() {
			locker.Lock()
			defer locker.Unlock()
			fmt.Println("loop function", i)
		}()
	}

}

func mistakeDefer() {
	locker := sync.Mutex{}
	for i := 0; i < 10; i++ {
		locker.Lock()
		//因为defer 是在函数退出时执行，所以不会在一次循环中释放。
		defer locker.Unlock()
		fmt.Println("loop function", i)
	}
}
