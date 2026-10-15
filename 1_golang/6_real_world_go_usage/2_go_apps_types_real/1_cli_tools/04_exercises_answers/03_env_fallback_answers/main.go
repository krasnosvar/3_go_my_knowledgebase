// Упражнение 03: флаг с fallback на переменную окружения (Решение).
package main

import (
	"flag"
	"fmt"
	"io"
)

func token(args []string, getenv func(string) string) (string, error) {
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	tok := fs.String("token", "", "API token")
	if err := fs.Parse(args); err != nil {
		return "", err
	}
	set := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "token" {
			set = true
		}
	})
	v := *tok
	if !set {
		v = getenv("TOOL_TOKEN")
	}
	if v == "" {
		return "", fmt.Errorf("token is required (-token or TOOL_TOKEN)")
	}
	return v, nil
}

func main() {
	fmt.Println("--- Проверка 03_env_fallback (Решение) ---")

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
