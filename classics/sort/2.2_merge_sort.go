package main

import (
	"fmt"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7}
	fmt.Println(arr)
	result := MergeSort(arr)
	fmt.Println(result)
	fmt.Println("")
}

// 归并排序（Merge Sort） O(nlogn)
// 采用分治法（Divide and Conquer）, 将已有序的子序列合并，得到完全有序的序列
// 即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	left := arr[:mid]
	right := arr[mid:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for idx, i, j := 0, 0, 0; idx < len(result); idx++ {
		if i >= len(left) {
			result[idx] = right[j]
			j++
		} else if j >= len(right) {
			result[idx] = left[i]
			i++
		} else if left[i] < right[j] {
			result[idx] = left[i]
			i++
		} else {
			result[idx] = right[j]
			j++
		}
	}

	return result
}
