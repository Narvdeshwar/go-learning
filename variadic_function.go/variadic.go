package main

import "fmt"

// ability of function that can handle n number of argument Inside, nums behaves like a slice ([]int).

func test(nums ...int) int {
	var result int = 0
	for _, val := range nums {
		result += val
	}
	return result
}

func main() {
	fmt.Println(test(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// passing the slice in variadic function

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("passing slice in variadic function", test(nums...))
}
