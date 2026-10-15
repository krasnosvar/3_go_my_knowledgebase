// Упражнение 01: стек вместо кучи.
//
// Задача: перепишите sumSquares(n) так, чтобы она не делала аллокаций в
// куче. Сейчас она копит квадраты в слайсе через append — при росте
// слайса runtime выделяет новые массивы в куче. Сумму можно
// накопить в обычной локальной переменной — она живёт на стеке.
//
// Проверка: результат верный, а testing.AllocsPerRun показывает 0 аллокаций.
//
// Запуск: go run ./03_exercises/01_stack_vs_heap/main.go
package main

import (
	"fmt"
	"testing"
)

// sumSquares возвращает 1*1 + 2*2 + ... + n*n.
func sumSquares(n int) int {
	// TODO: уберите промежуточный слайс — он и есть аллокация.
	var squares []int
	for i := 1; i <= n; i++ {
		squares = append(squares, i*i)
	}
	sum := 0
	for _, v := range squares {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("--- Проверка 01_stack_vs_heap ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(sumSquares(3) == 14, fmt.Sprintf("sumSquares(3) должна быть 14, получено %d", sumSquares(3)))
	check(sumSquares(0) == 0, "sumSquares(0) должна быть 0")

	n := 1000
	allocs := testing.AllocsPerRun(100, func() { _ = sumSquares(n) })
	check(allocs == 0, fmt.Sprintf("ожидалось 0 аллокаций, получено %.0f", allocs))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
