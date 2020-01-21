package chapter_2

import (
	"testing"
)

var (
	stackInputCases = [][]int{
		{1, 2, 3},
		{4, 5, 6, 7, 8},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{5, 3, 8, 2, 0, 3, 4, 6, 9, 21},
	}
)

func Test_Array_Stack_All(t *testing.T) {

	for _, inputs := range stackInputCases {
		as := NewArrayStack(5)

		for _, val := range inputs {
			as.Push(val)
		}

		for i := len(inputs) - 1; i >= 0; i-- {
			v := inputs[i]
			val := as.Pop()
			if v != val {
				t.Errorf("pop array stack error, get: %d, expect: %d", val, v)
			}
		}
	}

}

func Test_List_Stack_All(t *testing.T) {

	for _, inputs := range stackInputCases {
		ls := NewListStack()

		for _, val := range inputs {
			ls.Push(val)
		}

		for i := len(inputs) - 1; i >= 0; i-- {
			v := inputs[i]
			val := ls.Pop()
			if v != val {
				t.Errorf("pop list stack error, get: %d, expect: %d", val, v)
			}
		}
	}

}
