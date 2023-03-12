package practice1

func moveZeroes1(nums []int) {
	//表示最左边0的位置
	mostLeftZeroIndex := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			//如果是第一次遇到0，初始化zeroIndex 下标
			if mostLeftZeroIndex < 0 {
				mostLeftZeroIndex = i
			}
		} else {
			//如果不是0，且zeroIndex 已经初始化，只会存在2种情况
			//1.mostLeftZeroIndex 跟当前index紧挨着
			//2.zeroIdex 跟当前index 中间隔着连续的0
			//因此只要跟zeroIndex 交换位置，并且zeroIndex 右移一位
			if mostLeftZeroIndex >= 0 {
				nums[mostLeftZeroIndex], nums[i] = nums[i], 0
				mostLeftZeroIndex += 1
			}
		}
	}
}

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
