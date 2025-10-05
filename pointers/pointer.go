package main

import "fmt"

// ----------------------------------------
// 🔹 PASS BY VALUE FUNCTION
// ----------------------------------------
// Takes a copy of the variable, so changes
// inside this function do NOT affect the original.
func changeNumByValue(num int) {
	num = 5 // modifies only the local copy
	fmt.Println("Inside changeNumByValue function:", num)
}

// ----------------------------------------
// 🔹 PASS BY REFERENCE FUNCTION (with pointer safety)
// ----------------------------------------
// Takes the address of a variable, so it can modify the
// actual memory value. Includes nil safety check.
func changeNumByReference(num *int) {
	// Check if the pointer is not nil before dereferencing
	if num == nil {
		fmt.Println("Pointer is nil — cannot modify the value!")
		return
	}

	*num = 6 // safely modify the original value
	fmt.Println("Inside changeNumByReference function:", *num)
}

func main() {
	num := 1

	// ------------------------------
	// ✅ PASS BY VALUE EXAMPLE
	// ------------------------------
	fmt.Println("----- Pass by Value -----")
	fmt.Println("Before calling changeNumByValue:", num)
	fmt.Println("Memory address of num:", &num)

	changeNumByValue(num)
	fmt.Println("After calling changeNumByValue:", num)
	// num is unchanged because function modified only a copy.

	// ------------------------------
	// ✅ PASS BY REFERENCE EXAMPLE
	// ------------------------------
	fmt.Println("\n----- Pass by Reference -----")
	fmt.Println("Before calling changeNumByReference:", num)
	fmt.Println("Memory address of num:", &num)

	changeNumByReference(&num)
	fmt.Println("After calling changeNumByReference:", num)
	// num is updated since we modified the real variable through pointer.

	// ------------------------------
	// ⚠️ EDGE CASE: Nil pointer
	// ------------------------------
	fmt.Println("\n----- Edge Case: Nil Pointer -----")
	var ptr *int = nil // nil pointer (not pointing to any memory)
	changeNumByReference(ptr) // safely handled (won’t panic)
}
