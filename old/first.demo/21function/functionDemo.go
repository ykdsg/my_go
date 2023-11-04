package main

import (
	"fmt"
	"io"
	"os"
)

var (
	myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, a...)
	}
)

//匿名函数使用了定义它的函数 setup 的局部变量 task，这样的匿名函数在 Go 中也被称为闭包（Closure）。
func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

func main() {
	fmt.Printf("%T\n", myFprintf) // func(io.Writer, string, ...interface {})
	//函数可以放在变量中
	myFprintf(os.Stdout, "%s\n", "Hello, Go") // 输出Hello，Go

	teardown := setup("demo")
	teardown()
	defer teardown()
	println("do some bussiness stuff")
}
