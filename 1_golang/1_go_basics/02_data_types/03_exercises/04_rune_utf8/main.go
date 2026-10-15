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

	allPassed := true

	// Тест 1
	got := firstRune("Go")
	if got == 'G' {
		fmt.Printf("Тест 1 (%q): PASS\n", "Go")
	} else {
		fmt.Printf("Тест 1 (%q): FAIL (ожидалось %c, получено %c)\n", "Go", 'G', got)
		allPassed = false
	}

	// Тест 2: 'П' - это руна 1055, но в байтах s[0] == 208
	got = firstRune("Привет")
	if got == 'П' {
		fmt.Printf("Тест 2 (%q): PASS\n", "Привет")
	} else {
		fmt.Printf("Тест 2 (%q): FAIL (ожидалось %c, получено %c)\n", "Привет", 'П', got)
		allPassed = false
	}

	// Тест 3
	got = firstRune("")
	if got == rune(0) {
		fmt.Printf("Тест 3 (%q): PASS\n", "")
	} else {
		fmt.Printf("Тест 3 (%q): FAIL (ожидалось %c, получено %c)\n", "", rune(0), got)
		allPassed = false
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
