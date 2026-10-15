// Упражнение 06: generic-множество Set[T comparable] (Решение).
package main

import "fmt"

// Set — generic-множество на map[T]struct{}.
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet создаёт пустое множество.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func main() {
	fmt.Println("--- Проверка 06_generic_set (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ints := NewSet[int]()
	ints.Add(1)
	ints.Add(2)
	ints.Add(2)
	check(ints.Len() == 2, fmt.Sprintf("Set[int]: ожидался Len()=2, получено %d", ints.Len()))
	check(ints.Has(1) && !ints.Has(3), "Set[int]: Has(1)=true, Has(3)=false")

	ints.Remove(1)
	check(!ints.Has(1) && ints.Len() == 1, "после Remove(1) элемента 1 быть не должно")

	words := NewSet[string]()
	words.Add("go")
	check(words.Has("go") && words.Len() == 1, "Set[string]: ожидался один элемент \"go\"")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
