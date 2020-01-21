package chapter_4

//快速排序，分治法：每一轮挑选一个基准元素，并让比它大的元素移动到数列的一边，比它小的袁术移动到另一边
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
