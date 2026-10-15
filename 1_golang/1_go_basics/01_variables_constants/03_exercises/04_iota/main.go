package main

import (
	"fmt"
)

// Priority представляет уровень приоритета задачи
type Priority int

// TODO: Объявите константы Low, Medium, High типа Priority с помощью iota так,
// чтобы Low равнялся 1, Medium - 2, High - 3.
//
// ВНИМАНИЕ: ниже стоят заглушки, чтобы программа компилировалась.
// Замените их объявлением через iota.
const (
	Low    Priority = 0
	Medium Priority = 0
	High   Priority = 0
)

func main() {
	fmt.Println("--- Проверка 04_iota ---")

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
