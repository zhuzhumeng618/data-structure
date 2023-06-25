// Package main 链表
package main

import (
	"fmt"
)

type Object interface{}

type Node struct {
	Data Object // 定义数据域
	Next *Node  // 定义地址域（指向下一个表的地址）
}

type List struct {
	headNode *Node // 头节点
}

// IsEmpty 判断链表是否为空，如果头节点为空，则列表为空
func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	}
	return false
}

// GetLength 获取列表长度
func (l *List) GetLength() int {
	current := l.headNode
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	return len
}

// Add 从链表头部添加元素
func (l *List) Add(obj Object) *Node {
	node := &Node{
		Data: obj,
	}
	node.Next = l.headNode
	l.headNode = node
	return node
}

// Append 从链表尾部添加元素
func (l *List) Append(obj Object) {
	node := &Node{
		Data: obj,
	}
	if l.IsEmpty() {
		l.headNode = node
	} else {
		current := l.headNode
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// Insert 在链表指定位置添加元素
func (l *List) Insert(index int, obj Object) {
	if index < 0 {
		l.Add(obj)
	} else if index > l.GetLength() {
		l.Append(obj)
	} else {
		pre := l.headNode
		len := 0
		for len < (index - 1) {
			pre = pre.Next
			len++
		}
		node := &Node{Data: obj}
		node.Next = pre.Next
		pre.Next = node
	}
}

// Remove 删除链表指定值的元素
func (l *List) Remove(obj Object) {
	pre := l.headNode
	if pre.Data == obj {
		l.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == obj {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// RemoveAtIndex 删除指定位置的元素
func (l *List) RemoveAtIndex(index int) {
	pre := l.headNode
	if index <= 0 {
		l.headNode = pre.Next
	} else if index > l.GetLength() {
		fmt.Println("超长了")
		return
	} else {
		len := 0
		for len != (index-1) && pre.Next != nil {
			len++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// Contains 查看链表是否包含某个元素
func (l *List) Contains(obj Object) bool {
	current := l.headNode
	for current != nil {
		if current.Data == obj {
			return true
		}
		current = current.Next
	}
	return false
}

// LoopList 遍历链表所有节点
func (l *List) LoopList() {
	if !l.IsEmpty() {
		current := l.headNode
		for {
			fmt.Printf("current.Data: %v\n", current.Data)
			if current.Next != nil {
				current = current.Next
			} else {
				break
			}
		}
	}
}

func main() {
	list := List{}
	list.Append("Java")
	list.Append("Python")
	list.Append("Golang")
	fmt.Printf("list.GetLength(): %v\n", list.GetLength())
	list.LoopList()

	fmt.Println("------------------------------")
	list.Add("before")
	list.LoopList()

	fmt.Println("------------------------------")
	list.Remove("Java")
	list.LoopList()

	fmt.Println("------------------------------")
	list.RemoveAtIndex(2)
	list.LoopList()
}
