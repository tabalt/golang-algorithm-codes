package chapter_1

import (
	"testing"
)

func Test_FindSameNum_All(t *testing.T) {

	arr := []int{3, 1, 2, 5, 4, 9, 7, 2}

	val, key1, key2 := FindSameNum1(arr)
	if val != 2 || key1 != 2 || key2 != 7 {
		t.Errorf("FindSameNum1 error, val: %d, key1: %d key2: %d", val, key1, key2)
	}

	val, key1, key2 = FindSameNum2(arr)
	if val != 2 || key1 != 2 || key2 != 7 {
		t.Errorf("FindSameNum2 error, val: %d, key1: %d key2: %d", val, key1, key2)
	}

}
