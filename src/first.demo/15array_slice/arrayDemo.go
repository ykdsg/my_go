package main

import "fmt"

func arrayDemo1() {
	var arr1 [6]int //[0 0 0 0 0 0]
	fmt.Printf("arr1 type=  %T\n", arr1)
	fmt.Printf("arr1 =  %v\n", arr1)

	var arr2 = [6]int{
		11, 12, 13, 14, 15, 16,
	}
	fmt.Printf("arr2 type=  %T\n", arr2)
	fmt.Printf("arr2 =  %v\n", arr2)

	var arr3 = [...]int{
		21, 22, 23,
	}
	fmt.Printf("arr3 type=  %T\n", arr3)
	fmt.Printf("arr3 =  %v\n", arr3)

	var arr4 = [...]int{
		9: 39, // 将第9个元素(下标值为9)的值赋值为39，其余元素值均为0
	}
	fmt.Printf("arr4 type=  %T\n", arr4)
	fmt.Printf("arr4 =  %v\n", arr4)
}

func sliceDemo() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(nums)) // 6

	nums = append(nums, 7)
	fmt.Println(len(nums))

	var s []int
	s = append(s, 1)
	fmt.Println("s append 1:", len(s), cap(s)) //1 1
	s = append(s, 2)
	fmt.Println("s append 2:", len(s), cap(s)) //2 2
	s = append(s, 3)
	fmt.Println("s append 3:", len(s), cap(s)) //3 4
	s = append(s, 4)
	fmt.Println("s append 4:", len(s), cap(s)) //4 4
	s = append(s, 5)
	fmt.Println("s append 5:", len(s), cap(s)) //5 8

	sl := make([]byte, 1, 3)
	fmt.Println("sl[0]", sl[0])
	fmt.Println("sl init:", len(sl), cap(sl)) //1 3
	sl = append(sl, 1)
	fmt.Println("sl[1]", sl[1])
	fmt.Println("sl append 1:", len(sl), cap(sl)) //2 3
	sl = append(sl, 2)
	fmt.Println("sl append 2:", len(sl), cap(sl)) //3 3

	sl = append(sl, 3)
	fmt.Println("sl append 3:", len(sl), cap(sl)) //4 8

}

func main() {
	arrayDemo1()
	sliceDemo()
}
