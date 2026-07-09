package main

import (
	"fmt"
)

type LogFormatter struct {
	Prefix string
}

// Format форматирует сообщение.
// Если Prefix не задан (пустая строка), метод должен использовать дефолтный префикс "[INFO] ".
// Иначе — использовать заданный Prefix.
func (lf LogFormatter) Format(msg string) string {
	// TODO: Реализуйте метод с учетом полезного нулевого значения.
	return ""
}

func main() {
	fmt.Println("--- Проверка 07_zero_values ---")

	// Нулевое значение структуры
	var defaultFormatter LogFormatter
	res1 := defaultFormatter.Format("тест 1") // "[INFO] тест 1"

	// Структура с заданным префиксом
	customFormatter := LogFormatter{Prefix: "[DEBUG] "}
	res2 := customFormatter.Format("тест 2") // "[DEBUG] тест 2"

	t1 := res1 == "[INFO] тест 1"
	t2 := res2 == "[DEBUG] тест 2"

	if t1 && t2 {
		fmt.Println("Тест нулевых значений: PASS")
	} else {
		fmt.Printf("Тест нулевых значений: FAIL (res1=%q, res2=%q)\n", res1, res2)
	}
}
