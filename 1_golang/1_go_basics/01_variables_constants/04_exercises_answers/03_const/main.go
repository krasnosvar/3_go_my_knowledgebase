package main

import (
	"fmt"
	"reflect"
)

// Объявление нетипизированной константы
const Pi = 3.14

// Объявление типизированной константы
const MaxCount int32 = 100

func main() {
	fmt.Println("--- Проверка 03_const (Решение) ---")

	fmt.Printf("Pi: %v (тип: %s)\n", Pi, reflect.TypeOf(Pi))
	fmt.Printf("MaxCount: %v (тип: %s)\n", MaxCount, reflect.TypeOf(MaxCount))

	if Pi == 3.14 {
		fmt.Println("Тест Pi: PASS")
	} else {
		fmt.Println("Тест Pi: FAIL")
	}

	if MaxCount == 100 && reflect.TypeOf(MaxCount).String() == "int32" {
		fmt.Println("Тест MaxCount: PASS")
	} else {
		fmt.Println("Тест MaxCount: FAIL")
	}
}
