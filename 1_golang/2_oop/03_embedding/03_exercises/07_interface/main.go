// Упражнение 07: реализация интерфейса через встраивание.
//
// Задача: встройте Base в Job, чтобы Job реализовывал Runner за счёт
// продвинутого метода Run. Затем реализуйте runAll — вызывает Run у всех
// Runner и склеивает результаты через ";".
//
// Запуск: go run ./03_exercises/07_interface/main.go
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

// Job — задача.
type Job struct {
	Name string
	// TODO: встройте Base.
}

// runAll запускает все Runner и склеивает результаты через ";".
func runAll(rs []Runner) string {
	// TODO: соберите r.Run() и strings.Join(..., ";").
	_ = strings.Join
	return ""
}

func main() {
	fmt.Println("--- Проверка 07_interface ---")

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
