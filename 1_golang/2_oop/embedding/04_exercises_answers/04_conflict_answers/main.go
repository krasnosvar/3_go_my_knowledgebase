package main

import "fmt"

type A struct{}

type B struct{}

func (A) Name() { fmt.Println("A") }
func (B) Name() { fmt.Println("B") }

type C struct {
	A
	B
}

func main() {
	c := C{}
	c.A.Name()
	c.B.Name()
}
