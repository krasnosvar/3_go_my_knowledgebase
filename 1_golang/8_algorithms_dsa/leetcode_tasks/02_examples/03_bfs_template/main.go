package main

import "fmt"

func bfs(g map[int][]int, s int) []int {
	q := []int{s}
	seen := map[int]bool{s: true}
	res := []int{}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		res = append(res, v)
		for _, to := range g[v] {
			if !seen[to] {
				seen[to] = true
				q = append(q, to)
			}
		}
	}
	return res
}

func main() { fmt.Println(bfs(map[int][]int{1: {2, 3}, 2: {}, 3: {}}, 1)) }
