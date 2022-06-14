package main

import (
	"fmt"
	"unicode/utf8"
)

func stringNormal() {
	var s string = "hello"
	//下面这种写法是有问题的
	//s[0] = 'k'
	s = "gopher"

	fmt.Println("new s=" + s)
	//支持原始字符串，所见即所得
	var s2 = `she 
	sh 
	sds
	//@\\
	sd`
	fmt.Println(s2)
}

func stringLength() {
	var s = "中国人"
	fmt.Printf("the length of s =%d\n", len(s))

	//字节视角
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
	fmt.Println()
}

func stringCount() {
	var s = "中国人"
	fmt.Println("the count in s is", utf8.RuneCountInString(s))

	//字符视角：使用range的方式
	for _, c := range s {
		fmt.Printf("0x%x ", c)
	}
	fmt.Println()
}

func main() {
	stringNormal()

	stringLength()

	stringCount()
}
