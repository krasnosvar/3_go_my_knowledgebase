// Упражнение 08: метрики контроллера.
//
// Задача: реализуйте (m *reconcileMetrics) Observe(controller string, err error)
// и Render() string. Формат строк (как у controller-runtime):
//
//	controller_runtime_reconcile_total{controller="<c>",result="success"} N
//	controller_runtime_reconcile_total{controller="<c>",result="error"} N
//
// Строки только для ненулевых значений, отсортированы, каждая с "\n".
//
// Запуск: go run ./03_exercises/08_metrics_export
package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type reconcileMetrics struct{ counts map[string]int }

var _ = sort.Strings
var _ = strings.Builder{}

func (m *reconcileMetrics) Observe(controller string, err error) {
	// TODO: result = "success" или "error", увеличьте счётчик по ключу-строке метрики.
}

func (m *reconcileMetrics) Render() string {
	// TODO: отсортированные строки "<ключ> <N>\n".
	return ""
}

func main() {
	fmt.Println("--- Проверка 08_metrics_export ---")

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
