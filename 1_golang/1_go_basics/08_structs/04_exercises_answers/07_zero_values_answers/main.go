package main

import (
	"fmt"
)

type LogFormatter struct {
	Prefix string
}

func (lf LogFormatter) Format(msg string) string {
	prefix := lf.Prefix
	if prefix == "" {
		prefix = "[INFO] "
	}
	return prefix + msg
}

func main() {
	fmt.Println("--- Проверка 07_zero_values (Решение) ---")

	var defaultFormatter LogFormatter
	res1 := defaultFormatter.Format("тест 1")

	customFormatter := LogFormatter{Prefix: "[DEBUG] "}
	res2 := customFormatter.Format("тест 2")

	t1 := res1 == "[INFO] тест 1"
	t2 := res2 == "[DEBUG] тест 2"

	if t1 && t2 {
		fmt.Println("Тест нулевых значений: PASS")
	} else {
		fmt.Printf("Тест нулевых значений: FAIL (res1=%q, res2=%q)\n", res1, res2)
	}
}
