package main

import (
	"fmt"
)

// checkAccess проверяет возраст и статус бана пользователя
func checkAccess(age int, isBanned bool) (bool, string) {
	if age < 18 {
		return false, "underage"
	}
	if isBanned {
		return false, "banned"
	}
	return true, "ok"
}

func main() {
	fmt.Println("--- Проверка 01_if_else (Решение) ---")

	ok1, msg1 := checkAccess(16, false)
	ok2, msg2 := checkAccess(20, true)
	ok3, msg3 := checkAccess(20, false)

	if !ok1 && msg1 == "underage" && !ok2 && msg2 == "banned" && ok3 && msg3 == "ok" {
		fmt.Println("Тест checkAccess: PASS")
	} else {
		fmt.Printf("Тест checkAccess: FAIL (1: %t/%s, 2: %t/%s, 3: %t/%s)\n", ok1, msg1, ok2, msg2, ok3, msg3)
	}
}
