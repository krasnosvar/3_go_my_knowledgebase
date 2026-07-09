package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// NewUser инициализирует и возвращает указатель на структуру User
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	fmt.Println("--- Проверка 08_constructor (Решение) ---")

	u := NewUser("Alice", 25)

	if u != nil && u.Name == "Alice" && u.Age == 25 {
		fmt.Println("Тест конструктора User: PASS")
	} else {
		fmt.Printf("Тест конструктора User: FAIL (получено: %v)\n", u)
	}
}
