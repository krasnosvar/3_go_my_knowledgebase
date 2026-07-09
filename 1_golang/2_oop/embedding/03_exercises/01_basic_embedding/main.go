package main

import "fmt"

type A struct{}

func (A) Hello() { fmt.Println("hello") }

type B struct {
	// TODO: embedded A
}

func main() {
	b := B{}
	b.Hello()
}
