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
