// Упражнение 06: переиспользование слайса (Решение).
package main

import (
	"fmt"
	"slices"
	"testing"
)

// filterInPlace оставляет в s только элементы, для которых keep(v) == true.
func filterInPlace(s []int, keep func(int) bool) []int {
	out := s[:0] // длина 0, но тот же массив и ёмкость
	for _, v := range s {
		if keep(v) {
			out = append(out, v) // пишет поверх уже прочитанных элементов
		}
	}
	return out
}

func isEven(v int) bool { return v%2 == 0 }

func main() {
	fmt.Println("--- Проверка 06_slice_reuse (Решение) ---")

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
