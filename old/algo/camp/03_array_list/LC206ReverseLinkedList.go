package _3_array_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head.Next
	head.Next = nil
	pre := head
	for p != nil {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}

	return pre

}

// 下面写法更简洁
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 递归写法，其实不太容易理解，需要画图才能比较清晰的看到为啥要操作当前节点的下一个节点的Next
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList3(head.Next)
	//当前节点的下一个节点的Next指向当前节点，这一步相当于反转
	head.Next.Next = head
	head.Next = nil
	return ret
}

// 递归写法，这个比较容易理解
func reverseList4(head *ListNode) *ListNode {
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
