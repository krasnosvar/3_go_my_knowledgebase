package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	b, _ := json.Marshal(User{Name: "Den", Age: 30})
	var u User
	_ = json.Unmarshal(b, &u)
	fmt.Println(string(b), u.Name)
}
