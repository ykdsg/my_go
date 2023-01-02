package parctice1

// 整体逻辑比较简单，但是时间复杂度有点高，k*n
func rotate(nums []int, k int) {
	leng := len(nums)
	n := k % leng
	for i := 0; i < n; i++ {
		pre := nums[leng-1]
		for j := 0; j < leng; j++ {
			cur := nums[j]
			nums[j] = pre
			pre = cur
		}
	}
}

func rotate2(nums []int, k int) {
	leng := len(nums)
	n := k % leng
	for start, count := 0, 0; count < leng; start++ {
		preIndex := start
		pre := nums[preIndex]
		//相当于do while(preIndex!= start)
		//这里搞这么复杂是因为有些情况下len 和k 刚好会形成反复循环，比如len=4，k=2 ，就会在0，2之间反复，这个时候就需要跳出循环，重新下一个索引
		for ok := true; ok; ok = preIndex != start {
			curIndex := (preIndex + n) % leng
			cur := nums[curIndex]
			nums[curIndex] = pre
			pre = cur
			preIndex = curIndex
			count++
		}
	}
}
