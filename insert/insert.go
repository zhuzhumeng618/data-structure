// Package main 插入排序
package main

import (
	"fmt"
)

func InsertSort(values []int) {
	l := len(values)
	// 总循环次数，元素个数
	for i := 0; i < l; i++ {
		// 从后面第二个开始，取出，找到合适的位置插入
		for j := i; j > 0; j-- {
			// 比较
			if values[j] > values[j-1] {
				// 交换
				tmp := values[j]
				values[j] = values[j-1]
				values[j-1] = tmp
			}
		}
	}
}

func main() {
	sli := []int{1, 2, 3, 4, 5, 6}
	InsertSort(sli)
	fmt.Println("降序排序后:", sli)
}
