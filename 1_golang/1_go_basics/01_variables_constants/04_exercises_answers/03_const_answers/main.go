package main

import (
	"fmt"
)

// Объявление нетипизированной константы
const Pi = 3.14

// Объявление типизированной константы
const MaxCount int32 = 100

func main() {
	fmt.Println("--- Проверка 03_const (Решение) ---")

	// %T печатает тип значения (подробнее о fmt — в теме 03_fmt_and_output)
	fmt.Printf("Pi: %v (тип: %T)\n", Pi, Pi)
	fmt.Printf("MaxCount: %v (тип: %T)\n", MaxCount, MaxCount)

	if Pi == 3.14 {
		fmt.Println("Тест Pi: PASS")
	} else {
		fmt.Println("Тест Pi: FAIL")
	}

	if MaxCount == 100 && fmt.Sprintf("%T", MaxCount) == "int32" {
		fmt.Println("Тест MaxCount: PASS")
	} else {
		fmt.Println("Тест MaxCount: FAIL")
	}
}
