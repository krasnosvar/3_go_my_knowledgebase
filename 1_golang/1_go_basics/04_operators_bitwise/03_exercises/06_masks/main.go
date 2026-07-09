package main

import (
	"fmt"
)

// isBitSet проверяет, установлен ли бит на позиции pos (0-индексированный) в числе val.
// Подсказка: создайте маску сдвигом (1 << pos) и выполните операцию побитового И (&).
func isBitSet(val uint8, pos uint) bool {
	// TODO: Верните true, если бит установлен, иначе false.
	return false
}

func main() {
	fmt.Println("--- Проверка 06_masks ---")

	var num uint8 = 0b00101100 // Установлены биты 2, 3, 5

	t1 := isBitSet(num, 2) // true
	t2 := isBitSet(num, 3) // true
	t3 := isBitSet(num, 4) // false
	t4 := isBitSet(num, 5) // true

	if t1 && t2 && !t3 && t4 {
		fmt.Println("Тест isBitSet: PASS")
	} else {
		fmt.Printf("Тест isBitSet: FAIL (t1=%t, t2=%t, t3=%t, t4=%t)\n", t1, t2, t3, t4)
	}
}
