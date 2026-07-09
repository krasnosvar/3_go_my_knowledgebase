package main

import (
	"fmt"
)

// performBitwise возвращает побитовое И, ИЛИ, XOR двух чисел
func performBitwise(a, b uint8) (uint8, uint8, uint8) {
	// TODO: Вычислите и верните побитовые AND, OR, XOR.
	return 0, 0, 0
}

func main() {
	fmt.Println("--- Проверка 04_bitwise_and_or_xor ---")

	and, or, xor := performBitwise(0b1100, 0b1010) // 12 и 10

	if and == 0b1000 && or == 0b1110 && xor == 0b0110 {
		fmt.Println("Тест performBitwise: PASS")
	} else {
		fmt.Printf("Тест performBitwise: FAIL (and=%04b, or=%04b, xor=%04b)\n", and, or, xor)
	}
}
