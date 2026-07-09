package main

import "fmt"

func First[T any](items []T) T { return items[0] }
func main()                    { fmt.Println(First([]string{"go", "kb"})) }
