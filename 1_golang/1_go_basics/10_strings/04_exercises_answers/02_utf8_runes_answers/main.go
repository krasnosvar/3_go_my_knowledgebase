package main

import (
	"fmt"
	"unicode/utf8"
)

func decodeFirstRune(s string) (r rune, size int) {
	return utf8.DecodeRuneInString(s)
}

func main() {
	fmt.Println("--- Проверка 02_utf8_runes (Решение) ---")

	r1, s1 := decodeFirstRune("Привет")
	r2, s2 := decodeFirstRune("Go")
	r3, s3 := decodeFirstRune("📊Test")

	t1 := r1 == 'П' && s1 == 2
	t2 := r2 == 'G' && s2 == 1
	t3 := r3 == '📊' && s3 == 4

	if t1 && t2 && t3 {
		fmt.Println("Тест decodeFirstRune: PASS")
	} else {
		fmt.Printf("Тест decodeFirstRune: FAIL (r1=%c/%d, r2=%c/%d, r3=%c/%d)\n", r1, s1, r2, s2, r3, s3)
	}
}
