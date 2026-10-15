// Упражнение 04: коды выхода.
//
// main() в CLI должен быть тонким: os.Exit(run(os.Args[1:], ...)). Логика —
// в run, которая возвращает код и поэтому тестируется.
//
// Задача: реализуйте run(args []string) int для утилиты "sum":
//   - все аргументы — целые числа; печатать ничего не нужно;
//   - нет аргументов → 2 (ошибка использования);
//   - аргумент не число → 2;
//   - сумма отрицательна → 1 (ошибка выполнения: «бизнес»-ошибка);
//   - иначе → 0.
//
// Запуск: go run ./03_exercises/04_exit_codes
package main

import (
	"fmt"
	"strconv"
)

const (
	exitOK    = 0
	exitError = 1
	exitUsage = 2
)

var _ = strconv.Atoi

func run(args []string) int {
	// TODO: верните exitUsage / exitError / exitOK по правилам.
	return -1
}

func main() {
	fmt.Println("--- Проверка 04_exit_codes ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(run([]string{"1", "2"}) == 0, "1 2 → 0")
	check(run(nil) == 2, "без аргументов → 2")
	check(run([]string{"1", "x"}) == 2, "не число → 2")
	check(run([]string{"1", "-5"}) == 1, "отрицательная сумма → 1")
	// В настоящей утилите: func main() { os.Exit(run(os.Args[1:])) }

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
