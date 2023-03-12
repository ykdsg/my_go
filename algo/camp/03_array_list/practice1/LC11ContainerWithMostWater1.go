package practice1

// 代码简洁，但是执行效率没有LC11ContainerWithMostWater 的高，因为没有对比自己短的进行跳过，会有一些无谓的计算。
func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		w := min(height[left], height[right])
		h := right - left
		maxArea = max(maxArea, h*w)
		if height[left] < height[right] {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	return maxArea
}

func min(x, y int) (min int) {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
