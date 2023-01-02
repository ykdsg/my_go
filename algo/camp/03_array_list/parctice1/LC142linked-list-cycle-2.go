package parctice1

// 代码是简单的，但是推导逻辑有点复杂
// 1.假设有快慢指针：fast(每次移动2步),slow(每次移动1步); 链表有a（头到环之间）+b（链表环）个结点，这道题目核心就是求a是多少。
// 2.相遇的时候移动的步数，f（fast 走过的步长）=2s（slow 走过的步长）。
// 3.相遇的时候f 比 s多n个环长度，f=s+nb , 可以推导出来：s=a+n1b+x, f=a+n2+x ,x 是在环中相遇的位置，f-s=(n2-n1)b
// 4.根据2、3 可以得出 s=nb ,f=2nb
// 5.指针从头一直走到链表入口的步数是：a+nb （这里的n 跟上面的n 都指代未知圈数）
// 6.当前s已经走了nb，如果再走a步 就会到 环入口，而这时候如果新的指针从头开始走a步，也会达到环入口，也就是说新指针会跟s 相遇。
// 7.因此在f和s 第一次相遇之后，再让一个新指针从头一步步走到跟s 相遇，就相当于走了a步。
// 8.再拓展一下，如果想求b 是多少 会更加简单，在f、s相遇之后继续跑到下一次相遇的步长就是b。
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	noCycle := true
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			noCycle = false
			break
		}
	}
	if noCycle {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
