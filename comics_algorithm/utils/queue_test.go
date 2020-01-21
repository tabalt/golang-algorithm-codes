package utils

import (
	"testing"
)

func TestQueue_Len(t *testing.T) {
	var myQueue Queue
	myQueue.In(1)
	myQueue.In("test")
	if myQueue.Len() != 2 {
		t.Error("Failed Queue.Len")
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	var myQueue Queue
	if !myQueue.IsEmpty() {
		t.Error("Failed Queue.IsEmpty")
	}
}

func TestQueue_In(t *testing.T) {
	var myQueue Queue
	myQueue.In(3)
	if myQueue.Len() != 1 {
		t.Error("Failed Queue.In")
	}
}

func TestQueue_Out(t *testing.T) {
	var myQueue Queue
	if _, err := myQueue.Out(); err == nil {
		t.Error("Failed Queue.Out")
	}
	myQueue.In("test")
	myQueue.In(3)
	if value, _ := myQueue.Out(); value != "test" || myQueue.Len() != 1 {
		t.Errorf("Failed Queue.Out, value is %d, len is %d", value, myQueue.Len())
	}
}
