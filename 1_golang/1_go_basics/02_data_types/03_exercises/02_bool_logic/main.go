package main

import (
	"fmt"
)

// canAccess возвращает true только если пользователь является администратором (admin)
// И при этом его аккаунт активен (active).
func canAccess(admin, active bool) bool {
	// TODO: реализуйте функцию
	return false
}

func main() {
	fmt.Println("--- Проверка 02_bool_logic ---")

	allPassed := true

	// Тест 1
	got := canAccess(true, true)
	if got == true {
		fmt.Printf("Тест 1 (admin=%t, active=%t): PASS\n", true, true)
	} else {
		fmt.Printf("Тест 1 (admin=%t, active=%t): FAIL (ожидалось %t, получено %t)\n", true, true, true, got)
		allPassed = false
	}

	// Тест 2
	got = canAccess(true, false)
	if got == false {
		fmt.Printf("Тест 2 (admin=%t, active=%t): PASS\n", true, false)
	} else {
		fmt.Printf("Тест 2 (admin=%t, active=%t): FAIL (ожидалось %t, получено %t)\n", true, false, false, got)
		allPassed = false
	}

	// Тест 3
	got = canAccess(false, true)
	if got == false {
		fmt.Printf("Тест 3 (admin=%t, active=%t): PASS\n", false, true)
	} else {
		fmt.Printf("Тест 3 (admin=%t, active=%t): FAIL (ожидалось %t, получено %t)\n", false, true, false, got)
		allPassed = false
	}

	// Тест 4
	got = canAccess(false, false)
	if got == false {
		fmt.Printf("Тест 4 (admin=%t, active=%t): PASS\n", false, false)
	} else {
		fmt.Printf("Тест 4 (admin=%t, active=%t): FAIL (ожидалось %t, получено %t)\n", false, false, false, got)
		allPassed = false
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
