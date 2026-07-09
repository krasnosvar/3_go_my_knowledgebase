package main

import "fmt"

type Parent struct{}

func (Parent) Print() { fmt.Println("parent") }

type Child struct{ Parent }

func (Child) Print() { fmt.Println("child") }

func main() { Child{}.Print() }
