package main

import "fmt"

type Getter[T any] interface{ Get() T }
type S struct{}

func (S) Get() int { return 1 }
func main()        { var g Getter[int] = S{}; fmt.Println(g.Get()) }
