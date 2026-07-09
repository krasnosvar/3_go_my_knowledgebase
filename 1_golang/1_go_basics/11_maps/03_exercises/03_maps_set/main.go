package main

import (
	"fmt"
)

// Set — структура множества строк на базе map[string]struct{}
type Set struct {
	data map[string]struct{}
}

func NewSet() *Set {
	return &Set{data: make(map[string]struct{})}
}

// Add добавляет элемент во множество
func (s *Set) Add(item string) {
	// TODO: Добавьте элемент
}

// Contains проверяет наличие элемента во множестве
func (s *Set) Contains(item string) bool {
	// TODO: Проверьте наличие элемента
	return false
}

func main() {
	fmt.Println("--- Проверка 03_maps_set ---")

	s := NewSet()
	s.Add("go")
	s.Add("rust")
	s.Add("go") // Повторное добавление

	t1 := s.Contains("go")
	t2 := s.Contains("rust")
	t3 := !s.Contains("c++")

	if t1 && t2 && t3 {
		fmt.Println("Тест Set: PASS")
	} else {
		fmt.Printf("Тест Set: FAIL (contains 'go': %t, contains 'rust': %t)\n", t1, t2)
	}
}
