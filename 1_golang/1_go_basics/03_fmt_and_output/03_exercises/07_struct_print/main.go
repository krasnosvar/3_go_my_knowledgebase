package main

import (
	"fmt"
)

type Device struct {
	ID   int
	Name string
}

// getGoSyntax возвращает строковое Go-представление структуры
func getGoSyntax(d Device) string {
	// TODO: Верните Go-представление структуры d с помощью fmt.Sprintf и нужного флага.
	return ""
}

func main() {
	fmt.Println("--- Проверка 07_struct_print ---")

	d := Device{ID: 101, Name: "Router"}
	output := getGoSyntax(d)

	// Ожидается полное Go-синтаксическое представление структуры
	expected := `main.Device{ID:101, Name:"Router"}`

	if output == expected {
		fmt.Println("Тест getGoSyntax: PASS")
	} else {
		fmt.Printf("Тест getGoSyntax: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}
}
