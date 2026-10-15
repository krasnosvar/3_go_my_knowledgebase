// Упражнение 04: поиск пути к конфигу.
//
// Задача: реализуйте configPath(flagPath string, getenv func(string) string, home string) string:
//   - явный флаг (непустой) → он;
//   - иначе APP_CONFIG из окружения → он;
//   - иначе XDG_CONFIG_HOME задан → <XDG_CONFIG_HOME>/app/config.yaml;
//   - иначе <home>/.config/app/config.yaml.
//
// Пути собирайте только через filepath.Join (не конкатенацией строк).
//
// Запуск: go run ./03_exercises/04_config_path
package main

import (
	"fmt"
	"path/filepath"
)

func configPath(flagPath string, getenv func(string) string, home string) string {
	// TODO: проверьте источники по приоритету.
	return ""
}

func main() {
	fmt.Println("--- Проверка 04_config_path ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	env := func(m map[string]string) func(string) string { return func(k string) string { return m[k] } }
	home := filepath.Join("home", "den")

	check(configPath("/etc/app.yaml", env(map[string]string{"APP_CONFIG": "x"}), home) == "/etc/app.yaml", "флаг важнее всего")
	check(configPath("", env(map[string]string{"APP_CONFIG": "/srv/c.yaml"}), home) == "/srv/c.yaml", "APP_CONFIG важнее XDG")
	check(configPath("", env(map[string]string{"XDG_CONFIG_HOME": "xdg"}), home) == filepath.Join("xdg", "app", "config.yaml"), "XDG_CONFIG_HOME/app/config.yaml")
	check(configPath("", env(nil), home) == filepath.Join(home, ".config", "app", "config.yaml"), "fallback в ~/.config")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
