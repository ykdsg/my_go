package _3_array_list

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 直接用递归存在的问题有重复计算问题，性能损耗很大，这个在n很大的时候就很坑
func climbStairs(n int) int {
	result, ok := cache[n]
	if ok {
		return result
	}

	result = climbStairs(n-1) + climbStairs(n-2)
	cache[n] = result
	return result
}

var cache = map[int]int{0: 1, 1: 1}

// 因为最终结果只跟n-1,n-2 相关，所以可以用一个长度为3的滚动数组实现
func climbStairs2(n int) int {
	r, s, t := 0, 0, 1
	for i := 1; i <= n; i++ {
		r = s
		s = t
		t = r + s
	}
	return t

}

// 下面的逻辑跟上面差不多，更简洁，省了一个空间，但是计算量应该更多一点
func climbStairs3(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		b += a
		a = b - a
	}
	return a

}
