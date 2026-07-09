package main

import "fmt"

type Repo[T any] struct{ items []T }

func (r *Repo[T]) Add(v T) { r.items = append(r.items, v) }
func main()                { r := &Repo[string]{}; r.Add("Den"); fmt.Println(r.items) }
