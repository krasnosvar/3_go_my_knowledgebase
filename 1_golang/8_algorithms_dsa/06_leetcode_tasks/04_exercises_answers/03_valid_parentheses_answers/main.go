// Упражнение 03: валидные скобки (LeetCode 20) и минимум удалений (LeetCode 1249) (Решение).
package main

import (
	"fmt"
	"strings"
)

// isValid проверяет строку из скобок ()[]{}.
func isValid(s string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if open, ok := pairs[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

// minRemoveToMakeValid удаляет минимум круглых скобок для валидности.
func minRemoveToMakeValid(s string) string {
	remove := make(map[int]bool)
	var stack []int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				remove[i] = true
			}
		}
	}
	for _, i := range stack {
		remove[i] = true
	}
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if !remove[i] {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

func main() {
	fmt.Println("--- Проверка 03_valid_parentheses (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	for _, c := range []struct {
		in   string
		want bool
	}{{"()", true}, {"()[]{}", true}, {"{[]}", true}, {"(]", false}, {"([)]", false}, {"]", false}} {
		got := isValid(c.in)
		check(got == c.want, fmt.Sprintf("isValid(%q): ожидалось %t, получено %t", c.in, c.want, got))
	}

	for _, c := range []struct{ in, want string }{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
	} {
		got := minRemoveToMakeValid(c.in)
		check(got == c.want, fmt.Sprintf("minRemoveToMakeValid(%q): ожидалось %q, получено %q", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
