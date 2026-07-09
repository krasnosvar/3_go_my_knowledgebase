package main

import (
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Rating float64
}

func main() {
	fmt.Println("--- Проверка 01_declaration (Решение) ---")

	m := Movie{Title: "Inception", Year: 2010, Rating: 8.8}

	if m.Title == "Inception" && m.Year == 2010 && m.Rating == 8.8 {
		fmt.Println("Тест Movie: PASS")
	} else {
		fmt.Printf("Тест Movie: FAIL (получено %+v)\n", m)
	}
}
