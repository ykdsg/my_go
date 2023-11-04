package main

import "fmt"

/*
通作为操作符（等号的右边）：
&符号的意思是对变量取地址，如：变量a的地址是&a
*符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了

作为类型符号（等号的左边）：
*符号代表指针类型
*/
func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 300
	//指向a的地址
	ptr = &a
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("ptr address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型，就是原始值的类型
	fmt.Printf("value type: %T\n", value)
	//指向指针ptr的地址
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d，prt=%x \n", *ptr, ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d, pptr=%x\n", **pptr, pptr)
}
