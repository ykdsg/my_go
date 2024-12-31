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

func (p field) print2(str string) {
	fmt.Println(str, p.name)

}

// 在for range 执行多线程的情况下，对指针的使用，需要特别当心
func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print("data1-module1-module1")
		//相当于下面的写法
		// go (*field).print(v, "data1-module1-2")

		go v.print2("data1-2-module1")
		// go field.print2(*v, "data1-2-2")
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print("data2-module1-module1")
		//相当于下面的写法
		go (*field).print(&v, "data2-module1-2") // 实际传入的是v的指针
		//打印出来都是six

		go v.print2("data2-2-module1")
		go field.print2(v, "data2-2-2")
	}

	time.Sleep(3 * time.Second)
}
