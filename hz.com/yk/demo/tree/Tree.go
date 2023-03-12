package main

import (
	"container/list"
	"fmt"
)

// 接口，关于树的操作
type tree interface {
	do()
}
type node struct {
	*list.List // (匿名字段)组合即继承，node 拥有list的特性
	name       string
}

// 叶子节点
type leaf struct {
	name string
}

func (l leaf) do() {
	fmt.Println(l.name + " leaf do something.")
}

func (n node) do() { // 定义即实现，node 实现了tree接口
	for e := n.Front(); e != nil; e = e.Next() { // node 拥有list特性
		e.Value.(tree).do()
	}
	fmt.Println(n.name + " node do something.")
}

func (n node) addSub(sub tree) {
	n.PushBack(sub)
}

func main() {
	n1 := node{list.New(), "n1"}
	n2 := node{list.New(), "n2"}
	l1 := leaf{"l1"}
	l2 := leaf{"l2"}

	n2.addSub(l2)
	n1.addSub(n2)
	n1.addSub(l1)

	n1.do()
}
