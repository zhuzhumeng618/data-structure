// Package main 归并排序
package main

import (
	"fmt"
)

func MergeSort(values []int, a int, b int) {
	// 长度小于等于 1，不用排序
	if b-a <= 1 {
		return
	}

	// 取中间值
	c := (a + b) / 2

	// 递归调用
	MergeSort(values, a, c) // 左边部分排序
	MergeSort(values, c, b) // 右边部分排序

	leftSli := make([]int, c-a)
	rightSli := make([]int, b-c)

	copy(leftSli, values[a:c])
	copy(rightSli, values[c:b])

	i := 0
	j := 0
	for k := a; k < b; k++ {
		switch {
		case i >= c-a:
			values[k] = rightSli[j]
			j++
		case j >= b-c:
			values[k] = leftSli[i]
			i++
		case leftSli[i] < rightSli[i]:
			values[k] = leftSli[i]
			i++
		default:
			values[k] = rightSli[j]
			j++
		}
	}
}

func main() {
	sli := []int{5, 4, 3, 2, 1}
	MergeSort(sli, 0, len(sli))
	fmt.Println("升序排序后:", sli)
}
