package main

import "fmt"

func bfs(g map[int][]int, start int) []int {
	q := []int{start}
	seen := map[int]bool{start: true}
	order := make([]int, 0)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		order = append(order, v)
		for _, to := range g[v] {
			if !seen[to] {
				seen[to] = true
				q = append(q, to)
			}
		}
	}
	return order
}

func main() {
	g := map[int][]int{1: {2, 3}, 2: {4}, 3: {}, 4: {}}
	fmt.Println(bfs(g, 1))
}
