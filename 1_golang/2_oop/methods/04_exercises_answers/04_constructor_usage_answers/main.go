package main

import "fmt"

type U struct{ Name string }

func NewU(n string) *U { return &U{Name: n} }
func (u U) Hello()     { fmt.Println(u.Name) }
func main()            { NewU("Den").Hello() }
