package main

import (
	"fmt"
)

// canAccess возвращает true только если пользователь является администратором (admin)
// И при этом его аккаунт активен (active).
func canAccess(admin, active bool) bool {
	// Используем логическое И (&&)
	return admin && active
}

func main() {
	fmt.Println("--- Проверка 02_bool_logic (Решение) ---")

	tests := []struct {
		admin  bool
		active bool
		want   bool
	}{
		{true, true, true},
		{true, false, false},
		{false, true, false},
		{false, false, false},
	}

	allPassed := true
	for i, tt := range tests {
		got := canAccess(tt.admin, tt.active)
		if got == tt.want {
			fmt.Printf("Тест %d (admin=%t, active=%t): PASS\n", i+1, tt.admin, tt.active)
		} else {
			fmt.Printf("Тест %d (admin=%t, active=%t): FAIL (ожидалось %t, получено %t)\n", i+1, tt.admin, tt.active, tt.want, got)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("Все тесты пройдены успешно!")
	}
}
