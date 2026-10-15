// Упражнение 04: пересечение множеств (Решение).
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

// intersection возвращает пересечение a и b.
func intersection(a, b map[int]struct{}) map[int]struct{} {
	if len(a) > len(b) {
		a, b = b, a
	}
	out := make(map[int]struct{})
	for k := range a {
		if _, ok := b[k]; ok {
			out[k] = struct{}{}
		}
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 04_intersection (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := sortedKeys(intersection(setOf(1, 2, 3, 4), setOf(3, 4, 5)))
	check(fmt.Sprint(got) == "[3 4]", fmt.Sprintf("ожидалось [3 4], получено %v", got))

	got = sortedKeys(intersection(setOf(1, 2), setOf(3)))
	check(len(got) == 0, fmt.Sprintf("нет общих элементов: ожидалось [], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
