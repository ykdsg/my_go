package main

import "fmt"

type T int

func (t T) Error() string {
	return "bad error"
}

func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	err1 = (*T)(nil)
	println("err1:", err1)            //(0x10c07f8,0x0)
	println("err1=nil:", err1 == nil) //false
	println("err2:", err2)            //(0x0,0x0)
	println("err2=nil:", err2 == nil) //true

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1=err2:", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1=err2:", err1 == err2)

}

func printEmptyInterfaceAndNonEmptyInterface() {
	println("printEmptyInterfaceAndNonEmptyInterface----------------------------")
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	//空接口类型变量和非空接口类型变量内部表示的结构有所不同（第一个字 段：_type vs. tab)，两者似乎一定不能相等。但 Go 在进行等值比较时，类型比较使用的
	//是 eface 的 _type 和 iface 的 tab._type，因此就像我们在这个例子中看到的那样，当 eif 和 err 都被赋值为T(5)时，两者之间是划等号的。
	println("eif=err:", eif == err) //true

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif =err:", eif == err) //false
}

func main() {
	printNonEmptyInterface()

	printEmptyInterfaceAndNonEmptyInterface()
}
