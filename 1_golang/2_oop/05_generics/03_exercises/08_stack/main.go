// Упражнение 08: обобщённый стек.
//
// Задача: реализуйте методы Stack[T] (LIFO на слайсе):
//   - Push(v) — положить наверх;
//   - Pop() (T, bool) — снять верхний; для пустого (zero, false);
//   - Peek() (T, bool) — посмотреть верхний без снятия;
//   - Len() int — размер.
//
// Запуск: go run ./03_exercises/08_stack/main.go
package main

import (
	"fmt"
)

// Stack — обобщённый стек.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	// TODO: append
}

func (s *Stack[T]) Pop() (T, bool) {
	// TODO: проверить пустоту, взять последний, укоротить слайс.
	var zero T
	return zero, false
}

func (s *Stack[T]) Peek() (T, bool) {
	// TODO
	var zero T
	return zero, false
}

func (s *Stack[T]) Len() int {
	// TODO
	return 0
}

func main() {
	fmt.Println("--- Проверка 08_stack ---")

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
