package practice2

import binarytree "23_binarytree"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历，先本身、再左子树、后右子树
func preOrderTraversal(root *TreeNode) []int {
	stack := binarytree.NewStack[*TreeNode]()
	stack.Push(root)
	result := []int{}
	for !stack.IsEmpty() {
		cur := stack.Pop()
		result = append(result, cur.Val)
		stack.Push(cur.Right)
		stack.Push(cur.Left)
	}
	return result
}

// 中序遍历，先左子树，再本身，后右子树
func inOrderTraversal(root *TreeNode) []int {
	stack := binarytree.NewStack[*TreeNode]()
	p := root
	result := []int{}
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.Left
		} else {
			//说明左子树已经遍历完，当前stack 的栈顶是最左子树
			cur := stack.Pop()
			result = append(result, cur.Val)
			//遍历最左子树的右子树
			p = cur.Right
		}
	}

	return result
}

// 后序遍历，先左子树，再右子树，后自身
func postOrderTraversal(root *TreeNode) []int {
	//使用2个stack，第一个stack 放入当前节点，出栈的时候压入stack2，因为当前节点最后输出。
	//然后stack1 陆续放入当前节点的left和right，这样压入stack2 的顺序就是cur,right,left  ，按出栈顺序就是后序遍历
	stack1 := binarytree.NewStack[*TreeNode]()
	stack2 := binarytree.NewStack[*TreeNode]()
	stack1.Push(root)
	for !stack1.IsEmpty() {
		cur := stack1.Pop()
		//实际使用的是stack2，所以自身节点先入栈
		stack2.Push(cur)
		stack1.Push(cur.Left)
		stack1.Push(cur.Right)
	}
	result := []int{}
	for !stack2.IsEmpty() {
		result = append(result, stack2.Pop().Val)
	}

	return result
}
