package xiaohui

import (
	"testing"
)

func Test_Doubly_Linked_List_All(t *testing.T) {

	dll := NewDoublyLinkdedList()

	// insert

	var insertCases = []struct {
		idx    int
		val    int
		expect int
	}{
		{0, 1, 1},
		{0, 2, 2},
		{2, 3, 3},
		{1, 4, 4},
		{2, 5, 5},
		{3, 6, 6},
		{-1, 6, -999},
		{999, 6, -999},
	}
	for _, ic := range insertCases {
		dll.Insert(ic.idx, ic.val)
		find := dll.Get(ic.idx)
		if find != nil {
			if find.Data != ic.expect {
				t.Errorf("insert item error, index: %d, value: %d, get: %d, expect: %d", ic.idx, ic.val, find.Data, ic.expect)
			}
		} else {
			if ic.expect != -999 {
				t.Errorf("insert item error, index: %d, value: %d, get: nil, expect: %d", ic.idx, ic.val, ic.expect)
			}
		}
	}

	// delete

	var deleteCases = []struct {
		idx    int
		expect int
	}{
		{0, 2},
		{4, 3},
		{2, 6},
		{-1, -999},
	}
	for _, dc := range deleteCases {
		node := dll.Delete(dc.idx)
		if node != nil {
			if node.Data != dc.expect {
				t.Errorf("delete item error, index: %d, get: %d, expect: %d", dc.idx, node.Data, dc.expect)
			}
		} else {
			if dc.expect != -999 {
				t.Errorf("delete item error, index: %d, get: nil, expect: %d", dc.idx, dc.expect)
			}
		}
	}

}
