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

func interfaceNil() {
	err := returnsError()
	var nilValue error

	//可以看到err 跟p 的内容已经不一样了。
	println("err:", err)
	println("nilValue:", nilValue)
	//err 打印出来是 (0x10c2438,0x0) ，有类型信息，跟nil（0x0,0x0）不能划等号。
	if err != nil {
		fmt.Printf("error occur:%+v\n", err)
	}
	// 如果要判断，这样是可以符合预期的
	nilError := (*MyError)(nil)
	if err == nilError {
		fmt.Println("err == nilError")
	}
	fmt.Println("ok")
}

func printNilInterface() {
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
