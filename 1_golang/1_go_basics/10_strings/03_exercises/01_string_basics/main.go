package main

import (
	"fmt"
)

// countBytesAndRunes возвращает количество байтов и количество символов (рун) в строке s.
func countBytesAndRunes(s string) (bytes int, runes int) {
	// TODO: Реализуйте подсчет.
	return 0, 0
}

func main() {
	fmt.Println("--- Проверка 01_string_basics ---")

	b1, r1 := countBytesAndRunes("Привет")
	b2, r2 := countBytesAndRunes("Go")

	t1 := b1 == 12 && r1 == 6 // "Привет" -> 12 байт, 6 рун
	t2 := b2 == 2 && r2 == 2  // "Go" -> 2 байта, 2 руны

	if t1 && t2 {
		fmt.Println("Тест countBytesAndRunes: PASS")
	} else {
		fmt.Printf("Тест countBytesAndRunes: FAIL (b1=%d, r1=%d, b2=%d, r2=%d)\n", b1, r1, b2, r2)
	}
}
