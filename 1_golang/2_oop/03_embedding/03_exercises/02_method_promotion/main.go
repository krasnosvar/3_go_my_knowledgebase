// Упражнение 02: продвижение методов и полей.
//
// Задача: встройте Logger в Service и реализуйте конструктор NewService,
// который задаёт Logger с Prefix "[<name>] ". Метод Format и поле Prefix
// должны «продвинуться» (promotion) и вызываться прямо у Service.
//
// Запуск: go run ./03_exercises/02_method_promotion/main.go
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

// Service — сервис с логгером.
type Service struct {
	Name string
	// TODO: встройте Logger.
}

// NewService создаёт сервис с префиксом логов "[name] ".
func NewService(name string) Service {
	// TODO: верните Service{Name: name, Logger: Logger{Prefix: ...}}.
	return Service{Name: name}
}

func main() {
	fmt.Println("--- Проверка 02_method_promotion ---")

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
