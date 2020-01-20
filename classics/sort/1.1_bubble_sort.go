package main

import (
	"fmt"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
	fmt.Println("")
}

//冒泡排序（Bubble Sort）
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}
