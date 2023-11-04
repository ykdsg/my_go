package main

import (
	"fmt"
	"github.com/ykdsg/base/util"
	"time"
)

func justPrint() {
	fmt.Println("I am print------------------")
}
func m2() {
	fmt.Println("m2 start----------------", util.GoID())
}

func m1() {
	fmt.Println("m1 start..............", util.GoID())
	go m2()
	justPrint()
	fmt.Println("m1 end..............", util.GoID())

	for i := 0; i < 50000; i++ {
		if i%500 == 0 {
			fmt.Println("current i=", i)
		}
	}
}

//从这个例子中可以看到goroutine 没有这么快启动，但也并不是等函数结束的时候才会启动。
func checkMulityGoroutineSequence() {
	fmt.Println("main start..............", util.GoID())
	go m1()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main end..............", util.GoID())
	time.Sleep(10 * time.Second)
	fmt.Println("main finish..............", util.GoID())
}

func main() {
	checkMulityGoroutineSequence()
}
