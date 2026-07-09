package main

import (
	"fmt"
)

// decodeFirstRune возвращает первый символ (rune) строки s и его размер в байтах.
// Если строка пустая, верните utf8.RuneError и 0.
func decodeFirstRune(s string) (r rune, size int) {
	// TODO: Реализуйте декодирование с использованием utf8.DecodeRuneInString.
	return 0, 0
}

func main() {
	fmt.Println("--- Проверка 02_utf8_runes ---")

	r1, s1 := decodeFirstRune("Привет") // П -> 2 байта
	r2, s2 := decodeFirstRune("Go")     // G -> 1 байт
	r3, s3 := decodeFirstRune("📊Test")  // 📊 -> 4 байта

	t1 := r1 == 'П' && s1 == 2
	t2 := r2 == 'G' && s2 == 1
	t3 := r3 == '📊' && s3 == 4

	if t1 && t2 && t3 {
		fmt.Println("Тест decodeFirstRune: PASS")
	} else {
		fmt.Printf("Тест decodeFirstRune: FAIL (r1=%c/%d, r2=%c/%d, r3=%c/%d)\n", r1, s1, r2, s2, r3, s3)
	}
}
