package practice1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	pre, in, post []int
}{

	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{3, 2, 1},
	},

	{
		[]int{1, 2, 4, 5, 3, 6, 7},
		[]int{4, 2, 5, 1, 6, 3, 7},
		[]int{4, 5, 2, 6, 7, 3, 1},
	},
	// 可以有多个 testCase
}

func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}

	if len(in) == 0 {
		return nil
	}

	//前序遍历，第一个节点是当前节点
	res := &TreeNode{
		Val: pre[0],
	}

	//找到中序遍历的位置
	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)
	//左子树的前序数组、中序数组
	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	//右子树的前序数组、中序数组
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}

func Test_preOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.pre, preOrderTraversal(root), "输入:%v", tc)
		ast.Equal(tc.pre, preOrderTraversal_stack(root), "输入:%v", tc)
	}
}

func Test_inOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.in, inOrderTraversal(root), "输入:%v", tc)
		ast.Equal(tc.in, inOrderTraversal_stack(root), "输入:%v", tc)
	}
}

func Test_postOrderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.post, postOrderTraversal(root), "输入:%v", tc)
		ast.Equal(tc.post, postOrderTraversal_stack(root), "输入:%v", tc)
	}
}
