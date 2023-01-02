package parctice1

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	v := &ListNode{
		Next: head,
	}
	pre := v
	cur := pre
	for cur != nil {
		//找到第k个节点
		for i := 0; i < k && cur != nil; i++ {
			cur = cur.Next
		}
		//不到k个直接跳出
		if cur == nil {
			break
		}
		nextGroopHead := cur.Next
		//从这掐断
		cur.Next = nil
		start := pre.Next
		pre.Next = reverse2(start)
		//这个时候start 翻转后到了队尾的位置，需要再续上
		start.Next = nextGroopHead

		pre = start
		cur = pre

	}

	return v.Next
}

func reverse2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre

}

func reverseKGroup2_recursive(head *ListNode, k int) *ListNode {
	index := head
	count := 0
	for index != nil && count < k {
		index = index.Next
		count++
	}
	//注意这里的区别，这里需要通过计数的方式，因为就算cur是nil，但是count 的数量到了，后面的逻辑还是要运行的
	if count != k {
		return head
	}

	nextGroopHead := reverseKGroup2_recursive(index, k)

	cur := head
	next := nextGroopHead
	for i := 0; i < k; i++ {
		oldNext := cur.Next
		cur.Next = next
		next = cur
		cur = oldNext
	}

	//注意这里返回的是next ，因为cur 在最后一个循环别设置为oldNext，这个就不对了。
	return next
}
