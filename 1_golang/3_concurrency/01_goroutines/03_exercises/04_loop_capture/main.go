// Упражнение 04: переменная цикла в горутине.
//
// Задача: collectIDs(ids) запускает по горутине на каждый элемент ids,
// каждая горутина отправляет «свой» id в канал результатов; функция
// собирает все значения и возвращает их отсортированными.
//
// Классическая ловушка (до Go 1.22): замыкание захватывало ОДНУ переменную
// цикла на все итерации, и горутины видели последнее значение. С Go 1.22
// у каждой итерации своя переменная, но идиоматичный и переносимый способ —
// явно передать значение аргументом: go func(id int) { ... }(id).
//
// Запуск: go run ./03_exercises/04_loop_capture/main.go
package main

import "fmt"

// collectIDs возвращает отсортированные id, отправленные из горутин.
func collectIDs(ids []int) []int {
	// TODO: на каждый id запустите горутину, передав id аргументом,
	// соберите ровно len(ids) значений из канала и отсортируйте.
	return nil
}

func main() {
	fmt.Println("--- Проверка 04_loop_capture ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	got := collectIDs([]int{5, 3, 1, 4, 2})
	want := []int{1, 2, 3, 4, 5}
	check(fmt.Sprint(got) == fmt.Sprint(want), fmt.Sprintf("collectIDs = %v, ожидалось %v", got, want))

	got = collectIDs(nil)
	check(len(got) == 0, "для пустого входа результат должен быть пустым")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
