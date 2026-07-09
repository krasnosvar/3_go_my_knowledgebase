package main

import "fmt"

type Base struct{ Name string }

type Child struct{ Base }

func NewChild(name string) Child {
	// TODO: return Child with initialized embedded Base.
	return Child{}
}

func main() {
	c := NewChild("Den")
	fmt.Println(c.Name)
}
