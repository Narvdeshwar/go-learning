package main

import "fmt"

func main() {
	m := []int{1, 2, 3, 4, 5}
	fmt.Println(m)

	// If we want to skip the index, we can use '_' in place of index
	for _, val := range m {
		fmt.Println("value =", val)
	}

	// Printing both index and value
	for i, v := range m {
		fmt.Println("index:", i, "value:", v)
	}

	// Printing only the index (key)
	for k := range m {
		fmt.Println("index only =", k)
	}

	// Iterating over a string
	for index, val := range "golang" {
		// 'val' here is a rune (Unicode code point), so it prints the integer value
		fmt.Println("index =", index, "rune =", val)

		// To print the actual character, convert rune to string
		fmt.Println("index =", index, "character =", string(val))
	}

	// Iterating over a map
	personAge := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 22,
	}

	// Printing key and value
	for name, age := range personAge {
		fmt.Println("name =", name, "age =", age)
	}

	// Printing only keys
	for name := range personAge {
		fmt.Println("name only =", name)
	}

	// Printing only values
	for _, age := range personAge {
		fmt.Println("age only =", age)
	}
}
