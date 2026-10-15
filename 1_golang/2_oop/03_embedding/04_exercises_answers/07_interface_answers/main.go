// Упражнение 07: реализация интерфейса через встраивание (Решение).
package main

import (
	"fmt"
	"strings"
)

// Runner — что-то запускаемое.
type Runner interface {
	Run() string
}

// Base — общая реализация Run.
type Base struct{}

func (Base) Run() string { return "base run" }

type Job struct {
	Name string
	Base
}

func runAll(rs []Runner) string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, r.Run())
	}
	return strings.Join(out, ";")
}

func main() {
	fmt.Println("--- Проверка 07_interface (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var v any = Job{Name: "backup"}
	r, ok := v.(Runner)
	check(ok, "Job должен реализовывать Runner (встройте Base)")
	if ok {
		got := runAll([]Runner{r, Base{}})
		check(got == "base run;base run", fmt.Sprintf("runAll: получено %q", got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
