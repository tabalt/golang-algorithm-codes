package chapter_4

import (
	"../utils"
	"testing"
)

var (
	bubbleSortInputCases = [][][]int{
		{
			{9, 6, 2, 5, 4, 8, 7},
			{2, 4, 5, 6, 7, 8, 9},
		},
		{
			{3, 2, 1, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
)

func Test_Bubble_Sort_All(t *testing.T) {

	for idx, inputs := range bubbleSortInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		var arr []int

		arr = make([]int, len(inputs[0]))
		copy(arr, inputs[0])
		BubbleSort(arr)
		// t.Logf("sort result\t%v", arr)
		if !utils.SliceEqual(arr, inputs[1]) {
			t.Errorf("BubbleSort error, index: %d, result: %v, expect: %v", idx, arr, inputs[1])
		}

		arr = make([]int, len(inputs[0]))
		copy(arr, inputs[0])
		BubbleSortBreakSorted(arr)
		// t.Logf("sort result\t%v", arr)
		if !utils.SliceEqual(arr, inputs[1]) {
			t.Errorf("BubbleSortBreakSorted error, index: %d, result: %v, expect: %v", idx, arr, inputs[1])
		}

		arr = make([]int, len(inputs[0]))
		copy(arr, inputs[0])
		BubbleSortBreakSortedAndBorder(arr)
		// t.Logf("sort result\t%v", arr)
		if !utils.SliceEqual(arr, inputs[1]) {
			t.Errorf("BubbleSortBreakSortedAndBorder error, index: %d, result: %v, expect: %v", idx, arr, inputs[1])
		}

	}

}
