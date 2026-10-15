// Упражнение 03: объединение множеств (Решение).
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

// union возвращает объединение a и b.
func union(a, b map[int]struct{}) map[int]struct{} {
	out := make(map[int]struct{}, len(a)+len(b))
	for k := range a {
		out[k] = struct{}{}
	}
	for k := range b {
		out[k] = struct{}{}
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 03_union (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	a, b := setOf(1, 2, 3), setOf(3, 4)
	got := sortedKeys(union(a, b))
	check(fmt.Sprint(got) == "[1 2 3 4]", fmt.Sprintf("ожидалось [1 2 3 4], получено %v", got))
	check(len(a) == 3 && len(b) == 2, "входные множества не должны изменяться")

	got = sortedKeys(union(setOf(), setOf(7)))
	check(fmt.Sprint(got) == "[7]", fmt.Sprintf("объединение с пустым: ожидалось [7], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
