package chapter_2

import (
	"testing"
)

func Test_Array_Struct_All(t *testing.T) {

	data := []int{3, 1, 2, 5, 4, 9, 7, 2}
	arr := NewArray(len(data))

	// init
	for idx, val := range data {
		arr.Insert(idx, val)
	}

	// read
	idx := 3
	val := 5
	v := arr.Get(idx)
	if v != val {
		t.Errorf("read item error, index: %d, value: %d, expect: %d", idx, v, val)
	}

	// update
	idx, val, expect := 3, 8, []int{3, 1, 2, 8, 4, 9, 7, 2}
	arr.Update(idx, val)
	if !SliceEqual(arr.Data(), expect) {
		t.Errorf("update item error, index: %d, value: %d, array: %d, expect: %d", idx, val, arr.Data(), expect)
	}

	// insert

	var insertCases = []struct {
		idx    int
		val    int
		expect []int
	}{
		{10, 10, []int{3, 1, 2, 8, 4, 9, 7, 2}},
		{8, 11, []int{3, 1, 2, 8, 4, 9, 7, 2, 11}},
		{6, 12, []int{3, 1, 2, 8, 4, 9, 12, 7, 2, 11}},
		{0, 13, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2, 11}},
		{-1, 14, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2, 11}},
	}
	for _, ic := range insertCases {
		arr.Insert(ic.idx, ic.val)
		if !SliceEqual(arr.Data(), ic.expect) {
			t.Errorf("insert item error, index: %d, value: %d, array: %d, expect: %d", ic.idx, ic.val, arr.Data(), ic.expect)
		}
	}

	// delete

	var deleteCases = []struct {
		idx    int
		expect []int
	}{
		{11, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2, 11}},
		{10, []int{13, 3, 1, 2, 8, 4, 9, 12, 7, 2}},
		{8, []int{13, 3, 1, 2, 8, 4, 9, 12, 2}},
		{0, []int{3, 1, 2, 8, 4, 9, 12, 2}},
		{-1, []int{3, 1, 2, 8, 4, 9, 12, 2}},
	}
	for _, dc := range deleteCases {
		arr.Delete(dc.idx)
		if !SliceEqual(arr.Data(), dc.expect) {
			t.Errorf("delete item error, index: %d, array: %d, expect: %d", dc.idx, arr.Data(), dc.expect)
		}
	}

	// delete unsort

	var deleteUnsortCases = []struct {
		idx    int
		expect []int
	}{
		{-1, []int{3, 1, 2, 8, 4, 9, 12, 2}},
		{0, []int{2, 1, 2, 8, 4, 9, 12}},
		{6, []int{2, 1, 2, 8, 4, 9}},
		{3, []int{2, 1, 2, 9, 4}},
	}
	for _, duc := range deleteUnsortCases {
		arr.DeleteUnsort(duc.idx)
		if !SliceEqual(arr.Data(), duc.expect) {
			t.Errorf("delete unsort item error, index: %d, array: %d, expect: %d", duc.idx, arr.Data(), duc.expect)
		}
	}

}
