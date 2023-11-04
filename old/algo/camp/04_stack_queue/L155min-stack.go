package _4_stack_queue

import "math"

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	minTop := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(val, minTop))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
