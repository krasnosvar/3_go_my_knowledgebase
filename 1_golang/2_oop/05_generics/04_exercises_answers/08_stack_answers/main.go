// Упражнение 08: обобщённый стек (Решение).
package main

import (
	"fmt"
)

// Stack — обобщённый стек.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func main() {
	fmt.Println("--- Проверка 08_stack (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s Stack[string]
	s.Push("a")
	s.Push("b")
	s.Push("c")
	check(s.Len() == 3, fmt.Sprintf("Len после 3 Push: ожидалось 3, получено %d", s.Len()))

	top, ok := s.Peek()
	check(ok && top == "c" && s.Len() == 3, fmt.Sprintf("Peek: ожидалось (\"c\", true) без снятия, получено (%q, %t), Len=%d", top, ok, s.Len()))

	order := ""
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		order += v
		if len(order) > 10 {
			break // защита от бесконечного цикла в заготовке
		}
	}
	check(order == "cba", fmt.Sprintf("порядок Pop (LIFO): ожидалось \"cba\", получено %q", order))

	_, ok = s.Pop()
	check(!ok, "Pop пустого стека: ожидалось ok=false")

	var ints Stack[int]
	ints.Push(7)
	n, _ := ints.Pop()
	check(n == 7, "Stack[int]: ожидалось 7")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
