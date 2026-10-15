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

// TODO: Структура Employee должна ВСТРАИВАТЬ структуру Person (поле без имени)
// и иметь дополнительное поле Position (string).
// ВНИМАНИЕ: ниже заглушка без встраивания: Person — обычное именованное поле,
// а Name и Greet() объявлены вручную и возвращают пустые значения.
// Замените поле `Person Person` на встраивание `Person`, удалите поле Name и метод Greet у Employee —
// тогда emp.Name и emp.Greet() будут «продвинуты» (promoted) из Person.
type Employee struct {
	Person   Person
	Name     string
	Position string
}

func (e Employee) Greet() string {
	return ""
}

func main() {
	fmt.Println("--- Проверка 05_embedded_struct ---")

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
}
