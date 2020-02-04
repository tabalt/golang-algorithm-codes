package main

import (
	"fmt"

	"../../comics_algorithm/chapter_3"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7, 1}
	fmt.Println(arr)
	HeapSort(arr)
	fmt.Println(arr)
	fmt.Println("")

}

// 堆排序 (Heap Sort)  O(nlogn)
// 最大堆排出从小到大，最小堆排出从大到小的序列
// 时间复杂度 O(n) + O(nlogn) = O(nlogn)
func HeapSort(arr []int) {
	bh := &chapter_3.BinaryHeap{}
	//构建二叉堆
	bh.Build(arr)

	//循环删除堆顶元素，移动到未排序的尾部，调整产生新对顶
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = arr[0]
		arr[0] = tmp

		bh.DownAdjustParent(arr, 0, i)
	}
}
