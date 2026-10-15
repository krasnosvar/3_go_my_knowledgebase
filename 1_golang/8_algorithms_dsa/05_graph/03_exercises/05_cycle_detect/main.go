// Упражнение 05: цикл в ориентированном графе (три цвета).
//
// Задача: реализуйте hasCycleDirected(adj) — есть ли цикл в ОРИЕНТИРОВАННОМ
// графе. DFS с тремя цветами:
//
//	white (0) — не посещена; gray (1) — в текущем стеке рекурсии;
//	black (2) — полностью обработана.
//
// Ребро в gray-вершину = обратное ребро = цикл. Запускайте DFS из каждой
// белой вершины (граф может быть несвязным).
//
// Запуск: go run ./03_exercises/05_cycle_detect/main.go
package main

import "fmt"

// hasCycleDirected сообщает, есть ли цикл в ориентированном графе.
func hasCycleDirected(adj [][]int) bool {
	// TODO: DFS с цветами white/gray/black.
	return false
}

func main() {
	fmt.Println("--- Проверка 05_cycle_detect ---")

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
