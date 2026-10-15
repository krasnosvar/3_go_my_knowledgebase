// Упражнение 01: стек на срезе (Push/Pop/Peek) (Решение).
package main

import "fmt"

// Stack — стек целых чисел на срезе.
type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) Len() int {
	return len(s.items)
}

func main() {
	fmt.Println("--- Проверка 01_push_pop (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s Stack
	_, ok := s.Pop()
	check(!ok, "Pop на пустом стеке должен вернуть ok=false")

	s.Push(1)
	s.Push(2)
	s.Push(3)
	check(s.Len() == 3, fmt.Sprintf("после трёх Push ожидался Len()=3, получено %d", s.Len()))

	v, ok := s.Peek()
	check(ok && v == 3, fmt.Sprintf("Peek: ожидалось (3, true), получено (%d, %t)", v, ok))

	var popped []int
	for s.Len() > 0 && len(popped) < 10 {
		v, _ := s.Pop()
		popped = append(popped, v)
	}
	check(fmt.Sprint(popped) == "[3 2 1]", fmt.Sprintf("порядок Pop: ожидалось [3 2 1], получено %v", popped))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
