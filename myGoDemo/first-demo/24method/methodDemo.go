package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print(str string) {
	fmt.Println(str, p.name)
}

// 在for range 执行多线程的情况下，对指针的使用，需要特别当心
func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print("data1-module1-module1")
		// 相当于下面的写法
		// go (*field).print(v, "data1-module1-2")

	}

	// 注意 go1.23 版本打印的不一样
	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print("data2-module1-module1")
		// 相当于下面的写法
		go (*field).print(&v, "data2-module1-2") // 实际传入的是v的指针
		// 打印出来都是six

	}
	time.Sleep(3 * time.Second)
}
