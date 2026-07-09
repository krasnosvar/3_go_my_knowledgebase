package main

import (
	"fmt"
)

// toggleBit инвертирует (переключает) бит на позиции pos в числе val и возвращает результат.
// Например, если бит был равен 1, он должен стать 0, и наоборот.
func toggleBit(val uint8, pos uint) uint8 {
	// TODO: Инвертируйте бит с помощью побитового исключающего ИЛИ (XOR: ^)
	return val
}

func main() {
	fmt.Println("--- Проверка 08_toggle_bits ---")

	var num uint8 = 0b0000 // Все нули

	// 1. Переключаем 1-й бит (индекс 1) -> ожидается 0b0010 (2)
	num = toggleBit(num, 1)
	t1 := num == 2

	// 2. Снова переключаем 1-й бит -> ожидается 0b0000 (0)
	num = toggleBit(num, 1)
	t2 := num == 0

	if t1 && t2 {
		fmt.Println("Тест toggleBit: PASS")
	} else {
		fmt.Printf("Тест toggleBit: FAIL (t1=%t, t2=%t, num=%04b)\n", t1, t2, num)
	}
}
