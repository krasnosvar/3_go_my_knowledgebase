package main

import (
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Rating float64
}

// newInception создаёт фильм «Inception».
func newInception() Movie {
	return Movie{Title: "Inception", Year: 2010, Rating: 8.8}
}

func main() {
	fmt.Println("--- Проверка 01_declaration (Решение) ---")

	m := newInception()

	// %+v печатает структуру вместе с именами полей.
	got := fmt.Sprintf("%+v", m)
	if got == "{Title:Inception Year:2010 Rating:8.8}" {
		fmt.Println("Тест Movie: PASS")
	} else {
		fmt.Printf("Тест Movie: FAIL (получено %s)\n", got)
	}
}
