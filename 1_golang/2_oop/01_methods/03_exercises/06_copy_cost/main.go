// Упражнение 06: стоимость копирования большого receiver.
//
// Задача: Big занимает 8 КБ ([1024]int). Value receiver копирует всю
// структуру при каждом вызове. Переведите оба метода на pointer receiver и
// реализуйте их:
//   - Fill(v) записывает v во все элементы data;
//   - Sum() возвращает сумму элементов.
//
// Тест проверяет, что значение Big{} НЕ реализует интерфейсы с этими методами
// (значит, получатели — указатели и копирования нет).
//
// Запуск: go run ./03_exercises/06_copy_cost/main.go
package main

import (
	"fmt"
)

// Big — «тяжёлая» структура.
type Big struct {
	data [1024]int
}

// Fill заполняет массив значением v.
// TODO: pointer receiver + цикл по b.data.
func (b Big) Fill(v int) {
}

// Sum возвращает сумму элементов.
// TODO: pointer receiver + суммирование.
func (b Big) Sum() int {
	return 0
}

func main() {
	fmt.Println("--- Проверка 06_copy_cost ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	b := &Big{}
	b.Fill(2)
	check(b.Sum() == 2048, fmt.Sprintf("после Fill(2): ожидалось Sum()=2048, получено %d", b.Sum()))

	var v any = Big{}
	_, hasFill := v.(interface{ Fill(int) })
	_, hasSum := v.(interface{ Sum() int })
	check(!hasFill, "Fill должен быть с pointer receiver (иначе копируется 8 КБ и изменения теряются)")
	check(!hasSum, "Sum должен быть с pointer receiver (иначе копируется 8 КБ на каждый вызов)")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
