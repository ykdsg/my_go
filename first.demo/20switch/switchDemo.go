package main

import "fmt"

func typeSwitch() {
	var x interface{} = 13
	switch x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("the type of x is int")
	default:
		println("don't support the type")
	}
}

func typeSwitchValue() {
	var x interface{} = 13
	//v 存储的是变量 x 的动态类型对应的值信息
	switch v := x.(type) {
	case nil:
		println("x is nil")
	case int:
		println("the type of x is int,v=", v)
	default:
		println("don't support the type")
	}
}

func caseStitch() {
	i := 1
	j := 2
	z := 3
	switch {
	case i == 1:
		fmt.Println("i==1")
		//fallthrough 将直接执行下一个case的内容，而不经过校验
		fallthrough
	case j != 2:
		fmt.Println("j==1")
	case z == 3:
		fmt.Println("z==3")
	default:
		fmt.Println("default")

	}

}

func main() {
	//typeSwitch()
	//typeSwitchValue()
	caseStitch()
}
