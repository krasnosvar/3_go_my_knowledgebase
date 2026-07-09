package main

import (
	"fmt"
)

// performBitwise возвращает побитовое И, ИЛИ, XOR двух чисел
func performBitwise(a, b uint8) (uint8, uint8, uint8) {
	return a & b, a | b, a ^ b
}

func main() {
	fmt.Println("--- Проверка 04_bitwise_and_or_xor (Решение) ---")

	and, or, xor := performBitwise(0b1100, 0b1010)

	if and == 0b1000 && or == 0b1110 && xor == 0b0110 {
		fmt.Println("Тест performBitwise: PASS")
	} else {
		fmt.Printf("Тест performBitwise: FAIL (and=%04b, or=%04b, xor=%04b)\n", and, or, xor)
	}
}
