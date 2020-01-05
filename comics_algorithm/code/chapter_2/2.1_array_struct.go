package chapter_2

type Array struct {
	data []int
	size int
}

func NewArray(n int) *Array {
	return &Array{make([]int, n), 0}
}

func (a *Array) Data() []int {
	return a.data[:a.size]
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) Get(idx int) int {
	if idx < 0 || idx >= a.size {
		return 0
	}

	return a.data[idx]
}

func (a *Array) Update(idx, val int) bool {
	if idx < 0 || idx >= a.size {
		return false
	}

	a.data[idx] = val
	return true
}

// T(n) = n+n = 2n = O(n)
// S(n) = n = O(n)
func (a *Array) Insert(idx, val int) {
	if idx < 0 || idx > a.size {
		return
	}

	// enlarge array
	if a.size >= len(a.data) {
		tmp := make([]int, len(a.data)*2)
		copy(tmp, a.data)
		a.data = tmp
	}

	//insert item
	for i := a.size - 1; i >= idx; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[idx] = val
	a.size++
}

// T(n) = n = O(n)
// S(n) = 1 = O(1)
func (a *Array) Delete(idx int) int {
	if idx < 0 || idx >= a.size {
		return ErrorIntValue
	}

	deleted := a.data[idx]
	for i := idx; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--

	return deleted
}

// T(n) = 1+1 = O(1)
// S(n) = 1 = O(1)
func (a *Array) DeleteUnsort(idx int) {
	if idx < 0 || idx >= a.size {
		return
	}

	a.data[idx] = a.data[a.size-1]
	a.size--
}
