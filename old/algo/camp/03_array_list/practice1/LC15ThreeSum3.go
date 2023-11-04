package practice1

import "sort"

func threeSum3(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	result := [][]int{}
	//循环数组，找一个基准
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return result
		}
		//去重判断
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left = findNextLeft(left, right, nums)
				right = findNextRight(left, right, nums)
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func findNextLeft(left int, right int, nums []int) int {
	for left < right && nums[left] == nums[left+1] {
		left++
	}
	return left + 1
}

func findNextRight(left int, right int, nums []int) int {
	for left < right && nums[right] == nums[right-1] {
		right--
	}
	return right - 1
}
