package main

import (
	"fmt"

	"../../comics_algorithm/chapter_4"
)

func main() {
	var arr []int

	arr = []int{9, 6, 2, 5, 4, 8, 7, 1, 3, 6, 7, 8, 4, 1}
	fmt.Println(arr)
	result := BucketSort(arr)
	fmt.Println(result)
	fmt.Println("")

	arr = []int{99, 96, 92, 95, 94, 98, 97, 91, 93, 96, 97, 98, 94, 91}
	fmt.Println(arr)
	result = BucketSort(arr)
	fmt.Println(result)
	fmt.Println("")

}

// 桶排序 (Bucket Sort)  O(n+k)
// 桶排序的基本思路是将数据根据计算规则来分组，并将数据依次分配到对应分组中。
// 分配时可能出现某分组里有多个数据，那么再进行分组内排序。
// 然后得到了有序分组，最后把它们依次取出来放到数组中即实现了整体排序。
func BucketSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	//求最大值、桶数量
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	num := len(arr)

	// 创建桶
	buckets := make([][]int, num)

	for i := 0; i < num; i++ {
		index := arr[i] * (num - 1) / max //分配桶

		buckets[index] = append(buckets[index], arr[i])
	}

	//分桶排序并拷贝
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			copy(arr[tmpPos:], chapter_4.QuickSort(buckets[i]))

			tmpPos += bucketLen
		}
	}

	return arr
}
