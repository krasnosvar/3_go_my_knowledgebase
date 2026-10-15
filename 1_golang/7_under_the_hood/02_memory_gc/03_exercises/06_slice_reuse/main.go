// Упражнение 06: переиспользование слайса.
//
// Задача: реализуйте filterInPlace(s, keep): оставить элементы, для которых
// keep(v) == true, переиспользуя массив s (приём out := s[:0]). Новый слайс
// создавать нельзя — результат должен разделять backing array с s.
//
// Проверка: верный результат, тот же backing array и 0 аллокаций.
//
// Запуск: go run ./03_exercises/06_slice_reuse/main.go
package main

import (
	"fmt"
	"slices"
	"testing"
)

// filterInPlace оставляет в s только элементы, для которых keep(v) == true.
func filterInPlace(s []int, keep func(int) bool) []int {
	// TODO: out := s[:0]; дописывайте подходящие элементы в out.
	var out []int
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func isEven(v int) bool { return v%2 == 0 }

func main() {
	fmt.Println("--- Проверка 06_slice_reuse ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	s := []int{1, 2, 3, 4, 5, 6}
	got := filterInPlace(s, isEven)
	check(slices.Equal(got, []int{2, 4, 6}), fmt.Sprintf("ожидалось [2 4 6], получено %v", got))
	check(len(got) > 0 && &got[0] == &s[0], "результат должен использовать тот же backing array, что и s")

	data := make([]int, 100)
	allocs := testing.AllocsPerRun(100, func() {
		for i := range data {
			data[i] = i
		}
		_ = filterInPlace(data, isEven)
	})
	check(allocs == 0, fmt.Sprintf("ожидалось 0 аллокаций, получено %.0f", allocs))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
