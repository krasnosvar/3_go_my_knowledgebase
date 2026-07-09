package main

import "fmt"

type U struct{ Name string }

func (u U) Valid() bool { return u.Name != "" }
func main()             { fmt.Println(U{Name: "Den"}.Valid()) }
