package main

import (
	"fmt"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7}
	fmt.Println(arr)
	SelectionSort(arr)
	fmt.Println(arr)
	fmt.Println("")
}

// 选择排序(Selection Sort)
// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
// 然后再从剩余未排序元素中继续寻找最小（大）元素，放到已排序序列的末尾
func SelectionSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		tmp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = tmp
	}

	return arr
}
