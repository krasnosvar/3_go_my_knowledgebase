// Упражнение 02: Strategy через func-тип (Решение).
package main

import "fmt"

type PriceStrategy func(base float64) float64

func apply(base float64, s PriceStrategy) float64 { return s(base) }

var regular PriceStrategy = func(base float64) float64 {
	return base
}

var blackFriday PriceStrategy = func(base float64) float64 {
	return base * 0.75
}

func main() {
	fmt.Println("--- Проверка 02_strategy_func (Решение) ---")

	const base = 200.0
	allOK := true

	if got := apply(base, regular); got != 200.0 {
		allOK = false
		fmt.Printf("FAIL: regular(%.0f) = %.2f, ожидалось 200.00\n", base, got)
	}
	if got := apply(base, blackFriday); got != 150.0 {
		allOK = false
		fmt.Printf("FAIL: blackFriday(%.0f) = %.2f, ожидалось 150.00\n", base, got)
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
