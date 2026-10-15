// Упражнение 02: кратчайший путь BFS в невзвешенном графе (Решение).
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

// bfsPath возвращает кратчайший путь от from до to или nil.
func bfsPath(adj [][]int, from, to int) []int {
	parent := make([]int, len(adj))
	for i := range parent {
		parent[i] = -1
	}
	visited := make([]bool, len(adj))
	visited[from] = true
	queue := []int{from}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == to {
			var path []int
			for cur := to; cur != -1; cur = parent[cur] {
				path = append(path, cur)
			}
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path
		}
		for _, n := range adj[v] {
			if !visited[n] {
				visited[n] = true
				parent[n] = v
				queue = append(queue, n)
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_bfs_path (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := bfsPath(sample, 0, 3)
	check(fmt.Sprint(got) == "[0 1 3]", fmt.Sprintf("0->3: ожидалось [0 1 3], получено %v", got))

	got = bfsPath(sample, 2, 3)
	check(fmt.Sprint(got) == "[2 0 1 3]", fmt.Sprintf("2->3: ожидалось [2 0 1 3], получено %v", got))

	got = bfsPath(sample, 4, 4)
	check(fmt.Sprint(got) == "[4]", fmt.Sprintf("4->4: ожидалось [4], получено %v", got))

	check(bfsPath(sample, 0, 5) == nil, "0->5: вершина 5 недостижима, ожидался nil")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
