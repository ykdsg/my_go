package parctice1

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//在head前面虚拟一个pre，能够统一处理相应情况，因为第二组的时候就会涉及前一个pre 指针的处理
	pre := &ListNode{
		Next: head,
	}
	cur := pre
	for cur.Next != nil && cur.Next.Next != nil {
		next1 := cur.Next
		next2 := next1.Next

		cur.Next = next2
		next1.Next = next2.Next
		next2.Next = next1

		cur = next1
	}
	return pre.Next
}

// 整体逻辑跟上面一致，但是代码组织的更简洁
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//通过虚拟的一个节点
	v := &ListNode{
		Next: head,
	}
	pre, cur := v, head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next = next
		cur.Next, next.Next = next.Next, cur
		pre, cur = cur, cur.Next
	}

	return v.Next
}

// 递归的解法就更简洁
func swapPairsRecursive1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsRecursive1(next.Next)
	next.Next = head

	return next
}
