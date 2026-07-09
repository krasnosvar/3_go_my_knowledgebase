package main

import "fmt"

func main() {
	a := map[int]struct{}{1: {}, 2: {}, 3: {}}
	b := map[int]struct{}{2: {}, 3: {}, 4: {}}
	res := make([]int, 0)
	for v := range a {
		if _, ok := b[v]; ok {
			res = append(res, v)
		}
	}
	fmt.Println(res)
}
