// Упражнение 07: расстояния от источника (BFS).
//
// Задача: реализуйте distances(adj, src) — для каждой вершины число рёбер
// кратчайшего пути от src; недостижимые вершины — -1.
// В невзвешенном графе BFS находит кратчайшие расстояния: вершина получает
// расстояние при ПЕРВОМ попадании в очередь (dist[n] = dist[v] + 1).
//
// Запуск: go run ./03_exercises/07_shortest_path_unweighted/main.go
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
	// TODO: BFS с массивом dist, заполненным -1.
	return nil
}

func main() {
	fmt.Println("--- Проверка 07_shortest_path_unweighted ---")

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
