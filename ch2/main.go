package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0

	for i:=0; i<len(arr); i++ {
		if smallest > arr[i] {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(arr []int) []int {
	copiedArr := make([]int, len(arr))
	copy(copiedArr, arr)
	newArr := []int{}

	n := len(copiedArr)
	for i:=0; i<n; i++ {
		smallest := findSmallest(copiedArr)
		newArr = append(newArr, copiedArr[smallest])
		copiedArr = append(copiedArr[:smallest], copiedArr[smallest+1:]...)
	}
return newArr
}

func main() {
	fmt.Println(selectionSort([]int{23,25,10,40,33}))
}