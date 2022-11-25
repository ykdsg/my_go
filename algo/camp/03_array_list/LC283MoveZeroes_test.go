package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	ast := assert.New(t)

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	ast.Equal([]int{1, 3, 12, 0, 0}, nums)

	nums = []int{0}
	moveZeroes(nums)
	ast.Equal([]int{0}, nums)

}
