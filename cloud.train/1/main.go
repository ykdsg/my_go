package main

import "fmt"

func main() {
	strArr := [5]string{"I", "am", "stupid", "and", "weak"}
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
