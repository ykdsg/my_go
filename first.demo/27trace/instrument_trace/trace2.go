package trace

import (
	"fmt"
	"runtime"
	"sync"
)

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "    "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	//通过 runtime.Caller 函数获得当前 Goroutine 的函数调用栈上的信息
	//runtime.Caller 的参数标识的是要获取的是哪一个栈帧的信息。
	//当参数为 0 时，返回的是 Caller 函数的调用者的函数信息，在这里就是 Trace 函数。但我们需要的是 Trace 函数的调用者的信息，于是我们传入1。
	pc, _, _, ok := runtime.Caller(1)
	//Caller 函数有四个返回值：第一个返回值代表的是程序计数（pc）；第二个和第三个参数 代表对应函数所在的源文件名以及所在行数,
	//最后一个参数代表是否能成功获取这些信息
	if !ok {
		panic("not found caller")
	}
	//通过 runtime.FuncForPC 函数和程序计数器（PC）得到被跟踪函数的函数名称
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()

	//map 不支持并发写，通过sync.Mutex 实例 mu 用于同步对 m 的写操作
	mu.Lock()
	//获取当前gid对应的缩进层次
	indents := m[gid]
	//缩进层次+1
	m[gid] = indents + 1
	mu.Unlock()
	printTrace(gid, name, "->", indents+1)
	return func() {
		mu.Lock()
		//当前guid对应的缩进层次
		indents := m[gid]
		//缩进层次-1
		m[gid] = indents - 1
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}

}
func A1() {
	defer Trace()()
	B1()
}

func B1() {
	defer Trace()()
	C1()
}

func C1() {
	defer Trace()()
	D()
}

func D() {
	defer Trace()()
}

func A2() {
	defer Trace()()
	B2()

}

func B2() {
	defer Trace()()
	C2()
}

func C2() {
	defer Trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()
}
