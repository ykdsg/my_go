package main

import (
	"fmt"
	"log"
	"time"
)

//代表处理器接口
type Handler interface {
	RealDo(string)
}

// logHandler，authHandler 本质是包装函数，内部利用了是配置函数类型
func logHandler(h Handler) Handler {
	return HandlerFunc(func(str string) {
		t := time.Now()
		log.Printf("str:[%s] %v \n", t, str)
		h.RealDo(str)
	})
}

func authHandler(h Handler) Handler {
	return HandlerFunc(func(str string) {
		log.Println("authHandler...")
		h.RealDo(str)
	})
}

//使用接口的上层方法
func doHandler(str string, handler Handler) {
	handler.RealDo(str)
}

//HandlerFunc 就是一个适配器，将普通函数转化为满足Handler接口的类型
type HandlerFunc func(string)

func (f HandlerFunc) RealDo(str string) {
	f(str)
}

func main() {
	doHandler("hello", HandlerFunc(justPrint))
	println("-------------middleware-----------------")
	doHandler("middle world", logHandler(authHandler(HandlerFunc(justPrint))))

}

func justPrint(str string) {
	fmt.Println("now we print:", str)
}
