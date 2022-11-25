package _3_array_list

//解题思路：
//1.暴力冒泡的方法，非0的看下前面是不是0，如果是就冒泡直到遇到非0
//2.优化遇到连续0的情况，不要傻傻的冒泡，直接跟最前面的0交换即可，只是需要多维护一个zeroIndex变量

// 关键是zeroIndex 的处理，为了提升效率
func moveZeroes(nums []int) {
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zeroIndex < 0 {
				zeroIndex = i
			}
		} else {
			if zeroIndex >= 0 {
				nums[zeroIndex] = nums[i]
				nums[i] = 0
				zeroIndex += 1
			}
		}
	}
}
