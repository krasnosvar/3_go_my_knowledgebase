// Упражнение 02: типизированный разбор переменных окружения (Решение).
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

func parse(env map[string]string) (config, error) {
	cfg := config{Timeout: 5 * time.Second}
	p, ok := env["PORT"]
	if !ok {
		return config{}, fmt.Errorf("PORT: required")
	}
	var err error
	if cfg.Port, err = strconv.Atoi(p); err != nil {
		return config{}, fmt.Errorf("PORT: %w", err)
	}
	if v, ok := env["TIMEOUT"]; ok {
		if cfg.Timeout, err = time.ParseDuration(v); err != nil {
			return config{}, fmt.Errorf("TIMEOUT: %w", err)
		}
	}
	if v, ok := env["DEBUG"]; ok {
		if cfg.Debug, err = strconv.ParseBool(v); err != nil {
			return config{}, fmt.Errorf("DEBUG: %w", err)
		}
	}
	return cfg, nil
}

func main() {
	fmt.Println("--- Проверка 02_typed_parse (Решение) ---")

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
