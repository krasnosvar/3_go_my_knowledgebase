// Упражнение 01: построение списка смежности (Решение).
package main

import "fmt"

// buildUndirected строит список смежности неориентированного графа.
func buildUndirected(n int, edges [][2]int) [][]int {
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	return adj
}

func main() {
	fmt.Println("--- Проверка 01_build_graph (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	adj := buildUndirected(5, [][2]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}})
	check(fmt.Sprint(adj) == "[[1 2] [0 2] [0 1 3] [2] []]",
		fmt.Sprintf("ожидалось [[1 2] [0 2] [0 1 3] [2] []], получено %v", adj))
	check(len(adj) == 5 && adj[4] != nil, "изолированная вершина 4 должна иметь пустой, но не nil список")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
