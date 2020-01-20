package chapter_3

import (
	"testing"

	"../utils"
)

var (
	priorityQueueInputCases = [][][]int{
		{
			{7, 1, 3, 10, 5, 2, 8, 9, 6},
			{1, 2, 3, 5, 6, 7, 8, 9, 10},
		},
	}
)

func Test_Priority_Queue_All(t *testing.T) {

	for idx, inputs := range priorityQueueInputCases {
		// t.Logf("    inputs \t%v", inputs[0])

		pq := NewPriotiryQueue()
		for _, v := range inputs[0] {
			pq.In(v)
		}

		result := []int{}
		for i := 0; i < len(inputs[0]); i++ {
			data := pq.Out()
			result = append(result, data)
		}

		// t.Logf("     outed \t%v", result)

		if !utils.SliceEqual(result, inputs[1]) {
			t.Errorf("sort heap error, index: %d, result: %v, expect: %v", idx, result, inputs[1])
		}
	}
}
