package _3_array_list

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
	fast := head.Next
	//	因为for 循环的关系，slow 和fast 不能在同一起跑线
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
