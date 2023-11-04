package practice1

// 使用两端靠近的方法，关键点是左右两端选择更短的那一根进行逼近，同时优化了无谓计算那部分
func maxArea2(height []int) int {
	left, right, max := 0, len(height)-1, 0
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left = nextLeft(left, right, height)
		} else {
			right = nextRight(left, right, height)
		}
	}

	return max
}

// 避免无谓的计算，因为只有长度比left 高，面积才有可能更大
func nextLeft(left int, right int, height []int) int {
	cur := height[left]
	for left < right {
		left = left + 1
		if height[left] > cur {
			return left
		}
	}
	return right
}

func nextRight(left int, right int, height []int) int {
	cur := height[right]
	for left < right {
		right = right - 1
		if height[right] > cur {
			return right
		}
	}
	return left
}
