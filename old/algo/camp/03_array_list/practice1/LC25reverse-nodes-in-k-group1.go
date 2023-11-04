package practice1

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// 核心关键是通过k次循环找到第k个节点，然后将node.Next=nil，这样就可以对start 进行翻转，翻转之后再接上下一节的head
func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	v := &ListNode{
		Next: head,
	}

	//注意这里pre 和 cur 指向同一位置
	pre := v
	cur := pre
	for cur.Next != nil {
		for i := 0; i < k && cur != nil; i++ {
			cur = cur.Next
		}
		//说明不足个数
		if cur == nil {
			break
		}
		nextGroupHead := cur.Next
		cur.Next = nil
		start := pre.Next
		//翻转整个链表
		pre.Next = reverse1(start)
		start.Next = nextGroupHead

		pre = start
		cur = pre
	}

	return v.Next
}

func reverse1(head *ListNode) *ListNode {
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
