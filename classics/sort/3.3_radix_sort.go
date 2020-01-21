package main

import (
    "fmt"
)

func main() {
    var arr []int

    arr = []int{9, 6, 2, 5, 4, 8, 7, 1, 3, 6, 7, 8, 4, 1}
    fmt.Println(arr)
    result := RadixSort(arr)
    fmt.Println(result)
    fmt.Println("")

    arr = []int{99, 96, 92, 95, 94, 98, 97, 91, 93, 96, 97, 98, 94, 91}
    fmt.Println(arr)
    result = RadixSort(arr)
    fmt.Println(result)
    fmt.Println("")

}

// 基数排序（Radix Sort） O(n*k)
// 按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位
func RadixSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    //计算最大值
    max := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }

    //计算最大值有几位数
    maxDigit := 0
    for max != 0 {
        max /= 10
        maxDigit++
    }

    //创建10个桶
    buckets := make([][]int, 10)

    //从低位开始进行每趟排序
    mod, div := 10, 1
    for i := 0; i < maxDigit; i++ {
        //按每个数第i位的值放入对应桶中
        for j := 0; j < len(arr); j++ {
            radix := arr[j] % mod / div
            buckets[radix] = append(buckets[radix], arr[j])
        }

        //合并各桶数据回原数组，并清空桶
        index := 0
        for j := 0; j < len(buckets); j++ {
            for k := 0; k < len(buckets[j]); k++ {
                arr[index] = buckets[j][k]
                index++
            }
            buckets[j] = []int{}
        }

        //调整mod和div，按上一位继续分桶排序
        mod *= 10
        div *= 10
    }

    return arr

}
