package chapter_4

import (
	"../utils"
	"testing"
)

var (
	quickSortInputCases = [][][]int{
		{
			{9, 6, 2, 5, 4, 8, 7, 1},
			{1, 2, 4, 5, 6, 7, 8, 9},
		},
		{
			{3, 2, 1, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
)

func Test_Quick_Sort_All(t *testing.T) {

	for idx, inputs := range quickSortInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		var arr []int

		arr = make([]int, len(inputs[0]))
		copy(arr, inputs[0])
		arr = QuickSort(arr)
		// t.Logf("sort result\t%v", arr)
		if !utils.SliceEqual(arr, inputs[1]) {
			t.Errorf("QuickSort error, index: %d, result: %v, expect: %v", idx, arr, inputs[1])
		}

	}

}
