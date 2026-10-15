package main

import (
	"fmt"
)

// firstByte возвращает первый БАЙТ строки.
// Если строка пустая, верните 0.
func firstByte(s string) byte {
	// ВАЖНО: Всегда проверяйте длину строки перед получением элемента по индексу s[0].
	// Обращение к индексу пустой строки вызовет panic: runtime error: index out of range.
	if len(s) == 0 {
		return 0
	}
	return s[0]
}

func main() {
	fmt.Println("--- Проверка 03_string_bytes (Решение) ---")

	allPassed := true

	// Тест 1
	got := firstByte("Go")
	if got == byte('G') {
		fmt.Printf("Тест 1 (%q): PASS\n", "Go")
	} else {
		fmt.Printf("Тест 1 (%q): FAIL (ожидалось %d, получено %d)\n", "Go", byte('G'), got)
		allPassed = false
	}

	// Тест 2
	got = firstByte("")
	if got == byte(0) {
		fmt.Printf("Тест 2 (%q): PASS\n", "")
	} else {
		fmt.Printf("Тест 2 (%q): FAIL (ожидалось %d, получено %d)\n", "", byte(0), got)
		allPassed = false
	}

	// Тест 3
	got = firstByte("Hello")
	if got == byte('H') {
		fmt.Printf("Тест 3 (%q): PASS\n", "Hello")
	} else {
		fmt.Printf("Тест 3 (%q): FAIL (ожидалось %d, получено %d)\n", "Hello", byte('H'), got)
		allPassed = false
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
