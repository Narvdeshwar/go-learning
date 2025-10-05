// A closure is a function that "remembers" variables from the environment where it was created.
package main

import "fmt"

// closureFunction returns an anonymous function (a closure) that maintains
// its own state for the variable `count` between calls.
func closureFunction() func() int {
	var count int = 0 // local variable, will be captured by the closure

	// The returned function forms a closure around `count`,
	// meaning it can access and modify `count` even after
	// closureFunction has finished executing.
	return func() int {
		count += 1 // increment the captured variable
		return count
	}
}

func main() {
	// Create a closure instance. Each call to closureFunction
	// creates a new independent `count` variable.
	increment := closureFunction()

	fmt.Println("1st call:", increment())  // Output: 1
	fmt.Println("2nd call:", increment())  // Output: 2

	// Here we print the function itself, not its result.
	// It will display the memory address or function signature.
	fmt.Println("Printing the function directly:", closureFunction())

	// 🧠 Edge Case Demonstration:
	// Each new call to closureFunction() gets its own copy of `count`.
	newIncrement := closureFunction()
	fmt.Println("New closure instance first call:", newIncrement()) // Output: 1 (independent state)
	fmt.Println("Old closure instance next call:", increment())     // Output: 3 (continues old state)
}
