package common

import "fmt"

// 基于分片实现栈
type Stack[T comparable] struct {
	data []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return len((*s).data) == 0

}

func (s *Stack[T]) Push(v T) {
	(*s).data = append((*s).data, v)
}

func (s *Stack[T]) Pop() (t T) {
	data := (*s).data
	if len(data) == 0 {
		return
	}

	t = data[len(data)-1]
	(*s).data = data[:len(data)-1]
	return

}

func (s *Stack[T]) Print() {
	data := (*s).data
	if s.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := len(data) - 1; i >= 0; i-- {
			fmt.Println((data)[i])
		}
	}

}
