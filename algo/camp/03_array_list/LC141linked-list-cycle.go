package _3_array_list

// 给你一个链表的头节点 head ，判断链表中是否有环。
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	//	因为for 循环的关系，slow 和fast 不能在同一起跑线，可以假设前面有个虚拟的开始节点
	for slow != fast {
		//因为fast 跑在前面，只要fast 不为nil，slow就不会是nil
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	//	这样就可以在统一起跑线出发
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
