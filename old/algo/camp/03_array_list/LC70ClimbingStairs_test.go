package _3_array_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, climbStairs(2))
	ast.Equal(3, climbStairs(3))
	ast.Equal(5, climbStairs(4))
	ast.Equal(8, climbStairs(5))

	count := climbStairs(44)
	fmt.Println(count)
}

func TestClimbStairs2(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, climbStairs2(2))
	ast.Equal(3, climbStairs2(3))
	ast.Equal(5, climbStairs2(4))
	ast.Equal(8, climbStairs2(5))

	count := climbStairs2(44)
	fmt.Println(count)
}

func TestClimbStairs3(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(2, climbStairs3(2))
	ast.Equal(3, climbStairs3(3))
	ast.Equal(8, climbStairs3(5))
	ast.Equal(5, climbStairs3(4))

	count := climbStairs3(44)
	fmt.Println(count)
}
