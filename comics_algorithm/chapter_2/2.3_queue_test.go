package chapter_2

import (
	"testing"
)

var (
	queueInputCases = [][]int{
		{1, 2, 3},
		{4, 5, 6, 7, 8},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{5, 3, 8, 2, 0, 3, 4, 6, 9, 21},
	}
)

func Test_Array_Queue_All(t *testing.T) {

	for _, inputs := range queueInputCases {
		aq := NewArrayQueue(5)

		for _, val := range inputs {
			aq.Enqueue(val)
		}

		for i := 0; i < len(inputs); i++ {
			v := inputs[i]
			val := aq.Dequeue()
			if v != val {
				t.Errorf("dequeue array queue error, get: %d, expect: %d", val, v)
			}
		}
	}

}

func Test_List_Queue_All(t *testing.T) {

	for _, inputs := range queueInputCases {
		lq := NewListQueue()

		for _, val := range inputs {
			lq.Enqueue(val)
		}

		for i := 0; i < len(inputs); i++ {
			v := inputs[i]
			val := lq.Dequeue()
			if v != val {
				t.Errorf("dequeue list queue error, get: %d, expect: %d", val, v)
			}
		}
	}

}
