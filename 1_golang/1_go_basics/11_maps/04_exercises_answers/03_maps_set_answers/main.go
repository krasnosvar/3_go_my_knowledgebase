package main

import "fmt"

/*
	ОТВЕТЫ К ЗАДАНИЮ 3
*/

// ================= РЕШЕНИЕ: Set =================
type Set struct {
	items map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[string]struct{}),
	}
}

func (s *Set) Add(item string) {
	s.items[item] = struct{}{}
}

func (s *Set) Remove(item string) {
	delete(s.items, item)
}

func (s *Set) Contains(item string) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Intersection(other *Set) *Set {
	result := NewSet()
	for item := range s.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// ===============================================

func main() {
	s1 := NewSet()
	s1.Add("apple")
	s1.Add("banana")

	s2 := NewSet()
	s2.Add("banana")
	s2.Add("orange")

	intersect := s1.Intersection(s2)
	fmt.Println("Contains banana:", intersect.Contains("banana")) // true
	fmt.Println("Contains apple:", intersect.Contains("apple"))   // false
}
