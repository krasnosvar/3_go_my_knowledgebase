package main

import (
	"fmt"
)

// isBitSet проверяет, установлен ли бит на позиции pos
func isBitSet(val uint8, pos uint) bool {
	return (val & (1 << pos)) != 0
}

func main() {
	fmt.Println("--- Проверка 06_masks (Решение) ---")

	var num uint8 = 0b00101100

	t1 := isBitSet(num, 2)
	t2 := isBitSet(num, 3)
	t3 := isBitSet(num, 4)
	t4 := isBitSet(num, 5)

	if t1 && t2 && !t3 && t4 {
		fmt.Println("Тест isBitSet: PASS")
	} else {
		fmt.Printf("Тест isBitSet: FAIL (t1=%t, t2=%t, t3=%t, t4=%t)\n", t1, t2, t3, t4)
	}
}
