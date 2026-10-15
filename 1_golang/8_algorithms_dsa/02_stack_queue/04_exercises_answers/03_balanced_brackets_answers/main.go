// Упражнение 03: проверка сбалансированности скобок (Решение).
package main

import "fmt"

// isBalanced проверяет корректность расстановки скобок.
func isBalanced(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println("--- Проверка 03_balanced_brackets (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"()[]{}", true},
		{"{[()()]}", true},
		{"f(a[1]) + {b}", true},
		{"([)]", false},
		{"(", false},
		{")(", false},
	}
	for _, c := range cases {
		got := isBalanced(c.in)
		check(got == c.want, fmt.Sprintf("isBalanced(%q): ожидалось %t, получено %t", c.in, c.want, got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
