// Упражнение 01: удаление дубликатов с сохранением порядка.
//
// Задача: реализуйте dedupe(xs) — возвращает новый срез без повторов,
// сохраняя порядок ПЕРВЫХ вхождений.
// Пример: [a b a c b] -> [a b c].
// Подсказка: множество seen := map[string]struct{}{} — O(n) вместо O(n^2).
//
// Запуск: go run ./03_exercises/01_deduplicate/main.go
package main

import "fmt"

// dedupe убирает повторы, сохраняя порядок первых вхождений.
func dedupe(xs []string) []string {
	// TODO: множество уже встреченных строк.
	return xs
}

func main() {
	fmt.Println("--- Проверка 01_deduplicate ---")

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
