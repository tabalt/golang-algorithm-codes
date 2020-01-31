package main

import (
	"fmt"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7, 1}
	fmt.Println(arr)
	result := QuickSort(arr)
	fmt.Println(result)
	fmt.Println("")

	arr = []int{9, 6, 2, 5, 4, 8, 7, 1}
	fmt.Println(arr)
	QuickSort2(arr, 0, len(arr)-1)
	fmt.Println(arr)
	fmt.Println("")

}

// 快速排序
// 分治法：每一轮挑选一个基准元素，并让比它大的元素移动到数列的一边，比它小的元素移动到另一边
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left := []int{}
	right := []int{}
	mid := arr[0]
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		if v <= mid {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	result := []int{}
	result = append(result, QuickSort(left)...)
	result = append(result, mid)
	result = append(result, QuickSort(right)...)

	return result
}

func QuickSort2(arr []int, left, right int) {
	if left < right {
		mid := partition(arr, left, right)
		QuickSort2(arr, left, mid-1)
		QuickSort2(arr, mid+1, right)
	}
}

func partition(arr []int, left, right int) int {
	//选择中轴元素
	pivot := arr[left]
	i, j := left+1, right
	for {
		//从左向右找到第一个大于pivot的元素位置
		for i <= j && arr[i] <= pivot {
			i++
		}
		//从右向左找到第一个小于pivot的元素位置
		for i <= j && arr[j] >= pivot {
			j--
		}

		if i >= j {
			break
		}

		//交换两个元素的位置，使得左边元素不大于pivot，右边元素不小于pivot
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	}

	//使中轴位置处于有序位置
	arr[left] = arr[j]
	arr[j] = pivot

	return j
}
