// Упражнение 02: типизированный разбор переменных окружения.
//
// Задача: реализуйте parse(env map[string]string) (config, error):
//   - PORT → int (strconv.Atoi), обязателен;
//   - TIMEOUT → time.Duration (time.ParseDuration), по умолчанию 5s;
//   - DEBUG → bool (strconv.ParseBool), по умолчанию false;
//   - ошибка разбора должна содержать имя переменной, например
//     "PORT: invalid syntax ..." (используйте fmt.Errorf("PORT: %w", err)).
//   - отсутствующий PORT → ошибка с текстом, содержащим "PORT".
//
// Запуск: go run ./03_exercises/02_typed_parse
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type config struct {
	Port    int
	Timeout time.Duration
	Debug   bool
}

var _ = strconv.Atoi

func parse(env map[string]string) (config, error) {
	// TODO: разберите PORT, TIMEOUT, DEBUG с дефолтами и понятными ошибками.
	return config{}, nil
}

func main() {
	fmt.Println("--- Проверка 02_typed_parse ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	c, err := parse(map[string]string{"PORT": "8080"})
	check(err == nil && c.Port == 8080 && c.Timeout == 5*time.Second && !c.Debug, "только PORT → дефолты для остального")
	c, err = parse(map[string]string{"PORT": "1", "TIMEOUT": "250ms", "DEBUG": "true"})
	check(err == nil && c.Timeout == 250*time.Millisecond && c.Debug, "TIMEOUT и DEBUG разбираются")
	_, err = parse(map[string]string{"PORT": "eighty"})
	check(err != nil && strings.Contains(err.Error(), "PORT"), "битый PORT → ошибка с именем переменной")
	_, err = parse(map[string]string{"PORT": "1", "TIMEOUT": "soon"})
	check(err != nil && strings.Contains(err.Error(), "TIMEOUT"), "битый TIMEOUT → ошибка с именем переменной")
	_, err = parse(map[string]string{})
	check(err != nil && strings.Contains(err.Error(), "PORT"), "отсутствующий PORT → ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
