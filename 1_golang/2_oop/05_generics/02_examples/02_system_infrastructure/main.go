package main

import "fmt"

type Number interface{ ~int | ~int64 | ~float64 }

func Sum[T Number](vals []T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}
func main() { fmt.Println(Sum([]int{1, 2, 3})) }
