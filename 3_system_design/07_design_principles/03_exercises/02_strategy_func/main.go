// Упражнение 02: Strategy через func-тип.
//
// Задача: реализовать взаимозаменяемые стратегии как значения-функции
// (идиоматичный Go — без иерархии классов).
//
// PriceStrategy — функция, считающая итоговую цену по базовой.
// Реализуйте две стратегии нужной сигнатуры:
//
//   - regular   — без изменения цены (возвращает base как есть);
//   - blackFriday — скидка 25% (возвращает base * 0.75).
//
// Функция apply просто применяет переданную стратегию.
//
// Заготовка компилируется и печатает FAIL — доведите до PASS.
//
// Запуск: go run ./03_exercises/02_strategy_func/main.go
package main

import "fmt"

// PriceStrategy — стратегия расчёта цены (алгоритм как значение).
type PriceStrategy func(base float64) float64

// apply применяет стратегию к базовой цене.
func apply(base float64, s PriceStrategy) float64 { return s(base) }

// regular — стратегия без скидки.
// TODO: верните цену без изменений.
var regular PriceStrategy = func(base float64) float64 {
	return 0 // TODO
}

// blackFriday — стратегия со скидкой 25%.
// TODO: верните base * 0.75.
var blackFriday PriceStrategy = func(base float64) float64 {
	return 0 // TODO
}

func main() {
	fmt.Println("--- Проверка 02_strategy_func ---")

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
