package main

import "fmt"

//函数类型是表示所有包含相同参数和返回类型的函数集合。
type Greeting func(name string) string

//作为Greeting的方法
func (g Greeting) say(name string) {
	fmt.Println(g(name))
}

//作为普通函数
func say(g Greeting, name string) {
	fmt.Println(g(name))
}

func english(name string) string {
	return "hello" + name

}

func main() {
	say(english, "world")

	greeting := Greeting(english)
	greeting.say("world2")
}
