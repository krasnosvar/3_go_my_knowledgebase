package main

import "fmt"

type Base struct{ Name string }

type Child struct{ Base }

func NewChild(name string) Child { return Child{Base: Base{Name: name}} }

func main() {
	c := NewChild("Den")
	fmt.Println(c.Name)
}
