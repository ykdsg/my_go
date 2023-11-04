package _3_array_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	result1 := cache.Get(1)
	ast := assert.New(t)
	ast.Equal(1, result1)

	cache.Put(3, 3)
	result2 := cache.Get(2)
	ast.Equal(-1, result2)

	cache.Put(4, 4)
	result1 = cache.Get(1)
	ast.Equal(-1, result1)

	ast.Equal(3, cache.Get(3))
	ast.Equal(4, cache.Get(4))
}

func TestLRUCache_Put2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 6)
	cache.Get(1)
	cache.Put(1, 5)
	cache.Put(1, 2)
	reslut1 := cache.Get(1)
	ast := assert.New(t)
	ast.Equal(2, reslut1)
	reslut2 := cache.Get(2)
	ast.Equal(6, reslut2)

}
