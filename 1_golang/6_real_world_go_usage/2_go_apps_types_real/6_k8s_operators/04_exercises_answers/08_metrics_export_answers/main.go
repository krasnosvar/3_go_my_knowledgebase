// Упражнение 08: метрики контроллера (Решение).
package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type reconcileMetrics struct{ counts map[string]int }

func (m *reconcileMetrics) Observe(controller string, err error) {
	result := "success"
	if err != nil {
		result = "error"
	}
	key := fmt.Sprintf("controller_runtime_reconcile_total{controller=%q,result=%q}", controller, result)
	m.counts[key]++
}

func (m *reconcileMetrics) Render() string {
	keys := make([]string, 0, len(m.counts))
	for k := range m.counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		fmt.Fprintf(&b, "%s %d\n", k, m.counts[k])
	}
	return b.String()
}

func main() {
	fmt.Println("--- Проверка 08_metrics_export (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	m := &reconcileMetrics{counts: map[string]int{}}
	boom := errors.New("conflict")
	m.Observe("bucket", nil)
	m.Observe("bucket", nil)
	m.Observe("bucket", boom)
	m.Observe("app", nil)

	want := `controller_runtime_reconcile_total{controller="app",result="success"} 1
controller_runtime_reconcile_total{controller="bucket",result="error"} 1
controller_runtime_reconcile_total{controller="bucket",result="success"} 2
`
	got := m.Render()
	check(got == want, "Render:\n"+got)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
