# Interface

An interface in Go is a type that defines a set of methods.
If a type implements those methods, it automatically satisfies the interface — no “implements” keyword needed!

<i>Go me interface ek aisa type hota hai jisme methods ka contract likha hota hai.
Agar koi struct un methods ko implement karta hai, to wo automatically us interface ko satisfy karta hai.
Kisi “implements” ya “extends” keyword ki zarurat nahi hoti jaise Java/C++ me hoti hai.</i>

### Example 1: Basic

```go
package main

import "fmt"

// Step 1: Define an interface
// Interface is like a contract — it defines a behavior (method) that any type can implement.
// Here, any type that has a Speak() method returning string is considered a Speaker.
type Speaker interface {
	Speak() string
}

// Step 2: Define a struct (data type)
// 'Person' struct represents a person with a 'name' field.
type Person struct {
	name string
}

// Step 3: Implement the interface method
// Since 'Person' has a 'Speak() string' method, it automatically satisfies the 'Speaker' interface.
// (Go me interface explicit nahi hota — agar method match kar gaya to implement ho gaya.)
func (p Person) Speak() string {
	return "Hello I'm " + p.name
}

func main() {
	// Create a new 'Person' object with name "Ashrith"
	newPerson := Person{"Ashrith"}

	// Call the Speak() method defined on 'Person'
	// It will print "Hello I'm Ashrith"
	fmt.Println(newPerson.Speak())
}

```

## Multiple Structs Implementing Same Interface

Multiple structs can implement the same interface, as long as they all have the required methods.
This enables polymorphism — the same interface variable can hold different types.

<i>
Kai alag-alag structs ek hi interface ko implement kar sakte hain.
Isse polymorphism possible hota hai — ek interface variable alag type ke objects rakh sakta hai.
</i>

### Example 2: Polymorphism

```go
package main

import "fmt"

// Step 1: Define an interface
// 'Shape' defines a behavior — any type that has an Area() method returning float32
// is considered a 'Shape'.
type Shape interface {
	Area() float32
}

// Step 2: Define a Circle struct
type Circle struct {
	radius float32
}

// Step 3: Implement the Area() method for Circle
// Since Circle has an Area() method, it automatically satisfies Shape interface.
func (c Circle) Area() float32 {
	// Formula for circle area = π * r * r
	return 3.14 * c.radius * c.radius
}

// Step 4: Define a Rectangle struct
type Rectangle struct {
	length, breadth float32
}

// Step 5: Implement the Area() method for Rectangle
func (r Rectangle) Area() float32 {
	return r.length * r.breadth
}

// Step 6: Polymorphic function
// This function can accept ANY type that implements Shape interface.
// It doesn't care if it's a Circle or Rectangle — both have Area() method.
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// Step 7: main() function
func main() {
	// Create a new Circle
	newCircle := Circle{radius: 10}
	// Pass it directly to printArea() — works because Circle implements Shape
	printArea(newCircle)

	// Create a new Rectangle
	newRectangle := Rectangle{length: 5, breadth: 4}
	// Pass it directly to printArea() — works because Rectangle implements Shape
	printArea(newRectangle)
}

```

## 1. Empty Interface (interface{})

- interface{} means no method — so every type satisfies it. <br/>
- Used for any type of value (like any in TypeScript).

### Example for empty interface

```go
package main

import "fmt"

// describe() function accepts an empty interface as parameter
// 'interface{}' (empty interface) means it can take ANY type of value.
// In Go, every type implements at least zero methods, so every type satisfies interface{}.
func describe(i interface{}) {
	// %v → prints the value
	// %T → prints the actual type of the value stored inside 'i'
	fmt.Printf("Value of i = %v, and type = %T\n", i, i)
}

func main() {
	// Passing different primitive data types
	describe(43)              // int
	describe("Jai shree ram") // string
	describe(true)             // bool

	// ---------- Composite Data Types ----------

	// 1️⃣ Slice example
	// Slice is a dynamic array that can grow or shrink.
	// Here we create a slice of integers.
	arr := []int{1, 2, 3, 4, 5}
	describe(arr)

	// 2️⃣ Map example
	// Map is like a dictionary — it stores key-value pairs.
	// Keys and values can be of any type (but keys must be comparable).
	person := map[string]string{
		"name":         "Ram",
		"placeOfBirth": "Ayodhaya",
	}
	describe(person)

	// 3️⃣ Struct example
	// Struct is a user-defined type that groups related fields together.
	// Think of it as a lightweight object (without classes or inheritance).
	type User struct {
		name         string
		placeOfBirth string
	}

	// Create a new struct instance
	newUser := User{"Ram", "Ayodhaya"}
	describe(newUser)
}

```

## Type Asseration

```go
package main

import "fmt"

func main() {
	// Step 1️⃣: Create an empty interface variable
	// 'interface{}' can hold any type of value
	var i interface{} = 123 // Here, we store an int value

	// Step 2️⃣: Type assertion (safe version)
	// We are trying to extract 'string' type from interface 'i'
	// Syntax: value, ok := i.(Type)
	// - If 'i' actually holds that type → ok = true
	// - If not → ok = false (no panic occurs)
	str, ok := i.(string)

	// Step 3️⃣: Check if type assertion succeeded
	if !ok {
		// If 'i' does NOT contain a string → ok = false
		fmt.Println("Value is not string type")
	} else {
		// If 'i' contains a string → ok = true
		fmt.Println("Value is", str)
	}
}

```

### Embedding Interface
Go allows you to combine interfaces by embedding one interface inside another.
This helps build complex behaviours from simple one.

```go
package main

import "fmt"

// Reader interface declares a Read() method.
// Any type that implements this method satisfies the Reader interface.
type Reader interface {
	Read()
}

// Writer interface declares a Write() method.
// Any type that implements this method satisfies the Writer interface.
type Writer interface {
	Write()
}

// ReadWrite interface embeds both Reader and Writer interfaces.
// → Any type that implements both Read() and Write()
//   automatically implements ReadWrite.
type ReadWrite interface {
	Reader
	Writer
}

// File struct represents a concrete type that will implement both methods.
type File struct{}

// Read implements the Reader interface for File.
func (File) Read() {
	fmt.Println("Reading the File")
}

// Write implements the Writer interface for File.
func (File) Write() {
	fmt.Println("Writing the File")
}

func main() {
	// Declare a variable of interface type ReadWrite.
	// Assign a File value to it — this works because File implements both Read() and Write().
	var rw ReadWrite = File{}

	// Interface method calls — resolved at runtime through dynamic dispatch.
	rw.Read()
	rw.Write()
}

```
