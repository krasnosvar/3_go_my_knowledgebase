package main

import "fmt"

type Big struct{ A [1024]int }

func (b *Big) Touch() { b.A[0] = 1 }
func main()           { x := &Big{}; x.Touch(); fmt.Println(x.A[0]) }
