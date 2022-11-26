package parctice1

import "sort"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0
// 关键点是对重复数据的处理，第一个是在锚定值的选择上，第二个是在left 的选择上，都要对重复值左兼容
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	result := [][]int{}
	sort.Ints(nums)
	for i, num := range nums {
		if num > 0 {
			return result
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		right := len(nums) - 1
		for left := i + 1; left < right; left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}
			//右指针开始逼近
			for left < right && num+nums[left]+nums[right] > 0 {
				right--
			}
			if left == right {
				break
			}
			if num+nums[left]+nums[right] == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
			}
		}
	}
	return result
}
