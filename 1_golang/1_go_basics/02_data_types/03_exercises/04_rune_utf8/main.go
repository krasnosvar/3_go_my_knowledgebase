package main

import (
	"fmt"
)

// firstRune возвращает первую РУНУ (Unicode code point) строки s.
// Если строка пустая, верните 0.
// ВНИМАНИЕ: помните, что первый символ кириллицы занимает 2 байта!
func firstRune(s string) rune {
	// TODO: реализуйте функцию
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_rune_utf8 ---")

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
