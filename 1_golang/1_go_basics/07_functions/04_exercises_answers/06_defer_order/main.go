package main

import (
	"fmt"
)

// getSecretKey возвращает secretKey, увеличивая его в defer
func getSecretKey(base int) (secretKey int) {
	defer func() {
		secretKey += 10
	}()
	return base
}

func main() {
	fmt.Println("--- Проверка 06_defer_order (Решение) ---")

	res := getSecretKey(100)

	if res == 110 {
		fmt.Println("Тест модификации в defer: PASS")
	} else {
		fmt.Printf("Тест модификации в defer: FAIL (получено %d)\n", res)
	}
}
