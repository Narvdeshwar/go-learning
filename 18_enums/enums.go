package main

import "fmt"

type Day string

const (
	Sunday    Day = "Sunday"
	Mondey    Day = "Monday"
	Tuesday   Day = "Tuesday"
	Wednesday Day = "Wednesday"
	Thursday  Day = "Thursday"
	Friday    Day = "Friday"
	Saturday  Day = "Saturday"
)

func main() {
	fmt.Println(Mondey)
}
