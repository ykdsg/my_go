package parctice1

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

	nums = []int{3, 0, -2, -1, 1, 2}
	result = threeSum(nums)
	ast.Equal([][]int{[]int{-2, -1, 3}, []int{-2, 0, 2}, []int{-1, 0, 1}}, result)

	nums = []int{0, 0, 0}
	result = threeSum(nums)
	ast.Equal([][]int{[]int{0, 0, 0}}, result)
}

func TestThreeSum3(t *testing.T) {
	ast := assert.New(t)

	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum3(nums)
	ast.Equal([][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}, result)
}
