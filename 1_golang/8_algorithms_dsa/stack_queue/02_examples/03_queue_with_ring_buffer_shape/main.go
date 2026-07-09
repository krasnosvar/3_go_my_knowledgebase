package main

import "fmt"

func main() {
	q := make([]int, 0, 4)
	q = append(q, 1, 2, 3)
	fmt.Println(q[0])
	q = q[1:]
	fmt.Println(q)
}
