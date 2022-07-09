package main

import (
	"fmt"
	"reflect"
)

//输出方法集合
func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)
	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Println("\n")
}

type V struct {
}

func (V) M1() {
}

func (V) M2() {
}

func (*V) M3() {
}

func (*V) M4() {
}

//S类型和*S类型都没有包含任何方法，因为type S 定义了一个新类型。
type S V

func main() {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t V
	dumpMethodSet(t)
	dumpMethodSet(&t)

	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)
}
