// Упражнение 06: минимальное окно, содержащее все символы (LeetCode 76) (Решение).
package main

import "fmt"

// minWindow возвращает минимальную подстроку s, покрывающую t.
func minWindow(s, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing := len(t)
	bestL, bestR := 0, -1
	left := 0
	for right := 0; right < len(s); right++ {
		if need[s[right]] > 0 {
			missing--
		}
		need[s[right]]--
		for missing == 0 {
			if bestR < 0 || right-left < bestR-bestL {
				bestL, bestR = left, right
			}
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++
			}
			left++
		}
	}
	return s[bestL : bestR+1]
}

func main() {
	fmt.Println("--- Проверка 06_min_window (Решение) ---")

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
