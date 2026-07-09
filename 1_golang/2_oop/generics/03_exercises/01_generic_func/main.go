package main

import "fmt"

func Identity[T any](v T) T { /* TODO */ return v }
func main()                 { fmt.Println(Identity(7)) }
