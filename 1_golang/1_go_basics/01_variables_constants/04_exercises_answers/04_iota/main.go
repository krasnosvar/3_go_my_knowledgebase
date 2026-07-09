package main

import (
	"fmt"
)

// Priority представляет уровень приоритета задачи
type Priority int

// Объявляем константы Low, Medium, High со значениями 1, 2, 3 соответственно.
// Используем заглушку `_` для пропуска 0.
const (
	_      Priority = iota // 0
	Low                    // 1
	Medium                 // 2
	High                   // 3
)

func main() {
	fmt.Println("--- Проверка 04_iota (Решение) ---")

	if Low == 1 {
		fmt.Println("Тест Low: PASS")
	} else {
		fmt.Printf("Тест Low: FAIL (ожидалось 1, получено %d)\n", Low)
	}

	if Medium == 2 {
		fmt.Println("Тест Medium: PASS")
	} else {
		fmt.Printf("Тест Medium: FAIL (ожидалось 2, получено %d)\n", Medium)
	}

	if High == 3 {
		fmt.Println("Тест High: PASS")
	} else {
		fmt.Printf("Тест High: FAIL (ожидалось 3, получено %d)\n", High)
	}
}
