package main

import (
    "fmt"
)

func main() {
    var arr []int
    var v int

    arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    v = 7
    fmt.Println(arr, v)
    fmt.Println(BinarySearch(arr, v))
    fmt.Println("")

    arr = []int{1, 3, 4, 6, 8, 9, 14, 16}
    v = 9
    fmt.Println(arr, v)
    fmt.Println(BinarySearch2(arr, v, 0, len(arr)))
    fmt.Println("")
}

func BinarySearch(arr []int, v int) int {
    low, high := 0, len(arr)-1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == v {
            return mid
        } else if arr[mid] < v {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func BinarySearch2(arr []int, v, low, high int) int {
    mid := (low + high) / 2
    if arr[mid] == v {
        return mid
    } else if arr[mid] < v {
        return BinarySearch2(arr, v, mid+1, high)
    } else {
        return BinarySearch2(arr, v, low, mid-1)
    }

    return -1
}
