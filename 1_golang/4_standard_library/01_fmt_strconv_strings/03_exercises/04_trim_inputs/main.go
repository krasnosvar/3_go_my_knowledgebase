// Упражнение 04: нормализация пользовательского ввода.
//
// Задача: реализуйте normalizeInputs(in), которая для каждого элемента:
//   - обрезает пробелы по краям (strings.TrimSpace);
//   - приводит к нижнему регистру (strings.ToLower);
//   - пропускает элементы, ставшие пустыми.
//
// Исходный срез менять нельзя — верните новый.
//
// Запуск: go run ./03_exercises/04_trim_inputs/main.go
package main

import (
	"fmt"
	"strings"
)

var _ = strings.TrimSpace // держим импорт живым в заготовке

// normalizeInputs возвращает очищенные значения в нижнем регистре.
func normalizeInputs(in []string) []string {
	// TODO: TrimSpace + ToLower, пустые значения пропустить.
	return in
}

func main() {
	fmt.Println("--- Проверка 04_trim_inputs ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := []string{"  Alice ", "BOB", "   ", "\tCarol\n"}
	got := normalizeInputs(in)
	want := []string{"alice", "bob", "carol"}
	check(fmt.Sprint(got) == fmt.Sprint(want), fmt.Sprintf("ожидалось %q, получено %q", want, got))
	check(in[0] == "  Alice ", "исходный срез не должен меняться")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
