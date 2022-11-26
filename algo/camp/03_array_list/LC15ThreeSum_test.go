package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	ast := assert.New(t)

	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	ast.Equal([][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}, result)

	nums = []int{0, 1, 1}
	result = threeSum(nums)
	ast.Equal([][]int{}, result)

	nums = []int{0, 0, 0}
	ast.Equal([][]int{[]int{0, 0, 0}}, threeSum(nums))
}
