package main

import "fmt"

type A struct{}

func (A) Hello() { fmt.Println("hello") }

type B struct{ A }

func main() { B{}.Hello() }
