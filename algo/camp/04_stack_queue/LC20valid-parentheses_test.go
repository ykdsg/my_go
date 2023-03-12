package _4_stack_queue

import (
	"camp/04_stack_queue/practice1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	ast := assert.New(t)

	valid := isValid("{[]}")
	ast.True(valid == true)

	valid1 := practice1.IsValid1("{[]}")
	ast.True(valid1 == true)

}
