package practice1

type BinaryTree[T Number] struct {
	root *Node[T]
}

func NewBinaryTree[T Number](rootV T) *BinaryTree[T] {
	return &BinaryTree[T]{
		root: NewNode(rootV),
	}
}

func (this *BinaryTree[T]) Find(v T) *Node[T] {
	p := this.root
	for nil != p {
		if p.data == v {
			return p
		} else if p.data < v {
			p = p.right
		} else {
			p = p.left
		}
	}
	return nil
}

func (this *BinaryTree[T]) Insert(v T) bool {
	p := this.root
	for nil != p {
		if p.data == v {
			return false
		} else if p.data < v {
			if p.right == nil {
				p.right = NewNode(v)
				return true
			} else {
				p = p.right
			}
		} else {
			if p.left == nil {
				p.left = NewNode(v)
				return true
			} else {
				p = p.left
			}
		}
	}
	return true
}

// 树的各种遍历，用递归实现还是挺清晰简单的，而且相通性比较好，基本写出来一个排序之后，其他的排序差别不大。
// 中序遍历，递归实现
func InOrderTraverse_recursive[T Number](root *Node[T]) (result []T) {
	if root == nil {
		return result
	}
	result = append(result, InOrderTraverse_recursive(root.left)...)
	result = append(result, root.data)
	result = append(result, InOrderTraverse_recursive(root.right)...)
	return result
}

// 前序遍历，递归实现
func PreOrderTraverse_recursive[T Number](root *Node[T]) (result []T) {
	if root == nil {
		return result
	}
	result = append(result, root.data)
	result = append(result, PreOrderTraverse_recursive(root.left)...)
	result = append(result, PreOrderTraverse_recursive(root.right)...)
	return result
}

// 后序遍历，递归实现
func PostOrderTraverse_recursive[T Number](root *Node[T]) (result []T) {
	if root == nil {
		return result
	}
	result = append(result, PostOrderTraverse_recursive(root.left)...)
	result = append(result, PostOrderTraverse_recursive(root.right)...)
	result = append(result, root.data)
	return result
}

//树的遍历用非递归的方式实现，这个难度就上升了不少，需要栈来配合，而且不同遍历顺序的实现差别较大。

// 中序遍历，如果是二叉查找树就是顺序输出
func (this *BinaryTree[T]) InOrderTraverse() (result []T) {
	p := this.root
	if p == nil {
		return result
	}
	stack := NewStack[*Node[T]]()
	//p指向当前元素，先将p存入stack，如果p.left 不为nil，继续将该值存入stack
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			//p入栈，接着查看left
			stack.Push(p)
			p = p.left
		} else {
			// 如果p为空，说明p.left或者p.right 为nil，说明上一个节点已经处理完，从栈中拿出一个node 来处理，再将node.right入栈
			//通过stack.Pop 出来的只要处理右子树就可以了，因为左子树已经包含再stack 中就是current
			current := stack.Pop()
			//打印当前节点
			result = append(result, current.data)
			//p指向right，进行下一轮处理
			p = current.right
		}
	}
	return result
}

// 前序遍历
func (this *BinaryTree[T]) PreOrderTraverse() (result []T) {
	p := this.root
	if p == nil {
		return result
	}
	stack := NewStack[*Node[T]]()
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			result = append(result, p.data)
			stack.Push(p.right)
			p = p.left
		} else {
			p = stack.Pop()
		}
	}
	return result
}

// 后序遍历
func (this *BinaryTree[T]) PostOrderTraverse() (result []T) {
	p := this.root
	if p == nil {
		return
	}
	stack1 := NewStack[*Node[T]]()
	stack1.Push(p)
	stack2 := NewStack[*Node[T]]()
	for !stack1.IsEmpty() {
		current := stack1.Pop()
		if current == nil {
			continue
		}
		stack2.Push(current)
		stack1.Push(current.left)
		stack1.Push(current.right)
	}

	for !stack2.IsEmpty() {
		t := stack2.Pop()
		result = append(result, t.data)
		//fmt.Printf("%+v ", t.data)
	}
	return result
}
