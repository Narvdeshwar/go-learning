package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"full_name"`
	Age  int    `json:"age"`
}

func main() {
	newUser := User{"ashrith", 23}
	data, _ := json.Marshal(newUser)
	fmt.Println(string(data))
}
