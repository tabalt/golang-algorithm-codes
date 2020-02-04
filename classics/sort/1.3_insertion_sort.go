package main

import (
	"fmt"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)
	fmt.Println("")
}

// 插入排序（Insertion Sort） O(n^2)
// 通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
// 需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
func InsertionSort(arr []int) {
	current := 0
	for i := 0; i < len(arr)-1; i++ {
		current = arr[i+1]
		preIndex := i
		for preIndex >= 0 && current < arr[preIndex] {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}

		arr[preIndex+1] = current
	}
}
