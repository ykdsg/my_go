package main

import (
	"fmt"
	"reflect"
)

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// 获取Zero常量的反射类型对象
	typeOfA = reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	typeOfStruct := reflect.TypeOf(&cat{})
	fmt.Println(typeOfStruct.Name(), typeOfStruct.Kind())

}
