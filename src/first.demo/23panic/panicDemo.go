package main

import "fmt"

func foo() {
	println("call foo")
	bar()
	println("exit foo")
}

func bar() {
	//通过recover 对painc进行捕获并恢复
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover the painc:", e)
		}
	}()
	println("call bar")
	panic("panic occurs in bar")
	zoo()
	println("exit bar")
}

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func main() {
	println("call main")
	foo()
	println("exit main")
}
