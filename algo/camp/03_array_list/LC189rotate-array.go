package _3_array_list

// 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
// 这个性能太差基本不太可行，还是额外使用一个数组，直接根据最终位置进行拷贝性能比较好，但是会额外还费内存空间
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		pre := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			cur := nums[j]
			nums[j] = pre
			pre = cur
		}
	}
}

// 这里理解的难点在于判断什么时候整体的处理结束
// 由于最终回到了起点，故该过程恰好走了整数数量的圈，不妨设为 a 圈；再设该过程总共遍历了 b 个元素。因此，我们有an=bk
func rotate2(nums []int, k int) {
	leng := len(nums)
	k %= leng
	for start, count := 0, gcd(k, leng); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % leng
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

// 最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 通过变量count累加的方式，表示处理的结束，这个方式比上一种更简单直接好理解
func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, 0; count < len(nums); start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
			count++
		}

	}
}

// 通过数据翻转的方法，可以参考下面的解释
// nums = "----->-->"; k =3
// result = "-->----->";
//
// reverse "----->-->" we can get "<--<-----"
// reverse "<--" we can get "--><-----"
// reverse "<-----" we can get "-->----->"
func rotate4(nums []int, k int) {
	k %= len(nums)
	reverseArr(nums)
	reverseArr(nums[:k])
	reverseArr(nums[k:])
}

func reverseArr(nums []int) {
	n := len(nums) - 1
	if n <= 0 {
		return
	}
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i] = nums[n-i], nums[i]
	}

}
