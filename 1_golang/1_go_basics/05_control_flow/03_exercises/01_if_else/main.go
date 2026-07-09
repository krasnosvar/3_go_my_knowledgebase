package main

import (
	"fmt"
)

// checkAccess проверяет возраст и статус бана пользователя.
// Если возраст меньше 18 — вернуть false, "underage".
// Если пользователь забанен — вернуть false, "banned".
// Иначе — вернуть true, "ok".
// Используйте ранний возврат (guard clauses).
func checkAccess(age int, isBanned bool) (bool, string) {
	// TODO: Реализуйте проверки.
	return false, ""
}

func main() {
	fmt.Println("--- Проверка 01_if_else ---")

	ok1, msg1 := checkAccess(16, false) // false, "underage"
	ok2, msg2 := checkAccess(20, true)  // false, "banned"
	ok3, msg3 := checkAccess(20, false) // true, "ok"

	if !ok1 && msg1 == "underage" && !ok2 && msg2 == "banned" && ok3 && msg3 == "ok" {
		fmt.Println("Тест checkAccess: PASS")
	} else {
		fmt.Printf("Тест checkAccess: FAIL (1: %t/%s, 2: %t/%s, 3: %t/%s)\n", ok1, msg1, ok2, msg2, ok3, msg3)
	}
}
