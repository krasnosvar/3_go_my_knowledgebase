// Упражнение 06: чтение конфигурации из переменных окружения (Решение).
package main

import (
	"fmt"
	"os"
	"strconv"
)

// envInt читает целочисленную переменную окружения.
func envInt(key string, def int) (int, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("env %s=%q: %w", key, v, err)
	}
	return n, nil
}

func main() {
	fmt.Println("--- Проверка 06_env_read (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	const key = "KB_EXERCISE_WORKERS"
	defer os.Unsetenv(key)

	os.Unsetenv(key)
	n, err := envInt(key, 4)
	check(err == nil && n == 4, fmt.Sprintf("не задана: ожидалось 4, nil; получено %d, %v", n, err))

	os.Setenv(key, "16")
	n, err = envInt(key, 4)
	check(err == nil && n == 16, fmt.Sprintf("=16: ожидалось 16, nil; получено %d, %v", n, err))

	os.Setenv(key, "")
	_, err = envInt(key, 4)
	check(err != nil, "пустое значение должно давать ошибку, а не default")

	os.Setenv(key, "many")
	_, err = envInt(key, 4)
	check(err != nil, "нечисловое значение должно давать ошибку")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
