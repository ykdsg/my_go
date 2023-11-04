package main

import (
	"fmt"
	"reflect"
)

func Add(i, j int) int {
	return i + j
}

type Calculator struct {
}

func (c Calculator) Add(i, j int) int {
	return i + j
}

func main() {
	// 函数调用
	f := reflect.ValueOf(Add)
	var i = 5
	var j = 6
	vals := []reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)}
	ret := f.Call(vals)
	fmt.Println(ret[0].Int())

	// 方法调用，注意跟函数调用的区别
	c := reflect.ValueOf(Calculator{})
	m := c.MethodByName("Add")
	ret = m.Call(vals)
	fmt.Println(ret[0].Int()) // 11
}
