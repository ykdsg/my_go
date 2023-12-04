package main

import "unsafe"

type T struct {
	b byte
	i int64
	u uint16
}

type S struct {
	b byte
	u uint16
	i int64
}

// 字段排列顺序不同,导致它们的大小
func main() {
	var t T
	println(unsafe.Sizeof(t))
	var s S
	println(unsafe.Sizeof(s))
}
