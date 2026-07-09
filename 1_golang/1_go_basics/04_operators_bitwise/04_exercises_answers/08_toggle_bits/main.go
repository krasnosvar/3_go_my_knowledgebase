package main

import (
	"fmt"
)

// toggleBit инвертирует (переключает) бит на позиции pos в числе val
func toggleBit(val uint8, pos uint) uint8 {
	return val ^ (1 << pos)
}

func main() {
	fmt.Println("--- Проверка 08_toggle_bits (Решение) ---")

	var num uint8 = 0b0000

	num = toggleBit(num, 1)
	t1 := num == 2

	num = toggleBit(num, 1)
	t2 := num == 0

	if t1 && t2 {
		fmt.Println("Тест toggleBit: PASS")
	} else {
		fmt.Printf("Тест toggleBit: FAIL (t1=%t, t2=%t, num=%04b)\n", t1, t2, num)
	}
}
