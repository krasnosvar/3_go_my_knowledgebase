// Упражнение 04: стек с минимумом за O(1) (Решение).
package main

import "fmt"

// MinStack — стек с поддержкой минимума за O(1).
type MinStack struct {
	items []int
	mins  []int
}

func (s *MinStack) Push(v int) {
	s.items = append(s.items, v)
	if len(s.mins) == 0 || v < s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, v)
	} else {
		s.mins = append(s.mins, s.mins[len(s.mins)-1])
	}
}

func (s *MinStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.mins = s.mins[:len(s.mins)-1]
	return v, true
}

func (s *MinStack) Min() (int, bool) {
	if len(s.mins) == 0 {
		return 0, false
	}
	return s.mins[len(s.mins)-1], true
}

func main() {
	fmt.Println("--- Проверка 04_min_stack (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var s MinStack
	_, ok := s.Min()
	check(!ok, "Min на пустом стеке должен вернуть ok=false")

	s.Push(5)
	s.Push(3)
	s.Push(7)
	s.Push(2)
	m, ok := s.Min()
	check(ok && m == 2, fmt.Sprintf("после 5,3,7,2: ожидался Min=2, получено (%d, %t)", m, ok))

	s.Pop() // убрали 2
	m, ok = s.Min()
	check(ok && m == 3, fmt.Sprintf("после Pop: ожидался Min=3, получено (%d, %t)", m, ok))

	s.Pop() // убрали 7
	s.Pop() // убрали 3
	m, ok = s.Min()
	check(ok && m == 5, fmt.Sprintf("остался 5: ожидался Min=5, получено (%d, %t)", m, ok))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
