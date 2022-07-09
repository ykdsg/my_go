package main

import "fmt"

type Test struct {
	name string
}

//&是“取地址运算符”，是从一个变量获取地址
//
//*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
//*是“解引用运算符”，可以简单理解为“从地址取值”， 是&的逆运算
func main() {
	a := Test{name: "test"}
	fmt.Println("a:", a)   //{test}
	fmt.Println("&a:", &a) //&{test}
	//a 不是指针，所以不能用*a
	//fmt.Println("*a:", *a)
	//c 是a的一个指针，是*Test 类型
	c := &a
	// 等同于
	// var c *Test = &a
	fmt.Println("c:", c)           //&{test}
	fmt.Println("c.name:", c.name) //test
	//查看指针变量所指向的存储单元，也就是这个指向的值
	fmt.Println("*c:", *c) //{test}
	//查看指针变量的地址
	fmt.Println("&c:", &c) //0xc0000ac020

	//b 是一个Test* 类型，是一个指针
	b := &Test{"test"}
	fmt.Println("b:", b)   //&{test}
	fmt.Println("*b:", *b) // {test}
	//&b，变量b本身的地址。
	fmt.Println("&b:", &b)           //b指针地址：0xc00000e030
	fmt.Println("&b.name:", &b.name) //b.name指针地址：0xc000096260

	var i int = 1
	fmt.Println("i:", i)

	fmt.Println("&i", &i) //0xc00001e0f0
}
