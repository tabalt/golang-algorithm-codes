package main

import (
	"fmt"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7}
	fmt.Println(arr)
	ShellSort(arr)
	fmt.Println(arr)
	fmt.Println("")
}

// 希尔排序（Shell Sort） O(nlogn)
// 把记录按一定增量分组，对每组使用直接插入排序算法排序; 随着增量逐渐减少，每组包含的关键词越来越多
// 当增量减至1时，整个文件恰被分成一组，算法便终止
func ShellSort(arr []int) {
	length := len(arr)
	tmp := 0
	gap := length / 2
	for gap > 0 {
		for i := gap; i < length; i++ {
			tmp = arr[i]
			preIndex := i - gap
			for preIndex >= 0 && arr[preIndex] > tmp {
				arr[preIndex+gap] = arr[preIndex]
				preIndex -= gap
			}
			arr[preIndex+gap] = tmp
		}
		gap /= 2
	}
}
