package main

import "fmt"

type Signed interface{ ~int | ~int64 }

func Neg[T Signed](v T) T { return -v }
func main()               { fmt.Println(Neg(5)) }
