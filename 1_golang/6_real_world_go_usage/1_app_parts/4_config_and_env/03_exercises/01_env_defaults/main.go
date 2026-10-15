// Упражнение 01: чтение конфига из окружения с дефолтами.
//
// Задача: реализуйте load(getenv func(string) string) config:
//   - APP_ADDR → Addr, по умолчанию ":8080";
//   - APP_ENV  → Env,  по умолчанию "dev";
//   - пустая строка считается «не задано».
//
// getenv передаётся параметром (в проде это os.Getenv) — так функция
// тестируется без изменения реального окружения процесса.
//
// Запуск: go run ./03_exercises/01_env_defaults
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
	// TODO: прочитайте APP_ADDR и APP_ENV, подставьте значения по умолчанию.
	return config{}
}

func main() {
	fmt.Println("--- Проверка 01_env_defaults ---")

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
