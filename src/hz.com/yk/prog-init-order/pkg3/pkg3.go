package pkg3

import "fmt"

const (
	c1 = "c1"
)

var (
	_  = constInitCheck()
	v1 = variableInit("v1")
)

func variableInit(name string) string {
	fmt.Printf("pkg3: var %s has been initalized\n", name)
	return name
}

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("pkg3: const c1 has been initialized")
	}
	return ""
}

func init() {
	fmt.Println("pkg3: init func invoked")
}
