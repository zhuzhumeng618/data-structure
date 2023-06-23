// Package main 快速排序
package main

import (
	"fmt"
)

func QuickSort(values []int) []int {
	l := len(values)

	if l < 1 {
		return values
	}

	// 假设第一个元素是中间值
	middle := values[0]

	// 左边元素
	var left []int

	// 右边元素
	var right []int

	for i := 1; i < l; i++ {
		if middle < values[i] {
			right = append(right, []int{values[i]}...)
		} else {
			left = append(left, []int{values[i]}...)
		}
	}

	middle_s := []int{middle}
	// 递归排左边
	left = QuickSort(left)

	// 递归排右边
	right = QuickSort(right)

	result := append(append(left, middle_s...), right...)

	return result
}

func main() {
	sli := []int{5, 4, 3, 2, 1}
	sli = QuickSort(sli)
	fmt.Println("升序排序后:", sli)
}
