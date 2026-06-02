package main

import "fmt"

func binarySearch(arr []int, item int) int {
	high := len(arr) - 1
	low := 0

	for low <= high {
		mid := (high + low) / 2
		if arr[mid] == item {
			return mid
		} else if arr[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{-1, 3, 5, 7, 10, 15, 29, 950}, 5))
}
