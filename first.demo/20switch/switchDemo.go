package main

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

func main() {
	typeSwitch()
	typeSwitchValue()
}
