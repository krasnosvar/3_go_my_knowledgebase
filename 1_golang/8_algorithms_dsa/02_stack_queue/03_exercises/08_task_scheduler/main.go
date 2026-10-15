// Упражнение 08: планировщик round-robin на очереди.
//
// Задача: реализуйте roundRobin(tasks, quantum) — симуляцию планировщика.
// Задачи стоят в очереди; каждая по очереди получает quantum единиц работы.
// Если после этого работа осталась — задача уходит в конец очереди, иначе
// считается завершённой. Верните имена задач в порядке ЗАВЕРШЕНИЯ.
// Пример: A:3, B:1, C:2, quantum=1 -> [B C A].
// Если quantum <= 0 — верните nil (иначе цикл не завершится).
//
// Запуск: go run ./03_exercises/08_task_scheduler/main.go
package main

import "fmt"

// Task — задача с оставшимся объёмом работы.
type Task struct {
	Name string
	Work int
}

// roundRobin возвращает порядок завершения задач.
func roundRobin(tasks []Task, quantum int) []string {
	// TODO: очередь задач; уменьшайте Work на quantum.
	return nil
}

func main() {
	fmt.Println("--- Проверка 08_task_scheduler ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	tasks := []Task{{"A", 3}, {"B", 1}, {"C", 2}}
	got := roundRobin(tasks, 1)
	check(fmt.Sprint(got) == "[B C A]", fmt.Sprintf("quantum=1: ожидалось [B C A], получено %v", got))

	got = roundRobin(tasks, 2)
	check(fmt.Sprint(got) == "[B C A]", fmt.Sprintf("quantum=2: ожидалось [B C A], получено %v", got))

	got = roundRobin(tasks, 3)
	check(fmt.Sprint(got) == "[A B C]", fmt.Sprintf("quantum=3: ожидалось [A B C], получено %v", got))

	check(tasks[0].Work == 3, "входной срез tasks не должен изменяться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
