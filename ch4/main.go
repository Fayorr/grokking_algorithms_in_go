package main

import "fmt"

func sumRecursion(arr []int) int {


	//base case
	if len(arr) == 0 {
		return  0
	}
	//recursive case
	
	return arr[0] + sumRecursion(arr[1:])
}

// Write a recursive function to count the number of items in a list.
func countRecursion(arr []int) int {

	//base case
	if len(arr) == 0 {
		return  0
	}
	//recursive case
	
	return 1 + countRecursion(arr[1:])
}
// Write a recursive function to find the maximum number in a list.
func maxRecursion(arr []int) int {

	//base case
	if len(arr) == 1 {
		return arr[0]
	}
	maxOfRest := maxRecursion(arr[1:])
	//recursive case
	if arr[0] > maxOfRest {
		return arr[0]
	}
	return maxOfRest
}
func recursiveBinarySearch(arr []int, num int) int {
	return binarySearchHelper(arr, num, 0, len(arr)-1)
	}

func binarySearchHelper(arr []int, num, low, high int) int {
	if low > high {
        return -1 
    }

	mid := low + (high - low) / 2
	if arr[mid] == num {
		return mid
		} else if arr[mid] < num {
			return binarySearchHelper(arr, num, mid+1, high)
		} else {
			return binarySearchHelper(arr, num, low, mid-1)
		}
}


func main() {
	fmt.Println(sumRecursion([]int{2,4,6,8}))
	fmt.Println(countRecursion([]int{2,4,6,8}))
	fmt.Println(maxRecursion([]int{2,400,6,8,200, 50,64,66,3422,345,687,899,32,3,45,56565,43,323,455,5656,565345}))
	fmt.Println(recursiveBinarySearch([]int{1,3,5,8,9,12,20,33,48,55,67,86,92,105}, 12))
}