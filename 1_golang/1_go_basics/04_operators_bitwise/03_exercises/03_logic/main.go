package main

import (
	"fmt"
)

// isAllowed проверяет права: пользователь должен быть администратором ИЛИ (иметь активную подписку И не быть заблокированным)
func isAllowed(isAdmin, hasSubscription, isBlocked bool) bool {
	// TODO: Напишите логическое выражение
	return false
}

func main() {
	fmt.Println("--- Проверка 03_logic ---")

	// Тесты
	t1 := isAllowed(true, false, true)   // Админ, заблокирован -> true (админу можно всё)
	t2 := isAllowed(false, true, false)  // Не админ, подписка, не заблокирован -> true
	t3 := isAllowed(false, true, true)   // Не админ, подписка, заблокирован -> false
	t4 := isAllowed(false, false, false) // Не админ, нет подписки -> false

	if t1 && t2 && !t3 && !t4 {
		fmt.Println("Тест isAllowed: PASS")
	} else {
		fmt.Printf("Тест isAllowed: FAIL (t1=%t, t2=%t, t3=%t, t4=%t)\n", t1, t2, t3, t4)
	}
}
