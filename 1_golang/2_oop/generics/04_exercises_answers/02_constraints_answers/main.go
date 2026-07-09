package main

import "fmt"

type N interface{ ~int | ~int64 | ~float64 }

func Sum[T N](a, b T) T { return a + b }
func main()             { fmt.Println(Sum(2, 3)) }
