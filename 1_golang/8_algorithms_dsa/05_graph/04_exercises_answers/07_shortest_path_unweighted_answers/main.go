// Упражнение 07: расстояния от источника (BFS) (Решение).
package main

import "fmt"

// sample — неориентированный граф из 6 вершин списком смежности:
//
//	0 - 1 - 3
//	|   |
//	2 - 4    5 (изолирована)
var sample = [][]int{
	0: {1, 2},
	1: {0, 3, 4},
	2: {0, 4},
	3: {1},
	4: {1, 2},
	5: {},
}

// distances возвращает расстояния от src до всех вершин (-1 — недостижима).
func distances(adj [][]int, src int) []int {
	dist := make([]int, len(adj))
	for i := range dist {
		dist[i] = -1
	}
	dist[src] = 0
	queue := []int{src}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, n := range adj[v] {
			if dist[n] == -1 {
				dist[n] = dist[v] + 1
				queue = append(queue, n)
			}
		}
	}
	return dist
}

func main() {
	fmt.Println("--- Проверка 07_shortest_path_unweighted (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := distances(sample, 0)
	check(fmt.Sprint(got) == "[0 1 1 2 2 -1]", fmt.Sprintf("от 0: ожидалось [0 1 1 2 2 -1], получено %v", got))

	got = distances(sample, 3)
	check(fmt.Sprint(got) == "[2 1 3 0 2 -1]", fmt.Sprintf("от 3: ожидалось [2 1 3 0 2 -1], получено %v", got))

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
