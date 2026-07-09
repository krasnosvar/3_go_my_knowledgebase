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

// TODO: Объявите структуру Employee, которая встраивает структуру Person.
// Она должна также иметь дополнительное поле Position (string).

func main() {
	fmt.Println("--- Проверка 05_embedded_struct ---")

	// ВНИМАНИЕ: Раскомментируйте код ниже после встраивания структуры Person.

	/*
		emp := Employee{
			Person:   Person{Name: "Дмитрий"},
			Position: "Go Developer",
		}

		// Проверяем доступ к полю встроенной структуры напрямую
		nameOk := emp.Name == "Дмитрий"

		// Проверяем вызов унаследованного метода Greet
		greetOk := emp.Greet() == "Привет, меня зовут Дмитрий"

		if nameOk && greetOk {
			fmt.Println("Тест встраивания: PASS")
		} else {
			fmt.Printf("Тест встраивания: FAIL (nameOk=%t, greetOk=%t)\n", nameOk, greetOk)
		}
	*/

	fmt.Println("Раскомментируйте проверку в main")
}
