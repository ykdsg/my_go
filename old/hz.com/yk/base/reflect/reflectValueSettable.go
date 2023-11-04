package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	age  int
}

func main() {
	var n = 17
	fmt.Println("int:")
	val := reflect.ValueOf(n)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true

	fmt.Println("\n*int:")
	val = reflect.ValueOf(&n)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true

	//	通过reflect.Value的Elem方法可以得到代表该指针所指内存对象的Value反射对象。而这个反射对象是可设置和可寻址的
	val = reflect.ValueOf(&n).Elem()
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true

	fmt.Println("\nslice:")
	var sl = []int{5, 6, 7}
	val = reflect.ValueOf(sl)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Index(0)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true

	fmt.Println("\narray:")
	var arr = [3]int{5, 6, 7}
	val = reflect.ValueOf(arr)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Index(0)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true

	fmt.Println("\nptr to array:")
	var pArr = &[3]int{5, 6, 7}
	val = reflect.ValueOf(pArr)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Elem()
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true
	val = val.Index(0)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true

	fmt.Println("\nstruct:")
	p := Person{"tony", 33}
	val = reflect.ValueOf(p)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val1 := val.Field(0) // Name
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val1.CanSet(), val1.CanAddr(), val1.CanInterface()) // false false true
	val2 := val.Field(1) // age
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val2.CanSet(), val2.CanAddr(), val2.CanInterface()) // false false false

	fmt.Println("\nptr to struct:")
	pp := &Person{"tony", 33}
	val = reflect.ValueOf(pp)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Elem()
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true
	val1 = val.Field(0) // Name
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val1.CanSet(), val1.CanAddr(), val1.CanInterface()) // true true true
	//	通过reflect.Value的Elem方法可以得到代表该指针所指内存对象的Value反射对象。而这个反射对象是可设置和可寻址的
	val2 = val.Field(1) // age
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val2.CanSet(), val2.CanAddr(), val2.CanInterface()) // false true false

	fmt.Println("\ninterface:")
	var i interface{} = &Person{"tony", 33}
	val = reflect.ValueOf(i)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	val = val.Elem()
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // true true true

	fmt.Println("\nmap:")
	var m = map[string]int{
		"tony": 23,
		"jim":  34,
	}
	val = reflect.ValueOf(m)
	fmt.Printf("Settable = %v, CanAddr = %v, CanInterface = %v\n",
		val.CanSet(), val.CanAddr(), val.CanInterface()) // false false true
	//map类型被反射对象比较特殊，它的key和value都是不可寻址和不可设置的。但我们可以通过Value提供的SetMapIndex方法对map反射对象进行修改，这种修改会同步到被反射的map变量中。
	val.SetMapIndex(reflect.ValueOf("tony"), reflect.ValueOf(12))
	fmt.Println(m) // map[jim:34 tony:12]
}
