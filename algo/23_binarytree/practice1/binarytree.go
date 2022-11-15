package practice1

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
