package main

import (
	"fmt"
)

// firstRune возвращает первую РУНУ (Unicode code point) строки s.
// Если строка пустая, верните 0.
func firstRune(s string) rune {
	if len(s) == 0 {
		return 0
	}

	// Способ 1: Преобразование строки в слайс рун.
	// Это самый простой способ, но он создает новый слайс []rune в памяти (аллокация).
	// runes := []rune(s)
	// return runes[0]

	// Способ 2: Использование конструкции range.
	// range по строке в Go автоматически декодирует символы UTF-8.
	// Это более производительно, так как не аллоцирует память под весь слайс рун.
	for _, r := range s {
		return r // возвращаем первую же руну
	}

	return 0
}

func main() {
	fmt.Println("--- Проверка 04_rune_utf8 (Решение) ---")

	tests := []struct {
		input string
		want  rune
	}{
		{"Go", 'G'},
		{"Привет", 'П'}, // 'П' - это руна 1055, но в байтах s[0] == 208
		{"", 0},
	}

	allPassed := true
	for i, tt := range tests {
		got := firstRune(tt.input)
		if got == tt.want {
			fmt.Printf("Тест %d (%q): PASS\n", i+1, tt.input)
		} else {
			fmt.Printf("Тест %d (%q): FAIL (ожидалась руна %d (%c), получена %d (%c))\n",
				i+1, tt.input, tt.want, tt.want, got, got)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
