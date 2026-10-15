package main

import (
	"fmt"
)

func SafeSubstring(s string, start, end int) string {
	runes := []rune(s)
	if start < 0 {
		start = 0
	}
	if end > len(runes) {
		end = len(runes)
	}
	if start >= end || start >= len(runes) {
		return ""
	}
	return string(runes[start:end])
}

func main() {
	fmt.Println("--- Проверка 05_safe_substring (Решение) ---")

	s := "Привет, мир! 📊"

	res1 := SafeSubstring(s, 1, 6)
	res2 := SafeSubstring(s, 13, 14)

	t1 := res1 == "ривет"
	t2 := res2 == "📊"

	if t1 && t2 {
		fmt.Println("Тест SafeSubstring: PASS")
	} else {
		fmt.Printf("Тест SafeSubstring: FAIL (res1=%q, res2=%q)\n", res1, res2)
	}
}
