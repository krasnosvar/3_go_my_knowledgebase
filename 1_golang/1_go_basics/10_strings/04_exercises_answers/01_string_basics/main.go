package main

import (
	"fmt"
	"unicode/utf8"
)

func countBytesAndRunes(s string) (bytes int, runes int) {
	return len(s), utf8.RuneCountInString(s)
}

func main() {
	fmt.Println("--- Проверка 01_string_basics (Решение) ---")

	b1, r1 := countBytesAndRunes("Привет")
	b2, r2 := countBytesAndRunes("Go")

	t1 := b1 == 12 && r1 == 6
	t2 := b2 == 2 && r2 == 2

	if t1 && t2 {
		fmt.Println("Тест countBytesAndRunes: PASS")
	} else {
		fmt.Printf("Тест countBytesAndRunes: FAIL (b1=%d, r1=%d, b2=%d, r2=%d)\n", b1, r1, b2, r2)
	}
}
