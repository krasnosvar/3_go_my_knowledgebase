package main

import "fmt"

type B struct{ N int }

func (b *B) Add(v int) *B { b.N += v; return b }
func main()               { fmt.Println((&B{}).Add(2).Add(3).N) }
