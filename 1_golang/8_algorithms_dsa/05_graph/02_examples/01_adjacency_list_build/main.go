package main

import "fmt"

func main() {
	g := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4},
	}
	fmt.Println(g[1])
}
