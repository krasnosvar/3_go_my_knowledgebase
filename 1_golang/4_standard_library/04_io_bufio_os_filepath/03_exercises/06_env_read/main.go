// Упражнение 06: чтение конфигурации из переменных окружения.
//
// Задача: реализуйте envInt(key, def), которая читает целое из окружения:
//   - переменная не задана (os.LookupEnv вернул ok=false) → def, nil;
//   - задана, но пустая или не число → (0, ошибка с именем переменной);
//   - иначе → значение, nil.
//
// Важно различать «не задана» и «задана пустой»: os.Getenv этого не умеет,
// поэтому используйте os.LookupEnv.
//
// Запуск: go run ./03_exercises/06_env_read/main.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

var _ = strconv.Atoi // держим импорт живым в заготовке

// envInt читает целочисленную переменную окружения.
func envInt(key string, def int) (int, error) {
	// TODO: os.LookupEnv, при ok — strconv.Atoi с обёрткой ошибки.
	return def, nil
}

func main() {
	fmt.Println("--- Проверка 06_env_read ---")

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
