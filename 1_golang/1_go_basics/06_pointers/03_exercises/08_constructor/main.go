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
	// TODO: Создайте и верните указатель на структуру User с переданными значениями.
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_constructor ---")

	u := NewUser("Alice", 25)

	if u != nil && u.Name == "Alice" && u.Age == 25 {
		fmt.Println("Тест конструктора User: PASS")
	} else {
		fmt.Printf("Тест конструктора User: FAIL (получено: %v)\n", u)
	}
}
