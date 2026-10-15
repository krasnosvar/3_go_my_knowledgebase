package main

import (
	"fmt"
)

// TODO: Объявите нетипизированную константу `Pi` со значением 3.14
// TODO: Объявите типизированную константу `MaxCount` типа int32 со значением 100
//
// ВНИМАНИЕ: ниже стоят заглушки, чтобы программа компилировалась.
// Замените их своими объявлениями.
const Pi = 0
const MaxCount int = 0

func main() {
	fmt.Println("--- Проверка 03_const ---")

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
