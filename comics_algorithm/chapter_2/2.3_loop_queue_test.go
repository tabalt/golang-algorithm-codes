package chapter_2

import (
	"math/rand"
	"testing"
	"time"
)

var (
	loopInputCases = [][]int{
		{1, 2, 3},
		{4, 5, 6, 7, 8},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{5, 3, 8, 2, 0, 3, 4, 6, 9, 21},
	}
)

func Test_Loop_Queue_Enqueue_Dequeue(t *testing.T) {

	for _, inputs := range loopInputCases {
		lq := NewLoopQueue(len(inputs) + 1)

		for _, val := range inputs {
			err := lq.Enqueue(val)
			if err != nil {
				t.Errorf("enqueue loop queue error: %v", err)
			}
		}

		for i := 0; i < len(inputs); i++ {
			v := inputs[i]
			val, err := lq.Dequeue()
			if err != nil {
				t.Errorf("dequeue loop queue error: %v", err)
			} else if v != val {
				t.Errorf("dequeue loop queue error, get: %d, expect: %d", val, v)
			}
		}
	}

}

func Test_Loop_Queue_Loop(t *testing.T) {

	for _, inputs := range loopInputCases {
		lq := NewLoopQueue(len(inputs) + 1)
		rand.Seed(time.Now().UnixNano())
		dequeueCount := len(inputs) - 1 - rand.Intn(2)

		// t.Logf("loop queue status 1, front:%v, rear:%v, data: %v", lq.front, lq.rear, lq.storage)

		for _, val := range inputs {
			err := lq.Enqueue(val)
			if err != nil {
				t.Errorf("enqueue loop queue error: %v", err)
			}
		}
		if lq.front != 0 || lq.rear != len(inputs) {
			t.Errorf("enqueue loop queue error, front:%v, rear:%v, data: %v", lq.front, lq.rear, lq.storage)
		}

		// t.Logf("loop queue status 2, front:%v, rear:%v, data: %v", lq.front, lq.rear, lq.storage)

		for i := 0; i < dequeueCount; i++ {
			_, err := lq.Dequeue()
			if err != nil {
				t.Errorf("dequeue loop queue error: %v", err)
			}
		}
		if lq.front != dequeueCount || lq.rear != len(inputs) {
			t.Errorf("dequeue loop queue error, front:%v, rear:%v, data: %v", lq.front, lq.rear, lq.storage)
		}

		// t.Logf("loop queue status 3, front:%v, rear:%v, data: %v", lq.front, lq.rear, lq.storage)

		for idx, val := range inputs {
			if idx >= dequeueCount {
				break
			}

			err := lq.Enqueue(val)
			if err != nil {
				t.Errorf("enqueue loop queue error: %v", err)
			}
		}
		if lq.front != dequeueCount || lq.rear != (dequeueCount-1) {
			t.Errorf("enqueue loop queue error, front:%v, rear:%v, data: %v", lq.front, lq.rear, lq.storage)
		}

		// t.Logf("loop queue status 4, dequeueCount:%v, front:%v, rear:%v, data: %v", dequeueCount, lq.front, lq.rear, lq.storage)
	}

}
