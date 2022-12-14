package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	ast := assert.New(t)
	ast.Equal([]int{1, 2, 2, 3, 5, 6}, nums1)
}
