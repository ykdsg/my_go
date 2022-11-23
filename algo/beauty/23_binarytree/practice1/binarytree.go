package practice1

import "23_binarytree"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历，先本身、再左子树、后右子树
func preOrderTraversal_stack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var stack []*TreeNode
	var result []int
	stack = append(stack, root)
	//这里用栈能够比较方便处理是因为前序遍历是先处理当前节点，这样当前节点可以直接在栈中删除
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		result = append(result, current.Val)
		stack = stack[0 : len(stack)-1]
		stack = append(stack, current.Right)
		stack = append(stack, current.Left)
	}
	return result
}

// 使用递归的方式代码更清晰
func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var result []int
	result = append(result, root.Val)
	result = append(result, preOrderTraversal(root.Left)...)
	result = append(result, preOrderTraversal(root.Right)...)
	return result
}

// 中序遍历，栈实现
func inOrderTraversal_stack(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	stack := binarytree.NewStack[*TreeNode]()
	p := root
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.Left
		} else {
			t := stack.Pop()
			result = append(result, t.Val)
			p = t.Right
		}
	}
	return result
}

// 中序遍历，先左子树，再本身，后右子树
func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var result []int
	result = inOrderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inOrderTraversal(root.Right)...)
	return result
}

// 后序遍历，栈实现
func postOrderTraversal_stack(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	stack1 := binarytree.NewStack[*TreeNode]()
	stack2 := binarytree.NewStack[*TreeNode]()

	stack1.Push(root)
	for !stack1.IsEmpty() {
		t := stack1.Pop()
		stack2.Push(t)
		//注意这里需要先放left，后放right，因为这里实际要用的是stack2，因此stack1 先处理的，在stack2 这边就会后处理
		stack1.Push(t.Left)
		stack1.Push(t.Right)

	}
	for !stack2.IsEmpty() {
		t := stack2.Pop()
		result = append(result, t.Val)
	}
	return result
}

// 后序遍历，先左子树，再右子树，后自身
func postOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var result []int
	result = postOrderTraversal(root.Left)
	result = append(result, postOrderTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}
