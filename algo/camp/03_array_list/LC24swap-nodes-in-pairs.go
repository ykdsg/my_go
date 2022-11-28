package _3_array_list

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
// 难点在于对pre 节点Next指针的处理。需要画图多推演几个节点，可以跟下面的swapPairs3 对比，swapPairs3 假设head前有一个虚拟节点，这样整体的处理就不需要考虑特殊情况。
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	head = head.Next
	var pre *ListNode = nil
	for cur != nil {
		next := cur.Next
		//说明是最后一个
		if next == nil {
			break
		}
		next2 := next.Next
		next.Next = cur
		cur.Next = next2
		//因为第一个节点没有pre，但是后续节点都需要操作pre，所以这里判断下
		if pre != nil {
			pre.Next = next
		}
		pre = cur
		cur = next2
	}
	return head
}

// 递归解法，可以画图理解下，head表示当前节点，当前节点的Next 是下一组交换后的head节点。
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}

func swapPairs3(head *ListNode) *ListNode {
	//这里使用一个虚拟的pre来假设head 前面的指针，方便统一处理。
	pre := &ListNode{
		Val: 0,
	}
	pre.Next = head
	cur := pre
	for cur.Next != nil && cur.Next.Next != nil {
		next := cur.Next
		next2 := next.Next
		next3 := next2.Next

		cur.Next = next2
		next2.Next = next
		next.Next = next3

		cur = next
	}
	return pre.Next
}
