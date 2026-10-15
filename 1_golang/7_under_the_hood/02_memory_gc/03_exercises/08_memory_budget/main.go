// Упражнение 08: бюджет памяти.
//
// Задача:
//  1. Реализуйте int64SliceBytes(n): сколько байт занимает []int64 длины n
//     с len == cap: заголовок слайса (unsafe.Sizeof([]int64{})) + n элементов
//     по unsafe.Sizeof(int64(0)) байт.
//  2. Реализуйте maxItems(budget, perItem): сколько элементов по perItem байт
//     влезает в бюджет budget байт (perItem == 0 -> 0, чтобы не делить на ноль).
//
// Запуск: go run ./03_exercises/08_memory_budget/main.go
package main

import (
	"fmt"
	"unsafe"
)

// int64SliceBytes возвращает размер заголовка слайса плюс данных для n int64.
func int64SliceBytes(n int) uintptr {
	// TODO: заголовок слайса + n * размер int64.
	return 0
}

// maxItems возвращает, сколько элементов по perItem байт влезает в budget.
func maxItems(budget, perItem uint64) uint64 {
	// TODO: budget / perItem, с защитой от perItem == 0.
	return 0
}

func main() {
	fmt.Println("--- Проверка 08_memory_budget ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ptr := unsafe.Sizeof(uintptr(0))
	check(int64SliceBytes(0) == 3*ptr, fmt.Sprintf("пустой []int64: ожидалось %d байт (заголовок), получено %d", 3*ptr, int64SliceBytes(0)))
	check(int64SliceBytes(1000) == 3*ptr+8000, fmt.Sprintf("[]int64 из 1000: ожидалось %d, получено %d", 3*ptr+8000, int64SliceBytes(1000)))

	const mb = 1 << 20
	check(maxItems(64*mb, 1024) == 65536, fmt.Sprintf("64 MB / 1 KB: ожидалось 65536, получено %d", maxItems(64*mb, 1024)))
	check(maxItems(100, 0) == 0, "perItem == 0: ожидалось 0")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
