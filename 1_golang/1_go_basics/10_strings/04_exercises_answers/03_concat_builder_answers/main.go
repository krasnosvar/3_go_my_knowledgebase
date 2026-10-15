package main

import (
	"fmt"
	"strings"
)

func joinStrings(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var b strings.Builder
	// Считаем общую длину с учетом пробелов для Grow
	totalLen := 0
	for _, w := range words {
		totalLen += len(w)
	}
	totalLen += len(words) - 1 // пробелы
	b.Grow(totalLen)

	for i, w := range words {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(w)
	}
	return b.String()
}

func main() {
	fmt.Println("--- Проверка 03_concat_builder (Решение) ---")

	words := []string{"Go", "is", "awesome", "and", "fast"}
	res := joinStrings(words)

	expected := "Go is awesome and fast"
	if res == expected {
		fmt.Println("Тест joinStrings: PASS")
	} else {
		fmt.Printf("Тест joinStrings: FAIL (получено %q)\n", res)
	}
}
