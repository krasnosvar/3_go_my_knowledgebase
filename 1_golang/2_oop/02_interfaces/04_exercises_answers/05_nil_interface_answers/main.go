// Упражнение 05: ловушка typed nil в интерфейсе (Решение).
package main

import (
	"fmt"
)

// NotFoundError — ошибка «не найдено».
type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string { return fmt.Sprintf("id %d not found", e.ID) }

func find(id int) error {
	if id < 0 {
		return &NotFoundError{ID: id}
	}
	return nil
}

func main() {
	fmt.Println("--- Проверка 05_nil_interface (Решение) ---")

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
