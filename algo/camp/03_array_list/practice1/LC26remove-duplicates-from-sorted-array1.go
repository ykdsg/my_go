package practice1

// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
// 因为不需要考虑超出新长度之后的元素，所以相对还是简单的。
func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}
