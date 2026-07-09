package main

import "fmt"

func Eq[T comparable](a, b T) bool { return a == b }
func main()                        { fmt.Println(Eq("a", "a")) }
