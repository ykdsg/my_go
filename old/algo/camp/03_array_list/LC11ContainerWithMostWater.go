package _3_array_list

func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}
	left := 0
	right := len(height) - 1
	maxV := 0
	for left < right {
		w := right - left
		h, isLeft := min(height[left], height[right])
		maxV = max(maxV, w*h)
		if isLeft {
			left = moreHighIndex(left, right, isLeft, height)
		} else {
			right = moreHighIndex(left, right, isLeft, height)
		}
	}

	return maxV
}

// 左右逼近的方式查询下一个index
func moreHighIndex(left int, right int, isLeft bool, height []int) int {
	if isLeft {
		for left < right {
			if height[left] < height[left+1] {
				return left + 1
			} else {
				left += 1
			}
		}
	} else {
		for left < right {
			if height[right] < height[right-1] {
				return right - 1
			} else {
				right -= 1
			}
		}
	}

	return -1
}

func min(x, y int) (min int, isLeft bool) {
	if x < y {
		return x, true
	}
	return y, false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
