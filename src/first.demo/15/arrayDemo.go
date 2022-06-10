package main

import "fmt"

func arrayDemo1() {
	var arr1 [6]int //[0 0 0 0 0 0]
	fmt.Printf("arr1 type=  %T\n", arr1)
	fmt.Printf("arr1 =  %v\n", arr1)

	var arr2 = [6]int{
		11, 12, 13, 14, 15, 16,
	}
	fmt.Printf("arr1 type=  %T\n", arr2)
	fmt.Printf("arr1 =  %v\n", arr2)

	var arr3 = [...]int{
		21, 22, 23,
	}
	fmt.Printf("arr1 type=  %T\n", arr3)
	fmt.Printf("arr1 =  %v\n", arr3)
}

func main() {
	arrayDemo1()
}
