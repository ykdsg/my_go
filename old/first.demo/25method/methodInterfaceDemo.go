package main

type Interface interface {
	M1()
	M2()
}

type TT struct{}

func (t TT) M1()  {}
func (t *TT) M2() {}

func main() {
	var pt *TT
	var i Interface

	i = pt
	// var t TT
	// i = t // 报错：TT does not implement Interface (method M2 has pointer receiver)

	println(i)
}
