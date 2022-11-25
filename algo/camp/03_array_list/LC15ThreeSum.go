package _3_array_list

import "sort"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	result := [][]int{}
	//需要基于有序数组
	sort.Ints(nums)
	//相当于先锚定最左边的数，然后利用左右指针收拢的方式来找到可能的组合
	for i, num := range nums {
		if num > 0 {
			return result
		}
		//如果num 跟前一位相等需要跳过，因为已经遍历过该值的所有可能情况
		if i > 0 && nums[i-1] == num {
			continue
		}
		right := len(nums) - 1
		for left := i + 1; left < right; left++ {
			//如果跟上一次的值相同需要跳过
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}
			//如果>0 说明整体偏大，right左移，知道找到小于等于的数
			for left < right && num+nums[left]+nums[right] > 0 {
				right--
			}
			//左右指针重合说明没找到合适的组合
			if left == right {
				break
			}
			//说明知道了匹配的值
			if num+nums[left]+nums[right] == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
			}
		}
	}
	return result
}
