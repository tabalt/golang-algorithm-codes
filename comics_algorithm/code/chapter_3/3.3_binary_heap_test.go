package chapter_3

import (
	"testing"

	"../utils"
)

var (
	buildHeapInputCases = [][][]int{
		{
			{7, 1, 3, 10, 5, 2, 8, 9, 6},
			{1, 5, 2, 6, 7, 3, 8, 9, 10},
		},
	}

	sortHeapInputCases = [][][]int{
		{
			{1, 3, 2, 6, 5, 7, 8, 9, 10, 0},
			{10, 9, 8, 7, 6, 5, 3, 2, 1, 0},
		},
	}

	upAdjustInputCases = [][][]int{
		{
			{1, 3, 2, 6, 5, 7, 8, 9, 10, 0},
			{0, 1, 2, 6, 3, 7, 8, 9, 10, 5},
		},
	}
)

func Test_Binary_Heap_Build(t *testing.T) {

	for idx, inputs := range buildHeapInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		bh := &BinaryHeap{}
		bh.Build(inputs[0]) // 构建二叉堆

		// t.Logf("    heaps \t%v", inputs[0])

		if !utils.SliceEqual(inputs[0], inputs[1]) {
			t.Errorf("build heap error, index: %d, result: %v, expect: %v", idx, inputs[0], inputs[1])
		}
	}
}

func Test_Binary_Heap_Sort(t *testing.T) {

	for idx, inputs := range sortHeapInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		bh := &BinaryHeap{}
		bh.Sort(inputs[0])

		// t.Logf("    sorted \t%v", inputs[0])

		if !utils.SliceEqual(inputs[0], inputs[1]) {
			t.Errorf("sort heap error, index: %d, result: %v, expect: %v", idx, inputs[0], inputs[1])
		}
	}
}

func Test_Binary_Heap_UpAdjustChild(t *testing.T) {

	for idx, inputs := range upAdjustInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		bh := &BinaryHeap{}
		bh.UpAdjustChild(inputs[0], len(inputs[0])-1) //插入后上浮最后叶子节点

		// t.Logf("    adjust \t%v", inputs[0])

		if !utils.SliceEqual(inputs[0], inputs[1]) {
			t.Errorf("build heap error, index: %d, result: %v, expect: %v", idx, inputs[0], inputs[1])
		}
	}
}
