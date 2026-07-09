package main

import "fmt"

type Parent struct{}

func (Parent) Print() { fmt.Println("parent") }

type Child struct{ Parent }

// TODO: override Print() for Child.

func main() {
	Child{}.Print()
}
