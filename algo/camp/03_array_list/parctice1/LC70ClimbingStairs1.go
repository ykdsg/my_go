package parctice1

// 使用移动窗口的方式，比递归能大大提升执行效率
func climbStairs1(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	int2, int3 := 1, 2
	for i := 2; i < n; i++ {
		int2, int3 = int3, int2+int3
	}

	return int3
}

// 这种就是更简洁的写法
func climbStairs2(n int) int {
	int1, int2 := 1, 1
	for i := 0; i < n; i++ {
		int1, int2 = int2, int1+int2
	}
	return int1
}
