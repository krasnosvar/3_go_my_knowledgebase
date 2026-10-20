// Пример 02: самодельные метрики (counter/gauge) и эндпоинт /metrics
// в текстовом формате Prometheus.
//
// Реализуем потокобезопасные counter и gauge на sync/atomic, отдаём
// их через HTTP-хендлер в exposition-формате Prometheus, поднимаем
// сервер через httptest, делаем один запрос клиентом с таймаутом,
// печатаем тело /metrics — и программа завершается.
//
// Демонстрирует: atomic-счётчики, pull-модель (клиент забирает
// текст), формат "# HELP/# TYPE/имя{метки} значение".
//
// Запуск: go run ./02_examples/02_inprocess_metrics/main.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"time"
)

// counter — монотонно возрастающий счётчик (только Inc/Add).
type counter struct{ v atomic.Int64 }

func (c *counter) Add(n int64) { c.v.Add(n) }
func (c *counter) Get() int64  { return c.v.Load() }

// gauge — значение, которое может расти и падать (Set/Inc/Dec).
type gauge struct{ v atomic.Int64 }

func (g *gauge) Set(n int64) { g.v.Store(n) }
func (g *gauge) Add(n int64) { g.v.Add(n) }
func (g *gauge) Get() int64  { return g.v.Load() }

// registry держит все метрики процесса и умеет отдавать их текстом.
type registry struct {
	requests counter // общее число запросов (counter)
	errors   counter // число ошибок (counter)
	inflight gauge   // текущее число запросов в обработке (gauge)
}

// exposition собирает метрики в текстовый формат Prometheus.
func (r *registry) exposition() string {
	var b []byte
	b = append(b, "# HELP app_requests_total Total number of requests.\n"...)
	b = append(b, "# TYPE app_requests_total counter\n"...)
	b = append(b, fmt.Sprintf("app_requests_total %d\n", r.requests.Get())...)

	b = append(b, "# HELP app_errors_total Total number of failed requests.\n"...)
	b = append(b, "# TYPE app_errors_total counter\n"...)
	b = append(b, fmt.Sprintf("app_errors_total %d\n", r.errors.Get())...)

	b = append(b, "# HELP app_inflight_requests Requests currently in flight.\n"...)
	b = append(b, "# TYPE app_inflight_requests gauge\n"...)
	b = append(b, fmt.Sprintf("app_inflight_requests %d\n", r.inflight.Get())...)
	return string(b)
}

func main() {
	reg := &registry{}

	// Имитируем немного нагрузки, чтобы метрики были ненулевыми.
	reg.requests.Add(1027)
	reg.errors.Add(3)
	reg.inflight.Set(5)
	reg.inflight.Add(-2) // два запроса завершились -> 3 в обработке

	// /metrics-хендлер: Prometheus (pull-модель) сам ходит сюда и
	// забирает текущий срез. Здесь роль Prometheus играет наш клиент.
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		_, _ = io.WriteString(w, reg.exposition())
	})

	srv := httptest.NewServer(handler)
	defer srv.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(srv.URL + "/metrics")
	if err != nil {
		fmt.Println("ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("GET /metrics -> %d\n\n%s", resp.StatusCode, body)
	fmt.Println("Готово, программа завершается.")
}
