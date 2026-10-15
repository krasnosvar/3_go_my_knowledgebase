// Упражнение 03: валидация конфига со сбором всех ошибок (Решение).
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

func (c config) Validate() error {
	var errs []error
	if c.DSN == "" {
		errs = append(errs, errors.New("dsn is required"))
	}
	if c.Port < 1 || c.Port > 65535 {
		errs = append(errs, errors.New("port out of range"))
	}
	if c.Workers <= 0 {
		errs = append(errs, errors.New("workers must be > 0"))
	}
	return errors.Join(errs...)
}

func main() {
	fmt.Println("--- Проверка 03_validate_all (Решение) ---")

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
