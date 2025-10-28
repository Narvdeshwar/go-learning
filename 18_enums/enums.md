# Enums

Go does not have native enums (like C, Java, etc.),
but you can simulate enums using const + iota.

Enums are used when you want to represent a fixed set of related constant values — like days of the week, user roles, or order statuses.

Go me enum keyword nahi hota,
lekin hum const aur iota ka use karke enum jaisa structure bana sakte hain.

Enum ka matlab hota hai — ek fixed list of options, jaise:

* Days → Mon, Tue, Wed

* Roles → Admin, User, Guest

```go
package main

import "fmt"

// 'const' block is used to declare constant values in Go.
// Constants are fixed values that cannot be changed later.

// Step 1️⃣: Using 'iota' inside a constant block
// 'iota' is a special Go keyword that automatically generates
// a sequence of incrementing numbers — starting from 0.

// So here:
// Sunday   = 0
// Monday   = 1
// Tuesday  = 2
// Wednesday= 3
// Thursday = 4
// Friday   = 5
// Saturday = 6

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// Step 2️⃣: Print the value of Sunday
	// Since 'Sunday = 0' (because iota starts from 0), this will print 0.
	fmt.Println(Sunday)
}
```

```go
package main

import "fmt"

type Day int

const (
	Sunday Day = iota
	Mondey
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Day) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "WednesDay", "Thursday", "Friday", "Saturday"}[d]
}

func main() {
	fmt.Println(Sunday)
}

```

```go
package main

import "fmt"

type orderStatus int

const (
	pending orderStatus = iota
	received
	confirmed
	preparing
	shipped
	delivered
)

func (o orderStatus) String() string {
	return [...]string{"pending", "received", "confirmed", "preparing", "shipped", "delivered"}[o]
}

func nextStatus(status orderStatus) orderStatus {
	switch status {
	case pending:
		return received
	case received:
		return confirmed
	case confirmed:
		return preparing
	case preparing:
		return shipped
	case shipped:
		return delivered
	default:
		return status
	}
}

func main() {
	order := pending
	fmt.Print(order)
	order = nextStatus(order)
	fmt.Println(order)
}

```