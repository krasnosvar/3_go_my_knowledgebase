// Упражнение 05: самая длинная подстрока без повторов (LeetCode 3).
//
// Задача: реализуйте lengthOfLongestSubstring(s) — длина самой длинной
// подстроки без повторяющихся символов (байтов, вход — ASCII).
// Скользящее окно: правая граница идёт вперёд; если символ уже внутри окна,
// левая граница прыгает за его последнее вхождение. O(n).
//
// Запуск: go run ./03_exercises/05_longest_substring/main.go
package main

import "fmt"

// lengthOfLongestSubstring возвращает длину самой длинной подстроки без повторов.
func lengthOfLongestSubstring(s string) int {
	// TODO: окно [left, right] + последний индекс каждого символа.
	return 0
}

func main() {
	fmt.Println("--- Проверка 05_longest_substring ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	for _, c := range []struct {
		in   string
		want int
	}{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {"", 0}, {"abba", 2}, {"dvdf", 3}} {
		got := lengthOfLongestSubstring(c.in)
		check(got == c.want, fmt.Sprintf("lengthOfLongestSubstring(%q): ожидалось %d, получено %d", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
