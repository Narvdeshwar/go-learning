package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	name      string
	price     int
	createdAt time.Time
}

func newOrder(id int, name string, price int, createdAt time.Time) order {
	return order{id: id, name: name, price: price, createdAt: createdAt}
}

func main() {
	orders := order{
		id:        1,
		name:      "tea",
		price:     35,
		createdAt: time.Now(),
	}
	fmt.Println(newOrder(1, "ram", 34, time.Now()))

	fmt.Println(orders)

}
