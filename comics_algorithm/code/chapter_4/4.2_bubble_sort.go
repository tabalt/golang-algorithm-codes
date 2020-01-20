package chapter_4

//冒泡排序
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

//冒泡排序 优化版，有序后提前退出循环
func BubbleSortBreakSorted(arr []int) {
	for i := 0; i < len(arr); i++ {
		sorted := true
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp

				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}

//冒泡排序 更优版，有序后提前退出循环、增加有序区界定
func BubbleSortBreakSortedAndBorder(arr []int) {
	lastExchangeIndex := 0
	sortBorder := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		sorted := true
		for j := 0; j < sortBorder; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp

				sorted = false
				lastExchangeIndex = j
			}
		}
		sortBorder = lastExchangeIndex
		if sorted {
			break
		}
	}
}
