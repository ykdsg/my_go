package _3_array_list

//给你链表的头节点 head ，每k个节点一组进行翻转，请你返回修改后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。

// 使用递归的方法还是挺简单的
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nodes := []*ListNode{}
	p := head
	for i := 0; i < k; i++ {
		//说明长度不够k，不需要转换
		if p == nil {
			return head
		}
		nodes = append(nodes, p)
		p = p.Next
	}
	nextGroupHead := nodes[k-1].Next
	for i := k - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = reverseKGroup(nextGroupHead, k)

	return nodes[k-1]
}

// 非递归写法，这个比较有技巧，
// 虚拟一个head 前的节点dummy 这样就能统一地处理所有操作，而且dummy.Next 就是新链表的head。
// 通过end 的指针巧妙地标记当前段最后一位，end.Next就是下一group 的head 指针。同时姜end.Next 暂时设置为nil，相当于截断了链表，这个时候返转start相当于只反转了这段group的链表。
func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//创建一个在head 前的虚拟节点
	dummy := &ListNode{
		Val: 0,
	}
	dummy.Next = head

	pre := dummy
	end := pre

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 说明最后一组的数量不足k
		if end == nil {
			break
		}
		//下一组的head
		nextGroupHead := end.Next
		//这是为了翻转函数好判断什么时候结束
		end.Next = nil
		start := pre.Next
		pre.Next = reverse(start)
		start.Next = nextGroupHead
		//下一组翻转的起始坐标
		pre = start
		end = pre

	}

	return dummy.Next
}

// 反转整条链表
func reverse(head *ListNode) *ListNode {
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
