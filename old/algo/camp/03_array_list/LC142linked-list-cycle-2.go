package _3_array_list

//给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

//整体分析步骤比较复杂，总的来说分为2部分
//首先讨论的都是有环的情况下，假设设链表共有 a+b 个节点，其中 链表头部到链表入口 有 a 个节点， 链表环有b个节点
//第一部分：
//假设fast 和slow指针都指向head，fast每次走2步，slow每次走1步
//当fast 和slow 相遇的时候，可以得出：
//1.fast 走的步数是slow步数的 2 倍，即 f = 2s ；
//2.fast 比slow 多走n个环长度，因为双指针都走过a，然后在环内重合的时候，fast 比slow多走n环，f=s+nb
//因此得出：f=2nb ,s=nb

//第二部分：
//如果让指针从链表头部一直向前走并统计步数k，那么所有 走到链表入口节点时的步数 是：k=a+nb （因为进入环之后可以循环）
//根据第一部分分析得出的 s=nb ，也就是再走a步就会停留在入口节点。
//因此在第一次fast和slow 相遇之后，再派一个指针从head 开始走a步 就会跟slow 指针重合。那就是第一次重合之后另一个指针从head 开始一次一步，slow也一次一步，直到重合就走了a步。
//那就能算出a是多少

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	//	这样就可以在统一起跑线出发

	for true {
		//没有环的情况
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		// 指针第一次重叠
		if fast == slow {
			break
		}
	}

	//这里直接复用fast为新指针
	fast = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

// 用map来实现，整体逻辑就比较简单，但是空间复杂度就没有上面的那么优秀。
func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	cache := map[*ListNode]bool{}
	p := head
	for p != nil {
		if _, ok := cache[p]; ok == true {
			return p
		}
		cache[p] = true
		p = p.Next
	}
	return nil
}
