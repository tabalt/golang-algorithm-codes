package chapter_2

import (
	"testing"
)

func SliceEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func Test_Array_All(t *testing.T) {

	arr := []int{3, 1, 2, 5, 4, 9, 7, 2}

	// read
	idx := 3
	val := 5
	if GetArray(&arr, idx) != val {
		t.Errorf("read item error, index: %d, value: %d, expect: %d", idx, arr[idx], val)
	}

	// update
	idx, val, expect := 3, 8, []int{3, 1, 2, 8, 4, 9, 7, 2}
	UpdateArray(&arr, idx, val)
	if !SliceEqual(arr, expect) {
		t.Errorf("update item error, index: %d, value: %d, array: %d, expect: %d", idx, val, arr, expect)
	}

	// insert

	var insertCases = []struct {
		idx    int
		val    int
		expect []int
	}{
		{8, 10, []int{3, 1, 2, 8, 4, 9, 7, 2, 10}},
		{10, 11, []int{3, 1, 2, 8, 4, 9, 7, 2, 10, 0, 11}},
		{6, 12, []int{3, 1, 2, 8, 4, 9, 12, 7, 2, 10, 0}},
		{0, 13, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2, 10}},
		{-1, 14, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2, 10}},
	}
	for _, ic := range insertCases {
		InsertArray(&arr, ic.idx, ic.val)
		if !SliceEqual(arr, ic.expect) {
			t.Errorf("insert item error, index: %d, value: %d, array: %d, expect: %d", ic.idx, ic.val, arr, ic.expect)
		}
	}

	// delete

	var deleteCases = []struct {
		idx    int
		expect []int
	}{
		{11, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2, 10}},
		{10, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2, 0}},
		{8, []int{13, 3, 1, 2, 8, 4, 9, 12, 2, 0, 0}},
		{0, []int{3, 1, 2, 8, 4, 9, 12, 2, 0, 0, 0}},
		{-1, []int{3, 1, 2, 8, 4, 9, 12, 2, 0, 0, 0}},
	}
	for _, dc := range deleteCases {
		DeleteArray(&arr, dc.idx)
		if !SliceEqual(arr, dc.expect) {
			t.Errorf("delete item error, index: %d, array: %d, expect: %d", dc.idx, arr, dc.expect)
		}
	}

}
