package practice1

// 给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
func merge(nums1 []int, m int, nums2 []int, n int) {
	//核心关键是从队尾开始排
	cur := len(nums1) - 1
	n1 := m - 1
	n2 := n - 1
	for n2 >= 0 {
		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[cur] = nums1[n1]
			n1--
		} else {
			nums1[cur] = nums2[n2]
			n2--
		}
		cur--
	}
}
