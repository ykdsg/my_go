package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	ast := assert.New(t)

	list := []int{1, 2, 3, 4}
	head := NewHead(list)
	newHead := swapPairs(head)

	ast.Equal([]int{2, 1, 4, 3}, toList(newHead))

}

func TestSwapPairs3(t *testing.T) {
	ast := assert.New(t)

	list := []int{1, 2, 3, 4}
	head := NewHead(list)
	newHead := swapPairs3(head)

	ast.Equal([]int{2, 1, 4, 3}, toList(newHead))

}
