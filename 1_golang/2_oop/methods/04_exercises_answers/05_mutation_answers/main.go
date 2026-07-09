package main

import "fmt"

type S struct{ V int }

func (s *S) Set(v int) { s.V = v }
func main()            { s := &S{}; s.Set(5); fmt.Println(s.V) }
