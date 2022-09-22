package main

import "fmt"

type ISay interface {
	Say()
}

type SayFunc func()

// 实现ISay接口的Say方法，注意这里不能使用*SayFunc，否则会报错
func (s SayFunc) Say() {
	s()
}

// 入口函数具体执行的函数
func doSay(iSay ISay) {
	iSay.Say()
}

// 入口函数
func Say(handler func()) {
	doSay(SayFunc(handler))
}

//处理器函数
func HandlerSay() {
	fmt.Println("Hello World")
}

func main() {
	Say(HandlerSay)
	//实际相当于这样的方式，反而更简单
	//因为SayFunc和HandlerSay的入参和出参一致，可以做类型转换。
	sayFunc := SayFunc(HandlerSay)
	sayFunc.Say()
}
