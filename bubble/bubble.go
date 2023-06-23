// Package main 冒泡排序
package main

import (
	"fmt"
)

// BubbleASort 升序
func BubbleASort(values []int) {
	l := len(values)
	// 有几个数就比较几次
	for i := 0; i < l-1; i++ {
		// 比较总长度，每次减一
		for j := i + 1; j < l; j++ {
			// 如果后面的比前面的小就交换，所以是升序
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

// BubbleZSort 升序
func BubbleZSort(values []int) {
	l := len(values)
	// 有几个数就比较几次
	for i := 0; i < l-1; i++ {
		// 比较总长度，每次减一
		for j := i + 1; j < l; j++ {
			// 如果后面的比前面的小就交换，所以是升序
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func main() {
	// sli := []int{5, 4, 3, 2, 1}
	// BubbleASort(sli)
	// for i, v := range sli {
	// 	fmt.Printf("i: %v, v: %v\n", i, v)
	// }
	// fmt.Println("升序排序后: ", sli)

	sli := []int{1, 2, 3, 4, 5, 6}
	BubbleZSort(sli)
	for i, v := range sli {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
	fmt.Println("降序序排序后: ", sli)
}
