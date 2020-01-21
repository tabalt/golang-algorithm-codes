package chapter_1

// T(n) = 1+2+3+...+(n-1)+n = (n+1)*n/2 = 0.5*n^2+0.5n = O(n^2)  # 等差数列求和，高斯算法
// S(n) = 1 = O(1)
func FindSameNum1(arr []int) (val, key1, key2 int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] == arr[j] {
				return arr[i], j, i
			}
		}
	}
	return
}

// T(n) = n = O(n)
// S(n) = n = O(n)
func FindSameNum2(arr []int) (val, key1, key2 int) {
	exist := map[int]int{}
	for i := 0; i < len(arr); i++ {
		if _, ok := exist[arr[i]]; ok {
			return arr[i], exist[arr[i]], i
		}
		exist[arr[i]] = i
	}
	return
}
