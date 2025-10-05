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

// passing argurment as a function
func passIt(fn func(a int)) {

}

func main() {
	greet()
	fmt.Println(total(3, 4, 5), total1(1, 2))
	fn := func(){
		fmt.Println("Hi i will be passed into the function")
	}
	passIt(fn())

	
}
