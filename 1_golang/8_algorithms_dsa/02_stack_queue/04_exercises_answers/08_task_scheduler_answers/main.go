// Упражнение 08: планировщик round-robin на очереди (Решение).
package main

import "fmt"

// Task — задача с оставшимся объёмом работы.
type Task struct {
	Name string
	Work int
}

// roundRobin возвращает порядок завершения задач.
func roundRobin(tasks []Task, quantum int) []string {
	if quantum <= 0 {
		return nil
	}
	queue := append([]Task(nil), tasks...) // не портим входной срез
	var done []string
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		t.Work -= quantum
		if t.Work > 0 {
			queue = append(queue, t)
		} else {
			done = append(done, t.Name)
		}
	}
	return done
}

func main() {
	fmt.Println("--- Проверка 08_task_scheduler (Решение) ---")

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
