// Упражнение 03: флаг с fallback на переменную окружения.
//
// Задача: реализуйте token(args []string, getenv func(string) string) (string, error):
//   - флаг -token (свой FlagSet, ContinueOnError, вывод в io.Discard);
//   - если флаг явно передан (даже пустым) — используется он;
//     проверить «передан ли» можно через fs.Visit;
//   - иначе — getenv("TOOL_TOKEN");
//   - если в итоге пусто → ошибка "token is required (-token or TOOL_TOKEN)".
//
// Запуск: go run ./03_exercises/03_env_fallback
package main

import (
	"flag"
	"fmt"
	"io"
)

var _ = flag.NewFlagSet
var _ = io.Discard

func token(args []string, getenv func(string) string) (string, error) {
	// TODO: разберите -token, определите через fs.Visit, был ли он задан.
	return "", nil
}

func main() {
	fmt.Println("--- Проверка 03_env_fallback ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	env := func(m map[string]string) func(string) string { return func(k string) string { return m[k] } }
	withEnv := env(map[string]string{"TOOL_TOKEN": "from-env"})

	v, err := token([]string{"-token", "from-flag"}, withEnv)
	check(err == nil && v == "from-flag", "флаг важнее окружения")
	v, err = token(nil, withEnv)
	check(err == nil && v == "from-env", "без флага → TOOL_TOKEN")
	_, err = token(nil, env(nil))
	check(err != nil, "нигде не задан → ошибка")
	_, err = token([]string{"-token", ""}, withEnv)
	check(err != nil, "явно пустой флаг не подменяется окружением → ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
