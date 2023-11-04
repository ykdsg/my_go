package main

import (
	"fmt"
	"time"
	"unsafe"
)

type runtimeTimer struct {
	when   int64
	f      func(interface{}, uintptr)
	status uint32
}

type Timer struct {
	C <-chan time.Time
	r runtimeTimer
}

func NewTimer(d time.Duration) *Timer {
	c := make(chan time.Time, 1)
	t := &Timer{
		C: c,
		r: runtimeTimer{
			//when: when(d),
			//f:sendTime
		},
	}
	//startTimer(&t.r)
	return t
}

func main() {
	var t Timer
	sizeofTimer := unsafe.Sizeof(t)   // 结构体类型变量占用的内存大小
	offsetofR := unsafe.Offsetof(t.r) // 字段r在内存中相对于变量t起始地址的偏移量
	fmt.Println("size of Timer=", sizeofTimer)
	fmt.Println("offset of r=", offsetofR)
}
