package practice1

import "sort"

// 对排序数组选中1个数，进行左右逼近的方式
// 这里需要对重复解进行排除，所以每一次
func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	result := [][]int{}
	for i, num := range nums {
		if num > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		right := len(nums) - 1
		left := i + 1
		for left < right {
			if num+nums[left]+nums[right] == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				//判断去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if num+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}

		}
	}
	return result
}
