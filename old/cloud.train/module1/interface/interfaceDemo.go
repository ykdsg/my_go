package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName string
	lastName  string
}

//human是对象也代表了指针有对应的方法
func (human Human) getName() string {
	return human.firstName + human.lastName
}

type Plane struct {
	vendor string
	model  string
}

//plane是指针，但是对象并没有对应的方法
func (plane *Plane) getName() string {
	return plane.vendor + plane.model

}

func main() {
	//new返回的是指针
	ifs1 := new([]IF)
	fmt.Printf("ifs1 type=%T \n", ifs1)

	//make 返回的是对象
	ifs2 := make([]IF, 0)
	fmt.Printf("ifs2 type=%T \n", ifs2)

	//这样相当于make
	interfaces := []IF{}
	//h实际是指针
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)

	human := Human{
		firstName: "y",
		lastName:  "k",
	}
	interfaces = append(interfaces, human)
	plane := Plane{
		vendor: "t",
		model:  "119",
	}
	interfaces = append(interfaces, &plane)

}
