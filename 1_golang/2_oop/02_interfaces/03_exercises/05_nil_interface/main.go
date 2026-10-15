// Упражнение 05: ловушка typed nil в интерфейсе.
//
// Задача: find возвращает error. В заготовке баг: при успехе она возвращает
// nil-указатель *NotFoundError, обёрнутый в интерфейс error. Такой интерфейс
// НЕ равен nil (в нём есть тип), и вызывающий код считает это ошибкой.
// Исправьте find: при успехе возвращайте литерал nil, при id < 0 —
// &NotFoundError{ID: id}.
//
// Запуск: go run ./03_exercises/05_nil_interface/main.go
package main

import (
	"fmt"
)

// NotFoundError — ошибка «не найдено».
type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string { return fmt.Sprintf("id %d not found", e.ID) }

// find ищет запись по id.
func find(id int) error {
	// TODO: баг — typed nil. Верните nil явно, если ошибки нет.
	var e *NotFoundError
	if id < 0 {
		e = &NotFoundError{ID: id}
	}
	return e
}

func main() {
	fmt.Println("--- Проверка 05_nil_interface ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	err := find(1)
	check(err == nil, fmt.Sprintf("find(1): ожидался nil, получено %#v (typed nil в интерфейсе?)", err))

	err = find(-2)
	check(err != nil, "find(-2): ожидалась ошибка")
	if err != nil {
		check(err.Error() == "id -2 not found", fmt.Sprintf("текст ошибки: получено %q", err.Error()))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
