package main

import (
	"fmt"
)

// TODO: Объявите структуру Movie с тремя экспортируемыми полями:
// Title (string), Year (int), Rating (float64).
// ВНИМАНИЕ: ниже заглушка (пустая структура), чтобы программа компилировалась — замените её.
type Movie struct{}

// newInception создаёт фильм «Inception».
// TODO: после объявления полей верните Movie{Title: "Inception", Year: 2010, Rating: 8.8}.
func newInception() Movie {
	return Movie{}
}

func main() {
	fmt.Println("--- Проверка 01_declaration ---")

	m := newInception()

	// %+v печатает структуру вместе с именами полей.
	got := fmt.Sprintf("%+v", m)
	if got == "{Title:Inception Year:2010 Rating:8.8}" {
		fmt.Println("Тест Movie: PASS")
	} else {
		fmt.Printf("Тест Movie: FAIL (получено %s)\n", got)
	}
}
