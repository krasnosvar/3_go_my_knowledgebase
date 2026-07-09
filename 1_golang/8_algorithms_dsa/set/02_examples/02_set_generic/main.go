package main

import "fmt"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T)      { s[v] = struct{}{} }
func (s Set[T]) Has(v T) bool { _, ok := s[v]; return ok }

func main() {
	s := Set[int]{}
	s.Add(10)
	fmt.Println(s.Has(10))
}
