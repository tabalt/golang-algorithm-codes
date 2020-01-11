package utils

import (
	"errors"
)

type Queue []interface{}

func (queue Queue) Len() int {
	return len(queue)
}

func (queue Queue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue Queue) Cap() int {
	return cap(queue)
}

func (queue *Queue) In(value interface{}) {
	*queue = append(*queue, value)
}

func (queue *Queue) Out() (interface{}, error) {
	theQueue := *queue
	if len(theQueue) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theQueue[0]
	*queue = theQueue[1:]
	return value, nil
}
