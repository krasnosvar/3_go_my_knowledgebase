// Упражнение 01: чтение конфига из окружения с дефолтами (Решение).
package main

import (
	"fmt"
	"os"
)

type config struct {
	Addr string
	Env  string
}

func load(getenv func(string) string) config {
	cfg := config{Addr: ":8080", Env: "dev"}
	if v := getenv("APP_ADDR"); v != "" {
		cfg.Addr = v
	}
	if v := getenv("APP_ENV"); v != "" {
		cfg.Env = v
	}
	return cfg
}

func main() {
	fmt.Println("--- Проверка 01_env_defaults (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	fake := func(m map[string]string) func(string) string {
		return func(k string) string { return m[k] }
	}

	c := load(fake(nil))
	check(c.Addr == ":8080" && c.Env == "dev", "пустое окружение → дефолты")
	c = load(fake(map[string]string{"APP_ADDR": ":9000", "APP_ENV": "prod"}))
	check(c.Addr == ":9000" && c.Env == "prod", "значения из окружения перекрывают дефолты")
	c = load(fake(map[string]string{"APP_ENV": ""}))
	check(c.Env == "dev", "пустая строка = не задано")
	_ = load(os.Getenv) // в проде передаётся os.Getenv

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
