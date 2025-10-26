# what is _Struct_

A struct in Go is a user-defined data type that allows you to group together different types of data under one name.<br/>
It’s like a container that holds related data (like an object in OOP, but Go doesn’t have classes).

## syntax

```go
type struct_name struct{
    varName1 dataType
    varName2 dataType
}
```

### example

```go
package main

import "fmt"

// define the struct
type order struct {
	orderId string
	name    string
	price   float32
}

func main() {
	// creating the instance of struct
	o1 := order{orderId: "23os", name: "Ashrith", price: 24}
	fmt.Println(o1)
	// access fields using there name
	fmt.Println(o1.name)
}

```

# Different ways to use the struct

We can create structs in multiple ways, pass them to functions, and even use them with pointers.

## Example1: Different Ways to Initialize a Struct

```go
package main

import "fmt"

// Definition of car struct
type car struct {
	modelNo  string
	brand    string
	seatType int
}

func main() {
	// 1. Using field names
	c1 := car{modelNo: "1x2se", brand: "hundai", seatType: 4}

	// 2. Without field names (order matters)
	c2 := car{"134", "honda", 5}

	// 3. Zero value struct (fields take their default values)
	var c3 car
	fmt.Println(c3) // Output: {  0}

	// Assign values later
	c3.brand = "kia"
	c3.modelNo = "23x21"
	c3.seatType = 5

	fmt.Println(c1, c2, c3)
}

```

## Example2: Passing Structs to Functions

```go
package main

import "fmt"

// Define a struct type 'user' with three fields
type user struct {
	name   string
	branch string
	rollNo int
}

// Function to display user details (passed by value)
func show(u user) {
	fmt.Println("Name:", u.name, "Branch:", u.branch, "Roll No:", u.rollNo)
}

// Function to modify user's roll number (passed by pointer)
func modifyUser(u *user, newRollNo int) {
	u.rollNo = newRollNo // dereference pointer to modify original data
}

func main() {
	// Create a user struct instance using a composite literal
	user1 := user{"Ashrith", "Computer", 23}

	// Display initial details
	show(user1) // Output# Name: Ashrith Branch: Computer Roll No: 23

	// Modify roll number directly
	user1.rollNo = 45
	show(user1) // Output# Name: Ashrith Branch: Computer Roll No: 45

	// Modify roll number using a function with pointer receiver
	modifyUser(&user1, 56)
	show(user1) // Output# Name: Ashrith Branch: Computer Roll No: 56
}

```

## Example3: Nested (Embedded) Struct

```go
package main

import "fmt"

// Define a struct type 'engine' with one field
type engine struct {
	horsePower int
}

// Define a struct type 'car' that embeds 'engine' inside it
// This means 'car' will have all fields of 'engine' as if they were its own
type car struct {
	brand  string
	engine // embedded struct (composition)
}


func main() {
	// Create an instance of 'car' and initialize both fields
	car1 := car{
		brand:  "Hundai",
		engine: engine{horsePower: 300}, // initialize the embedded 'engine' struct
	}

	// Print the full 'car' struct (includes embedded struct fields)
	fmt.Println(car1)

	// Access and print only the embedded 'engine' struct
	fmt.Println(car1.engine)

	// You can also access the embedded field directly without using 'engine'
	fmt.Println("Horsepower:", car1.horsePower)
}
```

## Example 4: Anonymous Struct

```go
package main

import "fmt"

func main() {
	// Define and initialize an anonymous struct
	// (a struct without a named type)
	person := struct {
		name  string
		email string
	}{
		name:  "Ashrith",
		email: "ashrith02.oct@gmail.com",
	}

	// Print the entire struct
	fmt.Println(person)
}

```

### Anonymous vs Named Struct

| Feature               | Anonymous Struct                  | Named Struct                    |
| --------------------- | --------------------------------- | ------------------------------- |
| **Definition**        | Created inline                    | Declared with `type` keyword    |
| **Reusability**       | ❌ Only in that function or scope | ✅ Can be reused anywhere       |
| **Best for**          | Temporary or one-off structs      | Permanent, reusable data models |
| **Syntax simplicity** | Short                             | Slightly longer                 |
| **Can have methods?** | ❌ No                             | ✅ Yes                          |

