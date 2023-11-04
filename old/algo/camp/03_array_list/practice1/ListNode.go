package practice1

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(v int) *ListNode {
	return &ListNode{
		Val: v,
	}

}

func NewHead(list []int) *ListNode {
	if list == nil {
		return nil
	}
	head := newListNode(list[0])
	pre := head
	for i := 1; i < len(list); i++ {
		cur := newListNode(list[i])
		pre.Next = cur
		pre = cur
	}
	return head
}

func toList(node *ListNode) []int {
	if node == nil {
		return nil
	}
	cur := node
	result := []int{}
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}
