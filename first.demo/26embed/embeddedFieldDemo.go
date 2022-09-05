package main

import (
	"fmt"
	"io"
	"strings"
)

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type SecInt int

func (n *SecInt) Add(m int) {
	*n = *n + SecInt(m)
}

type I interface {
	M1()
	M2()
}
type S struct {
	*MyInt
	//如果存在相同方法的类型嵌入，直接使用s.Add会有编译错误
	//*SecInt
	t
	I
	io.Reader
	s string
	n int
}

func main() {
	m := MyInt(17)
	r := strings.NewReader("hello,go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}
	var sl = make([]byte, len("hello,go"))
	//s.Reader.Read(sl)
	s.Read(sl) // 等同于 s.Reader.Read(sl)
	fmt.Println(string(sl))
	//s.MyInt.Add(5)
	s.Add(5) // 等同于 s.MyInt.Add(5)
	fmt.Println(*(s.MyInt))
	//因为S 没有对应的M1方法，所以会报错
	//s.M1()

}
