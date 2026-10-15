// Упражнение 05: самая длинная подстрока без повторов (LeetCode 3) (Решение).
package main

import "fmt"

// lengthOfLongestSubstring возвращает длину самой длинной подстроки без повторов.
func lengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	best, left := 0, 0
	for right := 0; right < len(s); right++ {
		if i, ok := last[s[right]]; ok && i >= left {
			left = i + 1
		}
		last[s[right]] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

func main() {
	fmt.Println("--- Проверка 05_longest_substring (Решение) ---")

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
