package main

import "fmt"

func main() {
	// creation of map
	m := make(map[string]int)

	// adding values to map
	m["alice"] = 23
	m["bob"] = 34

	for index, key := range m {
		fmt.Println(index, "->", key)
	}
	// fmt.Println(m["ashrith"]) // it returns 0 when the key or value doesn;t exits

	// checking any value exits in map or not
	val, exits := m["ashrith"]
	if exits {
		fmt.Println("current value", val)
	} else {
		fmt.Println("current value doesn't exits", val)
	}

	delete(m, "bob")
	fmt.Println(m)

}
