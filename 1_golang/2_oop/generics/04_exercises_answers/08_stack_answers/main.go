package main

import "fmt"

type Stack[T any] struct{ a []T }

func (s *Stack[T]) Push(v T) { s.a = append(s.a, v) }
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.a) == 0 {
		var z T
		return z, false
	}
	v := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return v, true
}
func main() { st := &Stack[int]{}; st.Push(7); v, _ := st.Pop(); fmt.Println(v) }
