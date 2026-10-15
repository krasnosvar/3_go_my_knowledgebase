// Упражнение 03: поиск пути обходом в глубину (Решение).
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

// isValidPath проверяет, что path начинается в from, заканчивается в to,
// идёт по рёбрам графа и не повторяет вершины.
func isValidPath(adj [][]int, path []int, from, to int) bool {
	if len(path) == 0 || path[0] != from || path[len(path)-1] != to {
		return false
	}
	seen := map[int]bool{}
	for i, v := range path {
		if seen[v] {
			return false
		}
		seen[v] = true
		if i == 0 {
			continue
		}
		ok := false
		for _, n := range adj[path[i-1]] {
			if n == v {
				ok = true
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

// dfsPath возвращает какой-нибудь путь от from до to или nil.
func dfsPath(adj [][]int, from, to int) []int {
	visited := make([]bool, len(adj))
	var path []int
	var dfs func(v int) bool
	dfs = func(v int) bool {
		visited[v] = true
		path = append(path, v)
		if v == to {
			return true
		}
		for _, n := range adj[v] {
			if !visited[n] && dfs(n) {
				return true
			}
		}
		path = path[:len(path)-1] // откат
		return false
	}
	if dfs(from) {
		return path
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_dfs_path (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := dfsPath(sample, 2, 3)
	check(isValidPath(sample, got, 2, 3), fmt.Sprintf("2->3: путь %v некорректен", got))

	got = dfsPath(sample, 0, 0)
	check(fmt.Sprint(got) == "[0]", fmt.Sprintf("0->0: ожидалось [0], получено %v", got))

	check(dfsPath(sample, 3, 5) == nil, "3->5: вершина 5 недостижима, ожидался nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
