package _3_array_list

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	cache := map[int]int{}
	for i := 0; i < len(nums); i++ {
		cache[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		i2, ok := cache[t]
		if !ok || i == i2 {
			continue
		}
		return []int{i, i2}
	}
	return nil
}

// 这里更巧妙的是一遍循环就能实现，而且避免了自己跟自己配对的问题
func twoSum3(nums []int, target int) []int {
	cache := map[int]int{}
	for i, v := range nums {
		t := target - v
		if p, ok := cache[t]; ok {
			return []int{p, i}
		} else {
			cache[v] = i
		}
	}
	return nil
}
