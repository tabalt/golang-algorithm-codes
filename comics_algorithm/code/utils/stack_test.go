package utils

import (
	"testing"
)

func TestStack_Len(t *testing.T) {
	var myStack Stack
	myStack.Push(1)
	myStack.Push("test")
	if myStack.Len() != 2 {
		t.Error("Failed Stack.Len")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	var mStack Stack
	if !mStack.IsEmpty() {
		t.Error("Failed Stack.IsEmpty")
	}
}

func TestStack_Push(t *testing.T) {
	var mStack Stack
	mStack.Push(3)
	if mStack.Len() != 1 {
		t.Error("Failed Stack.Push")
	}
}

func TestStack_Top(t *testing.T) {
	var mStack Stack
	if _, err := mStack.Top(); err == nil {
		t.Error("Failed Stack.Top")
	}
	mStack.Push(3)
	if value, _ := mStack.Top(); value != 3 {
		t.Errorf("Failed Stack.Top, value is %d", value)
	}
}

func TestStack_Pop(t *testing.T) {
	var mStack Stack
	if _, err := mStack.Pop(); err == nil {
		t.Error("Failed Stack.Pop")
	}
	mStack.Push("test")
	mStack.Push(3)
	if value, _ := mStack.Pop(); value != 3 || mStack.Len() != 1 {
		t.Errorf("Failed Stack.Pop, value is %d, len is %d", value, mStack.Len())
	}
}
