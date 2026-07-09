package main

import "fmt"

type Box[T any] struct{ V T }

func main() { fmt.Println(Box[int]{V: 5}.V) }
