package practice2

type Stack[T comparable] struct {
	data []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (this *Stack[T]) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *Stack[T]) Push(v T) {
	this.data = append(this.data, v)
}

func (this *Stack[T]) Pop() (t T) {
	if len(this.data) == 0 {
		return
	}
	lastIndex := len(this.data) - 1
	t = this.data[lastIndex]
	this.data = this.data[:lastIndex]
	return t
}
