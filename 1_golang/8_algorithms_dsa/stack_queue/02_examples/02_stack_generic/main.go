package main

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Push(v T) { *s = append(*s, v) }
func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var z T
		return z, false
	}
	i := len(*s) - 1
	v := (*s)[i]
	*s = (*s)[:i]
	return v, true
}

func main() {
	st := Stack[string]{}
	st.Push("a")
	v, _ := st.Pop()
	fmt.Println(v)
}
