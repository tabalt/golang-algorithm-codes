package main

import (
	"fmt"
)

func main() {
	var arr []int
	var v int

	arr = []int{9, 6, 2, 5, 4, 8, 7}
	v = 8
	fmt.Println(arr, v)
	fmt.Println(SequenceSearch(arr, v))
	fmt.Println("")
}

func SequenceSearch(arr []int, v int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return i
		}
	}
	return -1
}
