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

func (this *BinaryTree[T]) Delete(v T) bool {
	//	找到对应的位置
	current, parent, isLeft := this.Search(v)
	if current == nil {
		return false
	}
	if current.left != nil && current.right != nil { // 存在左右2个节点，需要找到右子树的最左子结点
		leftNode, leftNodeParent := searchLeftNode(current.right)
		//如果如果leftNodeParent 不为nil，说明右子树最左子节点不是自己
		if leftNodeParent != nil {
			//删除原来的位置
			leftNodeParent.left = leftNode.right
			//调整最左子节点的右子节点为删除节点的右节点
			leftNode.right = current.right
		}
		//最左子节点的左节点调整为删除节点的左节点
		leftNode.left = current.left
		if parent == nil {
			this.root = leftNode
		} else if isLeft {
			parent.left = leftNode
		} else {
			parent.right = leftNode
		}
	} else { //删除的节点仅有一个子节点，这种写法能顺带解决叶子结点的情况
		var child *Node[T]
		if current.left != nil {
			child = current.left
		} else if current.right != nil {
			child = current.right
		} else {
			//删除节点是叶子结点
			child = nil
		}
		if parent == nil {
			this.root = child
		} else if isLeft {
			parent.left = child
		} else {
			parent.right = child
		}
	}
	//	判断是否叶子节点，针对叶子结点处理
	//if current.left == nil && current.right == nil {
	//	if parent == nil {
	//		this.root = nil
	//	} else if isLeft {
	//		parent.left = nil
	//	} else {
	//		parent.right = nil
	//	}
	//} else if current.left != nil && current.right == nil { // 判断是否只有左子节点
	//	if parent == nil {
	//		this.root = parent.left
	//	} else if isLeft {
	//		parent.left = current.left
	//	} else {
	//		parent.right = current.left
	//	}
	//} else if current.right != nil && current.left == nil { // 判断是否只有右子节点
	//	if parent == nil {
	//		this.root = current.right
	//	} else if isLeft {
	//		parent.left = current.right
	//	} else {
	//		parent.right = current.right
	//	}
	//
	//} else { // 存在左右2个节点，需要找到右子树的最左子结点
	//	node, nodeParent := searchLeftNode(current.right)
	//	nodeParent.left = node.right
	//	node.left = current.left
	//	node.right = current.right
	//	if parent == nil {
	//		this.root = node
	//	} else if isLeft {
	//		parent.left = node
	//	} else {
	//		parent.right = node
	//	}
	//}

	return true
}

// 找到最左边的结点
func searchLeftNode[T Number](current *Node[T]) (result *Node[T], parent *Node[T]) {
	if current == nil {
		return
	}
	p := current
	result = p
	for p.left != nil {
		result = p.left
		parent = p
		p = p.left
	}
	return
}

func (this *BinaryTree[T]) Search(v T) (p *Node[T], parent *Node[T], isLeft bool) {
	p = this.root
	if p == nil {
		return
	}
	for p != nil {
		if p.data == v {
			return
		} else if p.data > v {
			parent = p
			p = p.left
			isLeft = true
		} else {
			parent = p
			p = p.right
		}

	}
	return
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
