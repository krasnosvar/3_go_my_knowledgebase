// Упражнение 02: рост стека горутины.
//
// Задача:
//  1. Укажите в константе initialStackKB минимальный стартовый размер стека
//     горутины в Go (в килобайтах) — см. 01_theory/03_stack_and_preemption.md.
//  2. Реализуйте рекурсивную функцию depth(n), которая вызывает себя n раз и
//     возвращает глубину рекурсии (depth(0) == 0, depth(n) == depth(n-1)+1).
//     Локальный массив frame уже добавлен, чтобы каждый кадр занимал место.
//
// Проверка: depth(200000) отрабатывает без переполнения — стек вырос с
// нескольких килобайт до мегабайтов (runtime копирует стек в больший).
//
// Запуск: go run ./03_exercises/02_stack_growth/main.go
package main

import "fmt"

// TODO: подставьте минимальный стартовый размер стека горутины в KB.
const initialStackKB = 0

// depth рекурсивно спускается на n уровней и возвращает глубину.
func depth(n int) int {
	var frame [64]byte // занимает место в кадре стека
	frame[n%len(frame)] = byte(n)
	_ = frame
	// TODO: верните 0 при n == 0, иначе depth(n-1) + 1.
	return 0
}

func main() {
	fmt.Println("--- Проверка 02_stack_growth ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(initialStackKB == 2, fmt.Sprintf("минимальный стартовый стек горутины — 2 KB, указано %d", initialStackKB))

	done := make(chan int)
	go func() { done <- depth(200000) }() // в новой горутине: стартуем с маленького стека
	got := <-done
	check(got == 200000, fmt.Sprintf("depth(200000) должен вернуть 200000, получено %d", got))
	check(depth(0) == 0, "depth(0) должен вернуть 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
