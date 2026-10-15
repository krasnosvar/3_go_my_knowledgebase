// Упражнение 03: валидация конфига со сбором ВСЕХ ошибок.
//
// Задача: реализуйте (c config) Validate() error:
//   - DSN не пустой → иначе ошибка "dsn is required";
//   - Port в диапазоне 1..65535 → иначе "port out of range";
//   - Workers > 0 → иначе "workers must be > 0";
//   - вернуть все нарушения сразу через errors.Join (nil, если всё ок),
//     чтобы при старте сервиса оператор увидел полный список проблем.
//
// Запуск: go run ./03_exercises/03_validate_all
package main

import (
	"errors"
	"fmt"
	"strings"
)

type config struct {
	DSN     string
	Port    int
	Workers int
}

var _ = errors.Join

func (c config) Validate() error {
	// TODO: соберите нарушения в срез и верните errors.Join(errs...).
	return nil
}

func main() {
	fmt.Println("--- Проверка 03_validate_all ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	check(config{DSN: "postgres://x", Port: 5432, Workers: 4}.Validate() == nil, "валидный конфиг → nil")
	err := config{Port: 70000}.Validate()
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	check(strings.Contains(msg, "dsn is required"), "в ошибке есть про dsn")
	check(strings.Contains(msg, "port out of range"), "в ошибке есть про port")
	check(strings.Contains(msg, "workers must be > 0"), "в ошибке есть про workers")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
