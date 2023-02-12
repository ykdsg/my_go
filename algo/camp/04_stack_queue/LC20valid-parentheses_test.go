package _4_stack_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	ast := assert.New(t)

	valid := isValid("{[]}")
	ast.True(valid == true)
}
