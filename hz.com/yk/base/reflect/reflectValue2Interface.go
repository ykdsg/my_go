package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i = 5
	//通过reflect.Value.Interface()函数重建后得到的新变量（如例子中的r）
	//与原变量（如例子中的i）是两个不同的变量，它们的唯一联系就是值相同
	val := reflect.ValueOf(i)
	r := val.Interface().(int)
	fmt.Println(r) // 5
	r = 6
	fmt.Println(i, r) //5 6

	//如果我们反射的对象是一个指针（如例子中的&i），那么我们通过reflect.Value.Interface()
	//得到的新变量（如例子中的q）也是一个指针，且它所指的内存地址与原指针变量相同。
	val = reflect.ValueOf(&i)
	q := val.Interface().(*int)
	fmt.Printf("%p,%p,%d\n", &i, q, *q) //0xc000194008,0xc000194008,5
	*q = 7
	fmt.Println(i) //7
}
