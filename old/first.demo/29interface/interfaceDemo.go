package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happend"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	println("p:", p)
	return p

}

func returnsError2() error {
	if bad() {
		return &ErrBad
	}
	return nil
}

func interfaceNil() {
	println("interfaceNil---------------------------")
	err := returnsError()
	var nilValue error

	//可以看到err 跟p 的内容已经不一样了。
	println("err:", err)
	println("nilValue:", nilValue)
	//err 打印出来是 (0x10c2438,0x0) ，有类型信息，跟nil（0x0,0x0）不能划等号。
	if err != nil {
		fmt.Printf("error occur:%+v\n", err)
	}
	fmt.Println("ok")
	// 如果要判断，这样是可以符合预期的
	nilError := (*MyError)(nil)
	fmt.Println("err == nilError:", err == nilError)

	//注意这里跟err 的区别，因为err 实际返回的是MyError 类型，所以接口的附加信息会更加复杂。这个是由go 底层的interface 数据结构决定的。
	err2 := returnsError2()
	println("err2:", err2)
	fmt.Println("err2== nil:true", err2 == nil)

}

func printNilInterface() {
	println("printNilInterface--------------------")
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println("i", i)
	println("err", err)
	println("i=nil:", i == nil)
	println("err = nil:", err == nil)
	println("i =err:", i == err)
}

func main() {
	interfaceNil()
	printNilInterface()
}
