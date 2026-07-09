package main

import "fmt"

type Claims struct {
	UserID int
	Role   string
}

func main() { c := Claims{UserID: 1, Role: "admin"}; fmt.Println(c.Role) }
