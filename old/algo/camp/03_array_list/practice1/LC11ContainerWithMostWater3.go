package practice1

// 效率不是最优的，因为直接对left 和right 进行移位操作，没有判断是否比之前还短
func maxArea3(height []int) int {
	left, right, max := 0, len(height)-1, 0
	for left < right {
		w := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		area := w * h
		if max < area {
			max = area
		}
	}
	return max
}
