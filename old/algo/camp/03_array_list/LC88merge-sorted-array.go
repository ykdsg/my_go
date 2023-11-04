package _3_array_list

func merge(nums1 []int, m int, nums2 []int, n int) {
	cur := len(nums1) - 1
	m--
	n--
	for n >= 0 {
		for m >= 0 && nums1[m] > nums2[n] {
			nums1[cur] = nums1[m]
			m--
			cur--
		}
		nums1[cur] = nums2[n]
		n--
		cur--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	cur := len(nums1) - 1
	m--
	n--
	for n >= 0 {
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[cur] = nums1[m]
			m--
		} else {
			nums1[cur] = nums2[n]
			n--
		}
		cur--
	}
}
