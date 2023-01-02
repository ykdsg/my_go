package parctice1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	ast := assert.New(t)
	result := twoSum1([]int{3, 2, 4}, 6)
	ast.Equal([]int{1, 2}, result)

}
