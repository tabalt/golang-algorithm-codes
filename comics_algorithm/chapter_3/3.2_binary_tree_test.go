package chapter_3

import (
	"testing"

	"../utils"
)

var (
	null = EmptyNodeData

	binaryTreeInputCases = [][][]int{
		{
			{1, 2, 4, null, null, 5, null, null, 3, null, 6}, //input
			{1, 2, 4, 5, 3, 6},                               // pre order
			{4, 2, 5, 1, 3, 6},                               // in order
			{4, 5, 2, 6, 3, 1},                               // post order
			{1, 2, 3, 4, 5, 6},                               // level order
		},
		{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{9, 8, 7, 6, 5, 4, 3, 2, 1},
			{9, 8, 7, 6, 5, 4, 3, 2, 1},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			{3, 2, 9, null, null, 10, null, null, 8, 7, null, null, 4},
			{3, 2, 9, 10, 8, 7, 4},
			{9, 2, 10, 3, 7, 8, 4},
			{9, 10, 2, 7, 4, 8, 3},
			{3, 2, 8, 9, 10, 7, 4},
		},
	}
)

func Test_Binary_Tree_All(t *testing.T) {

	for idx, inputs := range binaryTreeInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		bt := NewBinaryTree(&inputs[0])
		bti := &BinaryTreeIterator{}

		var result *[]int

		result = &[]int{}
		bti.PreOrder(bt, result)
		// t.Logf(" pre order\t%v", result)
		if !utils.SliceEqual(*result, inputs[1]) {
			t.Errorf("traversal pre order error, index: %d, result: %v, expect: %v", idx, *result, inputs[1])
		}

		result = &[]int{}
		bti.InOrder(bt, result)
		// t.Logf("  in order\t%v", result)
		if !utils.SliceEqual(*result, inputs[2]) {
			t.Errorf("traversal in order error, index: %d, result: %v, expect: %v", idx, *result, inputs[2])
		}

		result = &[]int{}
		bti.PostOrder(bt, result)
		// t.Logf("post order\t%v", result)
		if !utils.SliceEqual(*result, inputs[3]) {
			t.Errorf("traversal post order error, index: %d, result: %v, expect: %v", idx, *result, inputs[3])
		}

		result = &[]int{}
		bti.LevelOrder(bt, result)
		// t.Logf("level order\t%v", result)
		if !utils.SliceEqual(*result, inputs[4]) {
			t.Errorf("traversal level order error, index: %d, result: %v, expect: %v", idx, *result, inputs[4])
		}

		result = &[]int{}
		bti.PreOrderWithStack(bt, result)
		// t.Logf(" pre order with stack\t%v", result)
		if !utils.SliceEqual(*result, inputs[1]) {
			t.Errorf("traversal pre order with stack error, index: %d, result: %v, expect: %v", idx, *result, inputs[1])
		}
	}

}
