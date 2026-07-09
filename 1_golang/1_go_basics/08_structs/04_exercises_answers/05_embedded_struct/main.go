package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Привет, меня зовут " + p.Name
}

type Employee struct {
	Person
	Position string
}

func main() {
	fmt.Println("--- Проверка 05_embedded_struct (Решение) ---")

	emp := Employee{
		Person:   Person{Name: "Дмитрий"},
		Position: "Go Developer",
	}

	nameOk := emp.Name == "Дмитрий"
	greetOk := emp.Greet() == "Привет, меня зовут Дмитрий"

	if nameOk && greetOk {
		fmt.Println("Тест встраивания: PASS")
	} else {
		fmt.Printf("Тест встраивания: FAIL (nameOk=%t, greetOk=%t)\n", nameOk, greetOk)
	}
}
