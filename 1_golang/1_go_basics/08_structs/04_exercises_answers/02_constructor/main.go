package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func NewUser(id int, name string) *User {
	if name == "" {
		return nil
	}
	return &User{ID: id, Name: name}
}

func main() {
	fmt.Println("--- Проверка 02_constructor (Решение) ---")

	u1 := NewUser(1, "Denis")
	u2 := NewUser(2, "")

	t1 := u1 != nil && u1.ID == 1 && u1.Name == "Denis"
	t2 := u2 == nil

	if t1 && t2 {
		fmt.Println("Тест NewUser: PASS")
	} else {
		fmt.Printf("Тест NewUser: FAIL (u1=%+v, u2=%+v)\n", u1, u2)
	}
}
