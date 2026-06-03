package main

import "fmt"

// Recursion

func fact(num int) int {
	if num == 1 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	fmt.Println(fact(4))
}
