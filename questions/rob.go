package main

import (
	"fmt"
)

func main() {
	var shops = []int{2, 7, 9, 3, 1}
	fmt.Println(shops)
	fmt.Println(doRob1(shops, len(shops)-1), "\n--------")
	fmt.Println(doRob2(shops, len(shops)-1), "\n--------")
	fmt.Println(doRob3(shops), "\n--------")
}

// func rob(nums []int) int {
// 	return doRob2(nums, len(nums)-1)
// }

//暴力递归
func doRob1(nums []int, idx int) int {
	if idx < 0 {
		return 0
	}

	fmt.Println(idx)

	now := nums[idx] + doRob1(nums, idx-2)
	next := doRob1(nums, idx-1)
	mx := max(now, next)

	fmt.Println("\t", idx, now, next, mx)

	return mx
}

var robCache map[int]int

func init() {
	robCache = map[int]int{}
}

// 暴力递归，缓存结果优化
func doRob2(nums []int, idx int) int {
	if idx < 0 {
		return 0
	}

	if mx, ok := robCache[idx]; ok {
		return mx
	}

	mx := max(nums[idx]+doRob2(nums, idx-2), doRob2(nums, idx-1))
	robCache[idx] = mx

	return mx
}

// 递推
func doRob3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var result = make([]int, len(nums))
	result[0] = nums[0]
	result[1] = max(nums[0], nums[1])

	for idx := 2; idx < len(nums); idx++ {
		result[idx] = max(nums[idx]+result[idx-2], result[idx-1])
	}

	return result[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
