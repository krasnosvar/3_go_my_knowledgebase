// Упражнение 04: приоритет источников конфигурации.
//
// Задача: реализуйте resolve(key string, flags, env, file map[string]string, def string) string:
// приоритет flags > env > file > def. Ключ «задан», если он присутствует
// в map (даже с пустым значением — значит его явно обнулили).
//
// Запуск: go run ./03_exercises/04_priority_merge
package main

import "fmt"

func resolve(key string, flags, env, file map[string]string, def string) string {
	// TODO: проверьте источники по убыванию приоритета через comma-ok.
	return ""
}

func main() {
	fmt.Println("--- Проверка 04_priority_merge ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	flags := map[string]string{"addr": ":1"}
	env := map[string]string{"addr": ":2", "level": "debug", "tag": ""}
	file := map[string]string{"addr": ":3", "level": "info", "region": "eu", "tag": "x"}

	check(resolve("addr", flags, env, file, ":0") == ":1", "flags побеждают всех")
	check(resolve("level", flags, env, file, "warn") == "debug", "env побеждает file")
	check(resolve("region", flags, env, file, "us") == "eu", "file побеждает default")
	check(resolve("missing", flags, env, file, "def") == "def", "нигде не задан → default")
	check(resolve("tag", flags, env, file, "def") == "", "пустое значение в env — явно задано")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
