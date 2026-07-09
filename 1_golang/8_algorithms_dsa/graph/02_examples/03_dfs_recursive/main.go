package main

import "fmt"

func dfs(g map[int][]int, v int, seen map[int]bool) {
	seen[v] = true
	fmt.Print(v, " ")
	for _, to := range g[v] {
		if !seen[to] {
			dfs(g, to, seen)
		}
	}
}

func main() {
	g := map[int][]int{1: {2, 3}, 2: {4}, 3: {}, 4: {}}
	seen := map[int]bool{}
	dfs(g, 1, seen)
	fmt.Println()
}
