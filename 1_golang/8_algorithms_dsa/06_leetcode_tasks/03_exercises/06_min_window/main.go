// Упражнение 06: минимальное окно, содержащее все символы (LeetCode 76).
//
// Задача: реализуйте minWindow(s, t) — самая короткая подстрока s, содержащая
// все символы t с учётом кратности. Если такой нет — "". Ответ единственный.
// Скользящее окно со счётчиками: расширяем правую границу, пока окно не
// покроет t (missing == 0), затем сжимаем левую, обновляя лучший ответ. O(|s|+|t|).
//
// Запуск: go run ./03_exercises/06_min_window/main.go
package main

import "fmt"

// minWindow возвращает минимальную подстроку s, покрывающую t.
func minWindow(s, t string) string {
	// TODO: need[c] — сколько символа c ещё нужно; missing — сколько всего не хватает.
	return ""
}

func main() {
	fmt.Println("--- Проверка 06_min_window ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	for _, c := range []struct{ s, t, want string }{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"aaflslflsldkalskaaa", "aaa", "aaa"},
	} {
		got := minWindow(c.s, c.t)
		check(got == c.want, fmt.Sprintf("minWindow(%q, %q): ожидалось %q, получено %q", c.s, c.t, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
