package main

import (
	"fmt"
)

type Set struct {
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{data: make(map[string]struct{})}
}

func (s *Set) Add(item string) {
	s.data[item] = struct{}{}
}

func (s *Set) Contains(item string) bool {
	_, ok := s.data[item]
	return ok
}

func main() {
	fmt.Println("--- Проверка 03_maps_set (Решение) ---")

	s := NewSet()
	s.Add("go")
	s.Add("rust")
	s.Add("go")

	t1 := s.Contains("go")
	t2 := s.Contains("rust")
	t3 := !s.Contains("c++")

	if t1 && t2 && t3 {
		fmt.Println("Тест Set: PASS")
	} else {
		fmt.Printf("Тест Set: FAIL (contains 'go': %t, contains 'rust': %t)\n", t1, t2)
	}
}
