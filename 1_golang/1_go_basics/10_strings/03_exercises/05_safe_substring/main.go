package main

import (
	"fmt"
)

// SafeSubstring возвращает подстроку от start до end (индексы по символам/рунам, не байтам).
// Если индексы выходят за границы, верните пустую строку или обрежьте по границам строки.
func SafeSubstring(s string, start, end int) string {
	// TODO: Реализуйте безопасную нарезку по рунам.
	return ""
}

func main() {
	fmt.Println("--- Проверка 05_safe_substring ---")

	s := "Привет, мир! 📊" // 15 символов (рун)

	// "ривет" -> символы с 1 по 6
	res1 := SafeSubstring(s, 1, 6)

	// "📊" -> символ 14 (индекс 13)
	res2 := SafeSubstring(s, 13, 14)

	t1 := res1 == "ривет"
	t2 := res2 == "📊"

	if t1 && t2 {
		fmt.Println("Тест SafeSubstring: PASS")
	} else {
		fmt.Printf("Тест SafeSubstring: FAIL (res1=%q, res2=%q)\n", res1, res2)
	}
}
