package main

import (
	"fmt"
)

// joinStrings объявляет strings.Builder, вызывает метод Grow для резервирования памяти
// и собирает переданные строки, разделяя их пробелом.
func joinStrings(words []string) string {
	// TODO: Реализуйте сборку строк через strings.Builder с Grow.
	return ""
}

func main() {
	fmt.Println("--- Проверка 03_concat_builder ---")

	words := []string{"Go", "is", "awesome", "and", "fast"}
	res := joinStrings(words)

	expected := "Go is awesome and fast"
	if res == expected {
		fmt.Println("Тест joinStrings: PASS")
	} else {
		fmt.Printf("Тест joinStrings: FAIL (получено %q)\n", res)
	}
}
