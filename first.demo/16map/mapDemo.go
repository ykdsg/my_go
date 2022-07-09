package main

import "fmt"

func mapdemo() {
	m := make(map[int]string)
	m[1] = "value1"
	m[2] = "value2"
	m[3] = "value3"

	fmt.Println(len(m))
	fmt.Println("map[3]=", m[3])

	v, ok := m[5]
	if !ok {
		fmt.Println("5 is not in m,v=", v)
	}

	delete(m, 1)
	fmt.Println("after delete m=", m)

	for k, v := range m {
		fmt.Printf("[%d,%s]", k, v)
	}

}
func main() {
	mapdemo()
}
