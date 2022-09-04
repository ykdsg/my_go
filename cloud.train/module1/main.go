package main

import "fmt"

func main() {
	strArr := [5]string{"I", "am", "stupid", "and", "weak"}
	for i, str := range strArr {
		fmt.Printf("i=%d,s=%s \n", i, str)
		//直接操作str 是不会影响strArr中的值
		str = "i"
	}
	fmt.Println(strArr)

	for i := range strArr {
		if i == 2 {
			strArr[i] = "smart"
		}
		if i == 4 {
			strArr[i] = "strong"
		}
	}
	fmt.Println(strArr)
}
