package practice2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	bst := initTree()
	ast := assert.New(t)
	ast.Equal([]int{1, 3, 2, 7, 5}, bst.PreOrderTraverse(), "前序遍历")
	ast.Equal([]int{1, 2, 3, 5, 7}, bst.InOrderTraverse(), "中序遍历")
	ast.Equal([]int{2, 5, 7, 3, 1}, bst.PostOrderTraverse(), "后序遍历")

}

// 删除的结点有左右子节点
func TestBST_Delete_HasLeftAndRight(t *testing.T) {
	bst := initTree()

	bst.PreOrderTraverse()
	ast := assert.New(t)
	result := bst.Delete(3)
	ast.True(result, "delete result:")
	ast.Equal([]int{1, 2, 5, 7}, bst.InOrderTraverse(), "删除后中序遍历验证")

}

// 删除叶子结点
func TestBST_Delete_Leaf(t *testing.T) {
	bst := initTree()
	ast := assert.New(t)
	result := bst.Delete(2)
	ast.True(result, "delete result:")
	ast.Equal([]int{1, 3, 5, 7}, bst.InOrderTraverse(), "删除后中序遍历验证")
}

// 删除的节点只有左节点或者右节点
func TestBST_Delete_HasLeftOrRight(t *testing.T) {
	bst := initTree()
	ast := assert.New(t)
	result := bst.Delete(7)
	ast.True(result, "delete result:")
	ast.Equal([]int{1, 2, 3, 5}, bst.InOrderTraverse(), "删除后中序遍历验证")
}

// 删除的节点是根节点s
func TestBST_Delete_Root(t *testing.T) {
	bst := initTree()
	ast := assert.New(t)
	result := bst.Delete(1)
	ast.True(result, "delete result:")
	ast.Equal([]int{2, 3, 5, 7}, bst.InOrderTraverse(), "删除后中序遍历验证")

}

// 初始化好一颗树
func initTree() *BinaryTree[int] {
	bst := NewBinaryTree(1)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)
	return bst
}
