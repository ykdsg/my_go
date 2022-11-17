package practice1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST_Find(t *testing.T) {
	bst := NewBinaryTree(1)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Find(2))
}

func TestBST_Insert(t *testing.T) {
	bst := NewBinaryTree(1)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	bst.PreOrderTraverse()

	bst.PostOrderTraverse()
	ast := assert.New(t)
	ast.Equal([]int{1, 2, 3, 5, 7}, bst.InOrderTraverse(), "中序遍历")
	ast.Equal([]int{1, 2, 3, 5, 7}, InOrderTraverse_recursive(bst.root), "中序递归遍历")

	ast.Equal([]int{1, 3, 2, 7, 5}, bst.PreOrderTraverse(), "前序遍历")
	ast.Equal([]int{1, 3, 2, 7, 5}, PreOrderTraverse_recursive(bst.root), "前序递归遍历")

	ast.Equal([]int{2, 5, 7, 3, 1}, bst.PostOrderTraverse(), "后序遍历")
	ast.Equal([]int{2, 5, 7, 3, 1}, PostOrderTraverse_recursive(bst.root), "后序递归遍历")
}
