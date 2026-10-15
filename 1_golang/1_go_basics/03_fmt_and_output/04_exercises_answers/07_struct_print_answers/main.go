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
	return fmt.Sprintf("%#v", d)
}

func main() {
	fmt.Println("--- Проверка 07_struct_print (Решение) ---")

	d := Device{ID: 101, Name: "Router"}
	output := getGoSyntax(d)

	expected := `main.Device{ID:101, Name:"Router"}`

	if output == expected {
		fmt.Println("Тест getGoSyntax: PASS")
	} else {
		fmt.Printf("Тест getGoSyntax: FAIL (ожидалось %q, получено %q)\n", expected, output)
	}
}
