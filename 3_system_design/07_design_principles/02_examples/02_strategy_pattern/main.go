// Пример 02: паттерн Strategy в идиоматичном Go.
//
// Strategy — взаимозаменяемые алгоритмы, которые подставляются в момент
// использования. В Go его выражают двумя идиоматичными способами:
//
//  1. func-тип: алгоритм — это ЗНАЧЕНИЕ-функция (никакой иерархии классов).
//  2. маленький интерфейс с одним методом.
//
// Здесь показаны оба: разные стратегии расчёта скидки (func-тип) и разные
// стратегии сортировки (func-тип, как это делает sort.Slice в stdlib).
//
// Программа детерминированно применяет разные стратегии и завершается.
//
// Запуск: go run ./02_examples/02_strategy_pattern/main.go
package main

import (
	"fmt"
	"sort"
)

// ---------------------------------------------------------------------------
// Вариант 1: стратегия скидки как func-тип.
// ---------------------------------------------------------------------------

// DiscountStrategy — алгоритм расчёта итоговой цены. Это просто функция.
type DiscountStrategy func(price float64) float64

// Конкретные стратегии — обычные функции нужной сигнатуры.
func noDiscount(price float64) float64    { return price }
func tenPercentOff(price float64) float64 { return price * 0.90 }
func fixedFiveOff(price float64) float64  { return price - 5 }

// checkout не знает, КАКОЙ алгоритм подставлен — он зависит от абстракции.
func checkout(price float64, strategy DiscountStrategy) float64 {
	return strategy(price)
}

// ---------------------------------------------------------------------------
// Вариант 2: стратегия сортировки (как sort.Slice — less-функция).
// ---------------------------------------------------------------------------

func sortWith(nums []int, less func(a, b int) bool) []int {
	out := append([]int(nil), nums...) // копия, чтобы не менять исходный слайс
	sort.Slice(out, func(i, j int) bool { return less(out[i], out[j]) })
	return out
}

func main() {
	const price = 100.0

	fmt.Println("=== Strategy через func-тип: скидки ===")
	strategies := []struct {
		name string
		fn   DiscountStrategy
	}{
		{"без скидки", noDiscount},
		{"минус 10%", tenPercentOff},
		{"минус 5 фикс.", fixedFiveOff},
	}
	for _, s := range strategies {
		fmt.Printf("%-14s цена %.2f -> %.2f\n", s.name, price, checkout(price, s.fn))
	}

	fmt.Println()
	fmt.Println("=== Strategy через less-функцию: сортировка ===")
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	asc := sortWith(nums, func(a, b int) bool { return a < b })
	desc := sortWith(nums, func(a, b int) bool { return a > b })
	fmt.Println("исходный:    ", nums)
	fmt.Println("по возрастанию:", asc)
	fmt.Println("по убыванию:   ", desc)

	fmt.Println()
	fmt.Println("Готово, программа завершается.")
}
