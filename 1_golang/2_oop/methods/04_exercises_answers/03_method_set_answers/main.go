package main

import "fmt"

type T struct{}

func (T) A()  { fmt.Println("A") }
func (*T) B() { fmt.Println("B") }
func main()   { t := &T{}; t.A(); t.B() }
