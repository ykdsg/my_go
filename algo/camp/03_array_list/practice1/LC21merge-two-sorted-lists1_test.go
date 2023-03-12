package practice1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}
	head1 := NewHead(list1)
	head2 := NewHead(list2)

	listNode := MergeTwoLists(head1, head2)

	ast := assert.New(t)
	ast.Equal([]int{1, 1, 2, 3, 4, 4}, toList(listNode))
}
