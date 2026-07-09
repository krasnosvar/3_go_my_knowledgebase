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

	tests := []struct {
		input string
		want  byte
	}{
		{"Go", 'G'},
		{"", 0},
		{"Hello", 'H'},
	}

	allPassed := true
	for i, tt := range tests {
		got := firstByte(tt.input)
		if got == tt.want {
			fmt.Printf("Тест %d (%q): PASS\n", i+1, tt.input)
		} else {
			fmt.Printf("Тест %d (%q): FAIL (ожидалось %d, получено %d)\n", i+1, tt.input, tt.want, got)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
