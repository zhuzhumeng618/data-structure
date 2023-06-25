// Package main 栈 后进先出
package main

import (
	"fmt"
)

// Stack 使用字符串切片存储数据
type Stack []string

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push 入栈、压栈
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop 出栈、弹栈
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index] // 获取栈顶元素
		*s = (*s)[:index]      // 删除元素
		return element, true   // 返回
	}
}

func main() {
	var stack Stack // 创建一个栈

	stack.Push("java")
	stack.Push("python")
	stack.Push("golang")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}

}
