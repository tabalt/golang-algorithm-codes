package chapter_4

import (
	"../utils"
	"testing"
)

var (
	countSortInputCases = [][][]int{
		{
			{9, 6, 2, 5, 4, 8, 7, 1, 3, 6, 7, 8, 4, 1},
			{1, 1, 2, 3, 4, 4, 5, 6, 6, 7, 7, 8, 8, 9},
		},
		{
			{99, 96, 92, 95, 94, 98, 97, 91, 93, 96, 97, 98, 94, 91},
			{91, 91, 92, 93, 94, 94, 95, 96, 96, 97, 97, 98, 98, 99},
		},
	}
)

func Test_Count_Sort_All(t *testing.T) {

	for idx, inputs := range countSortInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		var arr []int

		arr = make([]int, len(inputs[0]))
		copy(arr, inputs[0])
		arr = CountSort(arr)
		// t.Logf("sort result\t%v", arr)
		if !utils.SliceEqual(arr, inputs[1]) {
			t.Errorf("CountSort error, index: %d, result: %v, expect: %v", idx, arr, inputs[1])
		}

	}

}
