package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	ast := assert.New(t)

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(height)
	ast.Equal(49, area)

	height = []int{1, 2}
	area = maxArea(height)
	ast.Equal(1, area)

}
