// Упражнение 03: проверка сбалансированности скобок.
//
// Задача: реализуйте isBalanced(s) — true, если скобки (), [], {} в строке
// сбалансированы и правильно вложены. Прочие символы игнорируются.
// Алгоритм: открывающую скобку кладём в стек; на закрывающей снимаем вершину
// и проверяем пару; в конце стек должен быть пуст.
//
// Запуск: go run ./03_exercises/03_balanced_brackets/main.go
package main

import "fmt"

// isBalanced проверяет корректность расстановки скобок.
func isBalanced(s string) bool {
	// TODO: стек открывающих скобок + таблица пар.
	return false
}

func main() {
	fmt.Println("--- Проверка 03_balanced_brackets ---")

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
