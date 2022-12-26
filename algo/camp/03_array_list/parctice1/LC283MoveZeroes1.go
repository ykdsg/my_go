package parctice1

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
