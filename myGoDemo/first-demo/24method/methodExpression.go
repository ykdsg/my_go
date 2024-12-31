package main

import "fmt"

type T struct{ a int }

func (t T) Get() int { // 等价函数： func Get(t T) int { return t.a }
	return t.a
}

func (t *T) Set(a int) int { // 等价函数： func Set(t *T, a int) int {
	t.a = a
	return t.a
}

func main() {
	var t T
	t.Set(2)
	fmt.Println(t.Get())
}
