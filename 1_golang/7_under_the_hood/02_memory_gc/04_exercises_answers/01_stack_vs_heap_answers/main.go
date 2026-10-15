// Упражнение 01: стек вместо кучи (Решение).
package main

import (
	"fmt"
	"testing"
)

// sumSquares возвращает 1*1 + 2*2 + ... + n*n.
func sumSquares(n int) int {
	sum := 0 // локальная переменная: остаётся на стеке
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func main() {
	fmt.Println("--- Проверка 01_stack_vs_heap (Решение) ---")

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
