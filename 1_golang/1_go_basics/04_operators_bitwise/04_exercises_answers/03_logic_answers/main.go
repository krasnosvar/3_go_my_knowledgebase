package main

import (
	"fmt"
)

// isAllowed проверяет права
func isAllowed(isAdmin, hasSubscription, isBlocked bool) bool {
	return isAdmin || (hasSubscription && !isBlocked)
}

func main() {
	fmt.Println("--- Проверка 03_logic (Решение) ---")

	t1 := isAllowed(true, false, true)
	t2 := isAllowed(false, true, false)
	t3 := isAllowed(false, true, true)
	t4 := isAllowed(false, false, false)

	if t1 && t2 && !t3 && !t4 {
		fmt.Println("Тест isAllowed: PASS")
	} else {
		fmt.Printf("Тест isAllowed: FAIL (t1=%t, t2=%t, t3=%t, t4=%t)\n", t1, t2, t3, t4)
	}
}
