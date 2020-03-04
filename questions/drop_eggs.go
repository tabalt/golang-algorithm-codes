package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(superEggDrop(2, 100))
}

var result = map[string]int{}

func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}

	key := strconv.Itoa(K) + "_" + strconv.Itoa(N)

	if ret, ok := result[key]; ok {
		return ret
	}

	var res = int(^uint(0) >> 1)
	for i := 1; i < N+1; i++ {
		res = min(res, max(superEggDrop(K, N-i), superEggDrop(K-1, i-1))+1)
	}

	result[key] = res

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
