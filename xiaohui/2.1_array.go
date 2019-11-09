package xiaohui

func GetArray(arr *[]int, idx int) int {
	a := *arr
	if idx < 0 || idx >= len(a) {
		return 0
	}

	return a[idx]
}

func UpdateArray(arr *[]int, idx, val int) {
	a := *arr
	if idx < 0 || idx >= len(a) {
		return
	}

	a[idx] = val
}

// T(n) = n+n = 2n = O(n)
// S(n) = n = O(n)
func InsertArray(arr *[]int, idx, val int) {
	if idx < 0 {
		return
	}

	// enlarge array
	if idx >= len(*arr) {
		tmp := make([]int, idx+1)
		copy(tmp, *arr)
		*arr = tmp
	}

	//insert item
	a := *arr
	for i := len(a) - 1; i > idx; i-- {
		a[i] = a[i-1]
	}
	a[idx] = val
}

// T(n) = n = O(n)
// S(n) = 1 = O(1)
func DeleteArray(arr *[]int, idx int) {
	a := *arr
	if idx < 0 || idx >= len(a) {
		return
	}

	for i := idx; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = 0
}
