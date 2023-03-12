package practice1

// 迭代版本
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	//在遍历链表的时候相当于把next指针改为pre指针，因为结构体没有持有pre指针，所以需要使用pre变量存储前一个指针
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 递归版本
func reverseList2(head *ListNode) *ListNode {
	return doReverseList(head, nil)
}

func doReverseList(cur *ListNode, pre *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return doReverseList(next, cur)
}
