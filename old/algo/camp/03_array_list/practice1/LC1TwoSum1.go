package practice1

func twoSum1(nums []int, target int) []int {
	valueIndexMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		curValue := nums[i]
		nextValue := target - curValue
		if nextIndex, ok := valueIndexMap[nextValue]; ok {
			return []int{nextIndex, i}
		} else {
			valueIndexMap[curValue] = i
		}
	}
	return nil

}
