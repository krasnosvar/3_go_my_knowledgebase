// Упражнение 02: форматирование метрики в Prometheus-формат (Решение).
package main

import (
	"fmt"
	"strings"
)

type label struct {
	Key   string
	Value string
}

func formatMetric(name string, labels []label, value int64) string {
	if len(labels) == 0 {
		return fmt.Sprintf("%s %d", name, value)
	}
	var b strings.Builder
	b.WriteString(name)
	b.WriteByte('{')
	for i, l := range labels {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(l.Key)
		b.WriteString(`="`)
		b.WriteString(l.Value)
		b.WriteByte('"')
	}
	b.WriteByte('}')
	fmt.Fprintf(&b, " %d", value)
	return b.String()
}

func main() {
	fmt.Println("--- Проверка 02_prometheus_format (Решение) ---")

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
