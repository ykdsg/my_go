package practice2

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Node[T Number] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T Number](v T) *Node[T] {
	return &Node[T]{
		value: v,
	}

}
