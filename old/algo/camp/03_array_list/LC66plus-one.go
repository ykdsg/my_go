package _3_array_list

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			//如果尾部是9说明进位了，继续循环
			digits[i] = 0
		} else {
			//说明不用进位，在这里直接+1 返回
			digits[i] = digits[i] + 1
			return digits
		}
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
