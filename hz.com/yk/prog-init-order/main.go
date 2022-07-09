package main

import (
	"fmt"
	_ "hz.com/yk/prog-init-order/pkg1"
	_ "hz.com/yk/prog-init-order/pkg2"
)

const (
	c1 = "c1"
	c2 = "c2"
)

var (
	_  = constInitCheck()
	v1 = variableInit("v1")
)

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initalized\n", name)
	return name
}

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 has been initialized")
	}
	return ""
}

func init() {
	fmt.Println("main: firset init func invoked")
}

func init() {
	fmt.Println("main: second init func invoked")
}

func main() {

}
