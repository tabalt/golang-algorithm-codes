package chapter_4

//计数排序
//适用于元素取值在有限范围内的整数，如0-10、90-99等
//将元素出现的次数存在 和元素值相等的下标所在数组位置处，遍历数组输出即可
func CountSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	//求最大值和最小值
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	//计算偏移，计数
	bias := 0 - min
	counts := make([]int, max-min+1)
	for _, v := range arr {
		counts[v+bias]++
	}

	//遍历输出
	result := []int{}
	for v, count := range counts {
		for i := 0; i < count; i++ {
			result = append(result, v-bias)
		}
	}

	return result
}
