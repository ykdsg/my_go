package parctice1

// 这个解法比较有意思，index 是非0值下标，一路填充过来，循环完之后将index到len(nums)填充为0
// 胜在思路清晰简单，但是效率没有交换最左0值下标的方案高。
func moveZeroes2(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index += 1
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index += 1
	}
}
