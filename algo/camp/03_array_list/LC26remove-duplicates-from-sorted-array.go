package _3_array_list

// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 1 {
		return len(nums)
	}
	preIndex := 1

	for i := 1; i < len(nums); i++ {
		//存在重复数据
		if nums[i] != nums[i-1] {
			nums[preIndex] = nums[i]
			preIndex++
		}
	}
	return preIndex
}
