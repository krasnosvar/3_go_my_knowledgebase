// Упражнение 01: удаление дубликатов с сохранением порядка (Решение).
package main

import "fmt"

// dedupe убирает повторы, сохраняя порядок первых вхождений.
func dedupe(xs []string) []string {
	seen := make(map[string]struct{}, len(xs))
	out := make([]string, 0, len(xs))
	for _, x := range xs {
		if _, ok := seen[x]; ok {
			continue
		}
		seen[x] = struct{}{}
		out = append(out, x)
	}
	return out
}

func main() {
	fmt.Println("--- Проверка 01_deduplicate (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := dedupe([]string{"a", "b", "a", "c", "b"})
	check(fmt.Sprint(got) == "[a b c]", fmt.Sprintf("ожидалось [a b c], получено %v", got))

	got = dedupe([]string{"x", "x", "x"})
	check(fmt.Sprint(got) == "[x]", fmt.Sprintf("ожидалось [x], получено %v", got))

	got = dedupe(nil)
	check(len(got) == 0, fmt.Sprintf("пустой вход: ожидался пустой результат, получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
