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
