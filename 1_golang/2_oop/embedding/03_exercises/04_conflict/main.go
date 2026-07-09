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
	// TODO: вызови оба метода явно через поля A и B.
	c := C{}
	_ = c
}
