// Package main 二分查找
package main

import (
	"fmt"
	"sort"
)

// BinarySearchFor 二分查找循环实现
func BinarySearchFor(sli []int, k int) int {
	// 低位高位 长度
	lo, hi := 0, len(sli)-1
	// 如果低位小于等于高位一直循环
	for lo <= hi {
		// 取中间位，向右移动一位等于除以 2
		m := (lo + hi) >> 1 // 等同于 (lo + hi) / 2
		if sli[m] < k {
			lo = m + 1
		} else if sli[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}

	return -1 // 没找到是 -1
}

// BinarySearchRecursion 二分查找递归实现
func BinarySearchRecursion(sli *[]int, leftIndex int, rightIndex int, findVal int) {
	// 退出条件，递归必须写退出条件，不然栈会压爆
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}

	middleIndex := (leftIndex + rightIndex) / 2
	// 如果值大于中间索引处的值
	if findVal > (*sli)[middleIndex] {
		// 将索引 +1
		BinarySearchRecursion(sli, middleIndex+1, rightIndex, findVal)
	} else if findVal < (*sli)[middleIndex] {
		BinarySearchRecursion(sli, leftIndex, middleIndex-1, findVal)
	} else {
		// 如果相等，则就是这个中间索引
		fmt.Printf("找到了，下标为: %v\n", middleIndex)
	}
}

func main() {
	// 二分查找前提：必须是有序
	sli := []int{8, 7, 6, 5, 4, 3, 2, 1}
	// 先排序
	sort.Ints(sli)
	fmt.Printf("sli: %v\n", sli) // 1, 2, 3, 4, 5, 6, 7, 8

	result := BinarySearchFor(sli, 5) // 返回下标数
	fmt.Println(result)

	BinarySearchRecursion(&sli, 0, len(sli), 8)
}
