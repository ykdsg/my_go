package parctice1

// 维护非0下标的解法，多了一个判断就直接把0填充上了，不用再一个for循环填充0
func moveZeroes3(nums []int) {
	//非0元素下标
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}

// 这个基本是上面的思路，但是整体表现更简洁
func moveZeroes3_1(nums []int) {
	for i, lastNoZeroIndex := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNoZeroIndex], nums[i] = nums[i], nums[lastNoZeroIndex]
			lastNoZeroIndex++
		}
	}
}
