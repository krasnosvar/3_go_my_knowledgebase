// Упражнение 05: разность множеств (Решение).
package main

import (
	"fmt"
	"sort"
)

// sortedKeys возвращает ключи множества по возрастанию (для детерминированной проверки).
func sortedKeys(s map[int]struct{}) []int {
	out := make([]int, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}

// setOf строит множество из перечисленных значений.
func setOf(vals ...int) map[int]struct{} {
	s := make(map[int]struct{}, len(vals))
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

// difference возвращает элементы a, отсутствующие в b.
func difference(a, b map[int]struct{}) map[int]struct{} {
	out := make(map[int]struct{})
	for k := range a {
		if _, ok := b[k]; !ok {
			out[k] = struct{}{}
		}
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 05_difference (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a, b := setOf(1, 2, 3, 4), setOf(2, 4, 6)
	got := sortedKeys(difference(a, b))
	check(fmt.Sprint(got) == "[1 3]", fmt.Sprintf("a \\ b: ожидалось [1 3], получено %v", got))

	got = sortedKeys(difference(b, a))
	check(fmt.Sprint(got) == "[6]", fmt.Sprintf("b \\ a: ожидалось [6], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
