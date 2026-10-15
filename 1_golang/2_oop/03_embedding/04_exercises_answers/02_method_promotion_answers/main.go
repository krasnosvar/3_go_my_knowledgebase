// Упражнение 02: продвижение методов и полей (Решение).
package main

import (
	"fmt"
)

// Logger форматирует сообщения с префиксом.
type Logger struct {
	Prefix string
}

// Format добавляет префикс.
func (l Logger) Format(msg string) string { return l.Prefix + msg }

type Service struct {
	Name string
	Logger
}

func NewService(name string) Service {
	return Service{Name: name, Logger: Logger{Prefix: "[" + name + "] "}}
}

func main() {
	fmt.Println("--- Проверка 02_method_promotion (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var v any = NewService("api")
	f, ok := v.(interface{ Format(string) string })
	check(ok, "Service должен иметь продвинутый метод Format (встройте Logger)")
	if ok {
		got := f.Format("started")
		check(got == "[api] started", fmt.Sprintf("Format: ожидалось \"[api] started\", получено %q", got))
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
