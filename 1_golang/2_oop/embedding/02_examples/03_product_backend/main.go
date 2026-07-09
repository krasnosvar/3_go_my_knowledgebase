package main

import "fmt"

type Timestamps struct {
	CreatedAt string
	UpdatedAt string
}

type User struct {
	Timestamps
	Name string
}

func main() {
	u := User{Timestamps: Timestamps{CreatedAt: "today"}, Name: "Den"}
	fmt.Println(u.Name, u.CreatedAt)
}