## Example 5: Methods on Structs

```go
package main

import "fmt"

// Define a struct
type rectangle struct {
	length, breadth float32
}

// Method with value receiver (works on a copy)
func (r rectangle) area() float32 {
	return r.length * r.breadth
}

// Method with pointer receiver (works on original struct)
func (r *rectangle) perimeter(l float32, b float32) float32 {
	r.length = l
	r.breadth = b
	return 2 * (r.length + r.breadth)
}

func main() {
	// Create struct instance
	rectangeArea := rectangle{10, 20}

	// Calls value receiver method → no change to original struct
	fmt.Println(rectangeArea.area()) // Output: 200

	// Calls pointer receiver method → updates struct values
	fmt.Println(rectangeArea.perimeter(5, 4)) // Output: 18

	// Values updated → new area = 5 * 4 = 20
	fmt.Println(rectangeArea.area()) // Output: 20
}

```

_Go doesn’t have constructors like in other OOP languages (e.g., Java, C++). But we can create our own function that returns a new struct instance — and that acts like a constructor._

### Example 1: Basic Constructor Function

```go
package main

import "fmt"

// Define a struct type 'car'
type car struct {
	name  string
	price int
}

// Constructor function that returns a pointer to a new 'car' instance
func NewCar(name string, price int)car {
	return car{name, price}
}

func main() {
	// Create a new car using the constructor
	car1 := NewCar("hyndai", 23)

	// Prints the pointer to the struct
	fmt.Println(car1) // Output: &{hyndai 23}
}
```

- Humne _NewCar()_ function banaya jo ek _car_ struct return karta hai.

- Ye function constructor ki tarah behave karta hai.

- Naming convention: constructor functions usually start with New (like NewPerson, NewCar, etc.)

<br/>

# Usually, constructors in Go return a pointer to the struct instead of a value — this avoids copying and allows modification.

```go
package main

import "fmt"

// Define a struct 'Car' (exported because it starts with a capital letter)
type Car struct {
	Brand string
	Year  int
}

// Constructor function (idiomatic Go style)
// Returns a pointer to a new 'Car' instance
func NewCar(brand string, year int) *Car {
	return &Car{Brand: brand, Year: year}
}

func main() {
	// Create a new car using the constructor
	c := NewCar("Tesla", 2025)

	// Access and print struct fields
	fmt.Println(c.Brand, c.Year) // Output: Tesla 2025
}
```

# Adding Validation & Defaults in Constructor

A constructor function can also validate input, set default values, or do setup work before returning the struct.

```go
package main

import (
	"errors"
	"fmt"
)

// Account represents a simple bank account structure
type Account struct {
	Owner   string
	Balance float32
}

// NewAccount is a constructor function that validates input
// and returns a pointer to a new Account instance if valid.
func NewAccount(owner string, balance float32) (*Account, error) {
	// Validation: balance should not be negative
	if balance < 0 {
		return nil, errors.New("balance can't be negative")
	}

	// Validation: owner's name should be at least 2 characters long
	if len(owner) < 2 {
		return nil, errors.New("name length can't be below 2 characters")
	}

	// Return a pointer to the new Account if everything is valid
	return &Account{Owner: owner, Balance: balance}, nil
}

func main() {
	// Try creating a new account using the constructor
	ramAccount, err := NewAccount("Ashrith", 1.0)

	// Handle possible errors from constructor
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the successfully created account
	fmt.Println("New Account created:", *ramAccount)
}

```

# Important - struct tags

Struct tags are metadata attached to fields — usually for encoding/decoding (e.g. JSON, XML, DB).

