package practice1

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	v := &ListNode{}
	pre := v
	h1, h2 := list1, list2
	for h1 != nil || h2 != nil {
		if h1 == nil {
			pre.Next = h2
			break
		}
		if h2 == nil {
			pre.Next = h1
			break
		}
		if h1.Val < h2.Val {
			pre.Next = h1
			h1 = h1.Next
		} else {
			pre.Next = h2
			h2 = h2.Next
		}
		pre = pre.Next
	}
	return v.Next
}
