package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

var goroutineSpace = []byte("goroutine ")

//获取Goroutine ID，类似线程id
func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, goroutineSpace)

	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func Trace() func() {
	//通过 runtime.Caller 函数获得当前 Goroutine 的函数调用 栈上的信息，runtime.Caller 的参数标识的是要获取的是哪一个栈帧的信息
	//当参数为 0 时，返回的是 Caller 函数的调用者的函数信息，在这里就是 Trace 函数。但我们需要的是 Trace 函数的调用者的信息，于是我们传入 1
	//Caller 函数有四个返回值：第一个返回值代表的是程序计数（pc）；第二个和第三个参数 代表对应函数所在的源文件名以及所在行数,最后一个参数代表是 否能成功获取这些信息
	pc, _, _, ok := runtime.Caller(1)

	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	println("enter:", name)
	return func() {
		println("exit:", name)
	}
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

func main() {
	defer Trace()()
	foo()
}
