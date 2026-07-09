package main

import (
	"fmt"
)

// CapitalizeFirstSymbol принимает строку s, переводит первый символ (руну) в верхний регистр
// и возвращает измененную строку.
// Помните о неизменяемости строк и используйте приведение к []rune.
func CapitalizeFirstSymbol(s string) string {
	// TODO: Реализуйте капитализацию первого символа.
	return ""
}

func main() {
	fmt.Println("--- Проверка 04_immutability ---")

	res1 := CapitalizeFirstSymbol("привет")
	res2 := CapitalizeFirstSymbol("hello")

	t1 := res1 == "Привет"
	t2 := res2 == "Hello"

	if t1 && t2 {
		fmt.Println("Тест CapitalizeFirstSymbol: PASS")
	} else {
		fmt.Printf("Тест CapitalizeFirstSymbol: FAIL (res1=%q, res2=%q)\n", res1, res2)
	}
}
