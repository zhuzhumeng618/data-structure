// Package main 选择排序
package main

import (
	"fmt"
)

func SelectSort(values []int) {
	l := len(values)
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if values[j] < values[min] {
				min = j
			}
		}

		// 交换值
		tmp := values[i]
		values[i] = values[min]
		values[min] = tmp
	}
}

func main() {
	sli := []int{5, 4, 3, 2, 1}
	SelectSort(sli)
	fmt.Println("升序排序后:", sli)
}
