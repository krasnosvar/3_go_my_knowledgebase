package main

import (
	"fmt"
)

// isZeroValue проверяет, что все три параметра конфигурации имеют
// дефолтные нулевые значения своих типов (zero values).
func isZeroValue(enabled bool, count int, name string) bool {
	// TODO: реализуйте функцию
	return false
}

func main() {
	fmt.Println("--- Проверка 06_zero_values ---")

	allPassed := true

	// Тест 1
	got := isZeroValue(false, 0, "")
	if got == true {
		fmt.Printf("Тест 1 (enabled=%t, count=%d, name=%q): PASS\n", false, 0, "")
	} else {
		fmt.Printf("Тест 1 (enabled=%t, count=%d, name=%q): FAIL (ожидалось %t, получено %t)\n", false, 0, "", true, got)
		allPassed = false
	}

	// Тест 2
	got = isZeroValue(true, 0, "")
	if got == false {
		fmt.Printf("Тест 2 (enabled=%t, count=%d, name=%q): PASS\n", true, 0, "")
	} else {
		fmt.Printf("Тест 2 (enabled=%t, count=%d, name=%q): FAIL (ожидалось %t, получено %t)\n", true, 0, "", false, got)
		allPassed = false
	}

	// Тест 3
	got = isZeroValue(false, 5, "")
	if got == false {
		fmt.Printf("Тест 3 (enabled=%t, count=%d, name=%q): PASS\n", false, 5, "")
	} else {
		fmt.Printf("Тест 3 (enabled=%t, count=%d, name=%q): FAIL (ожидалось %t, получено %t)\n", false, 5, "", false, got)
		allPassed = false
	}

	// Тест 4
	got = isZeroValue(false, 0, "Admin")
	if got == false {
		fmt.Printf("Тест 4 (enabled=%t, count=%d, name=%q): PASS\n", false, 0, "Admin")
	} else {
		fmt.Printf("Тест 4 (enabled=%t, count=%d, name=%q): FAIL (ожидалось %t, получено %t)\n", false, 0, "Admin", false, got)
		allPassed = false
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
