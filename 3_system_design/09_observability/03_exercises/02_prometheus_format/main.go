// Упражнение 02: форматирование метрики в Prometheus-формат.
//
// Задача: реализовать formatMetric — вернуть одну строку экспозиции
// Prometheus вида:
//
//	name{k1="v1",k2="v2"} value
//
// Правила:
//   - метки идут в порядке, заданном срезом labels (не сортировать);
//   - значение метки оборачивается в двойные кавычки;
//   - если меток нет — вернуть "name value" (без фигурных скобок);
//   - строка НЕ заканчивается символом перевода строки.
//
// Реализуйте formatMetric.
//
// Запуск: go run ./03_exercises/02_prometheus_format/main.go
package main

import (
	"fmt"
	"strings"
)

// label — пара имя/значение метки.
type label struct {
	Key   string
	Value string
}

// Держим импорт strings живым (в решении удобен strings.Builder).
var _ = strings.TrimSpace

// formatMetric форматирует метрику в одну строку Prometheus-экспозиции.
func formatMetric(name string, labels []label, value int64) string {
	// TODO:
	//  1) если len(labels)==0 -> вернуть fmt.Sprintf("%s %d", name, value);
	//  2) иначе собрать k="v" через запятую, обернуть в {...};
	//  3) вернуть name{...} value без завершающего '\n'.
	return ""
}

func main() {
	fmt.Println("--- Проверка 02_prometheus_format ---")

	cases := []struct {
		name   string
		labels []label
		value  int64
		want   string
	}{
		{"app_up", nil, 1, "app_up 1"},
		{
			"http_requests_total",
			[]label{{"method", "GET"}, {"status", "200"}},
			1027,
			`http_requests_total{method="GET",status="200"} 1027`,
		},
		{
			"queue_depth",
			[]label{{"queue", "emails"}},
			42,
			`queue_depth{queue="emails"} 42`,
		},
	}

	allOK := true
	for _, c := range cases {
		got := formatMetric(c.name, c.labels, c.value)
		if got != c.want {
			allOK = false
			fmt.Printf("FAIL: formatMetric(%q,...) =\n  %q\n  ожидалось %q\n", c.name, got, c.want)
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
