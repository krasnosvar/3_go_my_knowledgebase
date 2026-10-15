// Упражнение 02: кратчайший путь BFS в невзвешенном графе.
//
// Задача: реализуйте bfsPath(adj, from, to) — кратчайший (по числу рёбер) путь
// от from до to включительно. Если пути нет — nil; если from == to — [from].
// Подсказка: BFS с массивом parent; путь восстанавливается от to к from
// по parent и разворачивается. Соседей обходите в порядке adj[v].
//
// Запуск: go run ./03_exercises/02_bfs_path/main.go
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
	// TODO: BFS + parent + восстановление пути.
	return nil
}

func main() {
	fmt.Println("--- Проверка 02_bfs_path ---")

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
