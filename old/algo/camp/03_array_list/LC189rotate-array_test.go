package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 3)
	ast := assert.New(t)
	ast.Equal([]int{5, 6, 7, 1, 2, 3, 4}, nums)
}

func TestRotate3(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	rotate3(nums, 2)
	ast := assert.New(t)
	ast.Equal([]int{3, 99, -1, -100}, nums)
}
