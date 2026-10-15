// Упражнение 04: приоритет источников конфигурации (Решение).
package main

import "fmt"

func resolve(key string, flags, env, file map[string]string, def string) string {
	for _, src := range []map[string]string{flags, env, file} {
		if v, ok := src[key]; ok {
			return v
		}
	}
	return def
}

func main() {
	fmt.Println("--- Проверка 04_priority_merge (Решение) ---")

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
