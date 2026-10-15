// Упражнение 02: рост стека горутины (Решение).
package main

import "fmt"

// Горутина стартует со стеком 2 KB (с Go 1.4). При нехватке места runtime
// выделяет стек вдвое больше и копирует туда кадры (contiguous stack).
const initialStackKB = 2

func depth(n int) int {
	var frame [64]byte
	frame[n%len(frame)] = byte(n)
	_ = frame
	if n == 0 {
		return 0
	}
	return depth(n-1) + 1
}

func main() {
	fmt.Println("--- Проверка 02_stack_growth (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(initialStackKB == 2, fmt.Sprintf("минимальный стартовый стек горутины — 2 KB, указано %d", initialStackKB))

	done := make(chan int)
	go func() { done <- depth(200000) }()
	got := <-done
	check(got == 200000, fmt.Sprintf("depth(200000) должен вернуть 200000, получено %d", got))
	check(depth(0) == 0, "depth(0) должен вернуть 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
