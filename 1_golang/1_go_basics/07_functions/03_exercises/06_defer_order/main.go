package main

import (
	"fmt"
)

// getSecretKey возвращает secretKey, но с помощью defer перед выходом из функции
// увеличивает возвращаемое значение на 10.
// Используйте именованное возвращаемое значение в сигнатуре: (secretKey int).
func getSecretKey(base int) (secretKey int) {
	// TODO: Объявите defer, который увеличит secretKey на 10.

	return base
}

func main() {
	fmt.Println("--- Проверка 06_defer_order ---")

	res := getSecretKey(100) // Ожидается 110

	if res == 110 {
		fmt.Println("Тест модификации в defer: PASS")
	} else {
		fmt.Printf("Тест модификации в defer: FAIL (получено %d)\n", res)
	}
}
