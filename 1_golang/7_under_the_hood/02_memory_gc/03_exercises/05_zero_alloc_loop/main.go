// Упражнение 05: цикл без аллокаций.
//
// Задача: перепишите appendInts(dst, nums) так, чтобы при достаточной
// ёмкости dst она не аллоцировала: дописывайте числа через strconv.AppendInt
// прямо в dst, разделяя их запятой. Сейчас используется fmt.Sprintf,
// который на каждой итерации создаёт новую строку.
//
// Проверка: результат "1,22,-3" и 0 аллокаций при dst с запасом ёмкости.
//
// Запуск: go run ./03_exercises/05_zero_alloc_loop/main.go
package main

import (
	"fmt"
	"strconv"
	"testing"
)

// appendInts дописывает nums в dst через запятую и возвращает новый слайс.
func appendInts(dst []byte, nums []int) []byte {
	// TODO: замените Sprintf на strconv.AppendInt.
	_ = strconv.AppendInt
	for i, n := range nums {
		if i > 0 {
			dst = append(dst, ',')
		}
		dst = append(dst, fmt.Sprintf("%d", n)...)
	}
	return dst
}

func main() {
	fmt.Println("--- Проверка 05_zero_alloc_loop ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := string(appendInts(nil, []int{1, 22, -3}))
	check(got == "1,22,-3", fmt.Sprintf("ожидалось \"1,22,-3\", получено %q", got))
	check(string(appendInts([]byte("ids="), []int{7})) == "ids=7", "числа должны дописываться к существующему содержимому dst")

	nums := []int{100, 200, 300, 400, 500}
	buf := make([]byte, 0, 64)
	allocs := testing.AllocsPerRun(100, func() { buf = appendInts(buf[:0], nums) })
	check(allocs == 0, fmt.Sprintf("ожидалось 0 аллокаций, получено %.0f", allocs))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
