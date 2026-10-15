// Упражнение 08: проверка двудольности (раскраска в 2 цвета) (Решение).
package main

import "fmt"

// isBipartite сообщает, двудолен ли граф.
func isBipartite(adj [][]int) bool {
	color := make([]int, len(adj))
	for i := range color {
		color[i] = -1
	}
	for start := range adj {
		if color[start] != -1 {
			continue
		}
		color[start] = 0
		queue := []int{start}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			for _, n := range adj[v] {
				if color[n] == -1 {
					color[n] = 1 - color[v]
					queue = append(queue, n)
				} else if color[n] == color[v] {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println("--- Проверка 08_is_bipartite (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	square := [][]int{0: {1, 3}, 1: {0, 2}, 2: {1, 3}, 3: {0, 2}}
	check(isBipartite(square), "цикл из 4 вершин двудолен")

	triangle := [][]int{0: {1, 2}, 1: {0, 2}, 2: {0, 1}}
	check(!isBipartite(triangle), "треугольник (нечётный цикл) не двудолен")

	mixed := [][]int{0: {1}, 1: {0}, 2: {3, 4}, 3: {2, 4}, 4: {2, 3}}
	check(!isBipartite(mixed), "несвязный граф с треугольником во второй компоненте не двудолен")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
