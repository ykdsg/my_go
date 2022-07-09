package main

import (
	"fmt"
	"time"
)

//for 循环的变量是公用的
func demo1Bug() {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 5)
}

//闭包函数，可以避免上面的问题
func demo2Fix() {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func(i int, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 5)
}

func main() {
	demo1Bug()
	fmt.Println("after fix ...")
	demo2Fix()
}
