// Упражнение 05: цикл в ориентированном графе (три цвета) (Решение).
package main

import "fmt"

// hasCycleDirected сообщает, есть ли цикл в ориентированном графе.
func hasCycleDirected(adj [][]int) bool {
	color := make([]int, len(adj))
	var dfs func(v int) bool
	dfs = func(v int) bool {
		color[v] = 1
		for _, n := range adj[v] {
			if color[n] == 1 {
				return true
			}
			if color[n] == 0 && dfs(n) {
				return true
			}
		}
		color[v] = 2
		return false
	}
	for v := range adj {
		if color[v] == 0 && dfs(v) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("--- Проверка 05_cycle_detect (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	dag := [][]int{0: {1, 2}, 1: {3}, 2: {3}, 3: {}}
	check(!hasCycleDirected(dag), "DAG с ромбом 0->1->3, 0->2->3: цикла нет")

	cyclic := [][]int{0: {1}, 1: {2}, 2: {0}, 3: {}}
	check(hasCycleDirected(cyclic), "0->1->2->0: ожидался цикл")

	disconnected := [][]int{0: {}, 1: {2}, 2: {1}}
	check(hasCycleDirected(disconnected), "цикл 1<->2 в несвязной компоненте должен находиться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
