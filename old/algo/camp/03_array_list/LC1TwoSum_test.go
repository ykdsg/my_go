package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	ast := assert.New(t)
	result := twoSum3([]int{2, 7, 11, 15}, 9)
	ast.Equal([]int{0, 1}, result)

}
