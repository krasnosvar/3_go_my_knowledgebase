package main

import (
	"fmt"
	"unicode"
)

func CapitalizeFirstSymbol(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func main() {
	fmt.Println("--- Проверка 04_immutability (Решение) ---")

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
