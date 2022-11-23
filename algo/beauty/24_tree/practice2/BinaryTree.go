package practice2

type BinaryTree[T Number] struct {
	root *Node[T]
}

func NewBinaryTree[T Number](rootv T) *BinaryTree[T] {
	return &BinaryTree[T]{
		root: NewNode(rootv),
	}
}

func (this *BinaryTree[T]) Insert(v T) bool {
	p := this.root
	for p != nil {
		if p.value == v {
			return false
		} else if v < p.value {
			if p.left != nil {
				p = p.left
			} else {
				p.left = NewNode(v)
			}
		} else if v > p.value {
			if p.right != nil {
				p = p.right
			} else {
				p.right = NewNode(v)
			}
		}
	}
	return true
}

func (this *BinaryTree[T]) Delete(v T) bool {
	p, pp, isLeft := this.Find(v)
	if p == nil {
		return false
	}
	//左右节点都有的情况，需要找到右节点的最左子节点（也可能是右节点自己）
	if p.left != nil && p.right != nil {
		mostLeft, mostLeftP := getRightNodeMostLeft(p)
		if pp == nil { //要删除的是父节点
			//右节点的最左子节点是自己，说明他没有左子结点
			if mostLeftP == p {
				mostLeft.left = p.left
			} else {
				mostLeftP.left = mostLeft.right
				mostLeft.left = p.left
				mostLeft.right = p.right
			}
			this.root = mostLeft
		} else { // 删除的不是父节点的情况
			if mostLeftP == p {
				mostLeft.left = p.left
			} else {
				mostLeftP.left = mostLeft.right
				mostLeft.left = p.left
				mostLeft.right = p.right
			}
			if isLeft {
				pp.left = mostLeft
			} else {
				pp.right = mostLeft
			}

		}

	} else { //进入else 说明要么是叶子结点，要么只有一个子节点
		inheritor := p.left
		if inheritor == nil {
			inheritor = p.right
		}

		//如果本身是root
		if pp == nil {
			this.root = inheritor
		} else {
			if isLeft {
				pp.left = inheritor
			} else {
				pp.right = inheritor
			}
		}
	}

	return true
}

// 查找该节点右子树的最左子节点
func getRightNodeMostLeft[T Number](node *Node[T]) (p *Node[T], pp *Node[T]) {
	if node == nil || node.right == nil {
		return
	}
	pp = node
	p = node.right
	for p != nil {
		if p.left == nil {
			return
		}
		pp = p
		p = p.left
	}
	return
}

// 最多存在一个非空的节点
func isMostOneNonNil[T Number](nodes ...*Node[T]) bool {
	if len(nodes) == 1 && nodes[0] != nil {
		return true
	}
	nonNilCount := 0
	for _, v := range nodes {
		if v != nil {
			nonNilCount += 1
		}
	}
	return nonNilCount <= 1
}

// 查找对应的节点，并返回父节点，当前节点为空说明没找到，父节点为空是root节点
func (this *BinaryTree[T]) Find(v T) (p *Node[T], pp *Node[T], isLeft bool) {
	node := this.root
	for node != nil {
		if node.value == v {
			p = node
			return
		} else if v < node.value {
			pp = node
			node = node.left
			isLeft = true
		} else {
			pp = node
			node = node.right
			isLeft = false
		}
	}
	return nil, nil, false
}

// 前序遍历
func (this *BinaryTree[T]) PreOrderTraverse() []T {
	root := this.root
	if root == nil {
		return nil
	}
	var result []T
	stack := NewStack[*Node[T]]()
	stack.Push(root)
	for !stack.IsEmpty() {
		p := stack.Pop()
		result = append(result, p.value)
		if p.right != nil {
			stack.Push(p.right)
		}
		if p.left != nil {
			stack.Push(p.left)
		}
	}
	return result
}

// 中序遍历
func (this *BinaryTree[T]) InOrderTraverse() []T {
	root := this.root
	if root == nil {
		return nil
	}
	var result []T
	p := root
	stack := NewStack[*Node[T]]()
	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.left
		} else {
			t := stack.Pop()
			result = append(result, t.value)
			p = t.right
		}
	}

	return result
}

func (this *BinaryTree[T]) PostOrderTraverse() []T {
	root := this.root
	if root == nil {
		return nil
	}
	var result []T
	//利用2个栈，stack1用来顺序遍历，stack2用来记录当前出栈的内容
	//相当于stack1出栈的节点先进stack2，再看该节点有无左右子节点，有的话再入stack1（按照顺序先左后右），
	//注意这样的话stack1有点像前序遍历，增加了stack2 之后，相当于把前序遍历反转了一下。
	stack1 := NewStack[*Node[T]]()
	stack2 := NewStack[*Node[T]]()

	stack1.Push(root)
	for !stack1.IsEmpty() {
		p := stack1.Pop()
		stack2.Push(p)
		if p.left != nil {
			stack1.Push(p.left)
		}
		if p.right != nil {
			stack1.Push(p.right)
		}
	}
	for !stack2.IsEmpty() {
		result = append(result, stack2.Pop().value)
	}
	return result
}
