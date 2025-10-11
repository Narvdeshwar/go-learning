package main

import "fmt"

// without argument function
func greet() {
	fmt.Println("Good morning")
}

// passing argument in a function
func total1(a int, b int) int {
	return a + b
}

// if the argument having the same dataType then we can make it more simple for above
func total(a, b, c int) int {
	return a + b + c
}

// Multiple return values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division with zero")
	}
	return a / b, nil
}

// Named return values
func reactangle(a, b int) (area int, perimeter int) {
	area = a * b
	perimeter = 2 * (a + b)
	return // implicit return
}

// first class and higher order funcion
func apply(fb func(int, int) int, a, b int) int {
	return fb(a, b)
}

func main() {
	greet()
	// fmt.Println(total(3, 4, 5), total1(1, 2))
	// fmt.Println(divide(1, 1))
	area, perimeter := reactangle(2, 3)
	fmt.Println("area", area, "perimeter", perimeter)

	add := func(a, b int) int {
		return a + b
	}
	multi := func(a, b int) int {
		return a * b
	}
	fmt.Println("Adding", apply(add, 3, 4))
	fmt.Println("Adding", apply(multi, 5, 6))

}
