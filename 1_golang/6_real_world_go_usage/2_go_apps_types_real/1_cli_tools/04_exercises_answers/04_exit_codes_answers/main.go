// Упражнение 04: коды выхода (Решение).
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

func run(args []string) int {
	if len(args) == 0 {
		return exitUsage
	}
	sum := 0
	for _, a := range args {
		n, err := strconv.Atoi(a)
		if err != nil {
			return exitUsage
		}
		sum += n
	}
	if sum < 0 {
		return exitError
	}
	return exitOK
}

func main() {
	fmt.Println("--- Проверка 04_exit_codes (Решение) ---")

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