```go
package main

import (
	"encoding/json" // For encoding and decoding JSON
	"fmt"           // For formatted input/output
)

// User struct defines how data will look in Go and how it maps to JSON keys.
// Struct tags (inside backticks) specify custom key names for JSON.
type User struct {
	Name string `json:"full_name"` // Go field Name → JSON key "full_name"
	Age  int    `json:"age"`       // Go field Age → JSON key "age"
}

func main() {
	// ---------------------------
	// 🧩 1. Go Struct → JSON (Marshal)
	// ---------------------------

	// Create a User instance
	newUser := User{"Ashrith", 23}

	// Convert Go struct into JSON (marshal)
	data, err := json.Marshal(newUser)

	// If we print 'data' directly, it will show byte (ASCII) values because the return type is []byte.
	// To see the actual JSON string, we need to convert it using string(data).

	if err != nil {
		fmt.Println("Error while marshalling:", err)
		return
	}

	// Print JSON as string
	fmt.Println("✅ JSON Output (Marshal):")
	fmt.Println(string(data)) // {"full_name":"Ashrith","age":23}

	// ---------------------------
	// 🔄 2. JSON → Go Struct (Unmarshal)
	// ---------------------------

	// Sample JSON input (from API, file, etc.)
	jsonInput := `{"full_name":"Aman","age":25}`

	// Create a variable to store the decoded data
	var userFromJSON User

	// Convert JSON string into Go struct (unmarshal)
	err = json.Unmarshal([]byte(jsonInput), &userFromJSON)
	if err != nil {
		fmt.Println("Error while unmarshalling:", err)
		return
	}

	// Print the resulting Go struct
	fmt.Println("\n✅ Go Struct Output (Unmarshal):")
	fmt.Printf("%+v\n", userFromJSON) // {Name:Aman Age:25}
}
```

### output

```cpp
✅ JSON Output (Marshal):
{"full_name":"Ashrith","age":23}

✅ Go Struct Output (Unmarshal):
{Name:Aman Age:25}

```

| Concept             | Explanation                                                         | Example / Tip                                 |
| ------------------- | ------------------------------------------------------------------- | --------------------------------------------- |
| **Structs**         | Custom data types to group related fields                           | `type Person struct { Name string; Age int }` |
| **JSON Tags**       | Control JSON field names during encode/decode                       | `` `json:"custom_name"` ``                    |
| **Marshalling**     | Converting Go → JSON                                                | `json.Marshal(p)`                             |
| **Unmarshalling**   | Converting JSON → Go                                                | `json.Unmarshal(data, &p)`                    |
| **Exported Fields** | JSON works **only** with fields that start with **capital letters** | ✅ `Name` ❌ `name`                             |
| **Error Handling**  | Always handle `err` properly                                        | `data, err := json.Marshal(p)`                |
| **Byte Slice**      | JSON data is returned as `[]byte`, not string                       | Convert using `string(data)`                  |
| **Packages**        | Always import `encoding/json` for JSON operations                   | `import "encoding/json"`                      |

# Q: JSON works only with fields that start with capital letters why ?
## 1. The Rule

In Go, only exported fields (those starting with a capital letter) are visible to other packages — including the encoding/json package.

So, if your struct field name starts with a lowercase letter,
👉 it is unexported (private), and encoding/json cannot access it.

## 2. Why This Happens (Go’s Visibility Rules)

Go’s visibility system is package-based:

| Field Name | Visibility                                  | Example            |
| ---------- | ------------------------------------------- | ------------------ |
| `Name`     | ✅ Exported (visible outside the package)    | JSON can see it    |
| `name`     | ❌ Unexported (private to this file/package) | JSON cannot see it |

The encoding/json package works outside your code’s package,
so it can only access exported fields — those starting with uppercase letters.

## 3. Example: What Happens in Code
### ❌ Case 1 — Lowercase fields (unexported)
```go
type Person struct {
    name string
    age  int
}

func main() {
    p := Person{name: "Aman", age: 25}
    data, _ := json.Marshal(p)
    fmt.Println(string(data))
}
```


### Output:
```cpp
{}
```

Because encoding/json can’t see name or age (they’re private).

### Case 2 — Uppercase fields (exported)
```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Aman", Age: 25}
    data, _ := json.Marshal(p)
    fmt.Println(string(data))
}
```
```cpp
Output:

{"Name":"Aman","Age":25}
```

Now json.Marshal can access and encode both fields.

| Field Name | Exported | Visible to JSON | Output in JSON |
| ---------- | -------- | --------------- | -------------- |
| `Name`     | ✅ Yes    | ✅ Yes           | ✅ Appears      |
| `name`     | ❌ No     | ❌ No            | ❌ Missing      |
