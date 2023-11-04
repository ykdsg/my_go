package practice1

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Node[T Number] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T Number](data T) *Node[T] {
	return &Node[T]{
		data: data,
	}
}

func (this *Node[T]) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", this.data, this.left, this.right)
}
