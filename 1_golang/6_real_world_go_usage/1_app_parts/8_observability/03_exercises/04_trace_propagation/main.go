// Упражнение 04: проброс trace ID между сервисами.
//
// Задача:
//   - extract(r *http.Request, gen func() string) context.Context: взять
//     заголовок "traceparent"-подобный X-Trace-ID; если нет — gen();
//     положить в context по ключу traceKey{};
//   - inject(ctx, req *http.Request): записать trace ID из ctx в заголовок
//     X-Trace-ID исходящего запроса (если в ctx его нет — ничего не делать).
//
// Так один ID проходит через всю цепочку вызовов и связывает логи.
//
// Запуск: go run ./03_exercises/04_trace_propagation
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type traceKey struct{}

func extract(r *http.Request, gen func() string) context.Context {
	// TODO: прочитайте X-Trace-ID или сгенерируйте, положите в context.
	return r.Context()
}

func inject(ctx context.Context, req *http.Request) {
	// TODO: если в ctx есть trace ID — запишите его в заголовок.
}

func main() {
	fmt.Println("--- Проверка 04_trace_propagation ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	in := httptest.NewRequest("GET", "/a", nil)
	in.Header.Set("X-Trace-ID", "t-123")
	ctx := extract(in, func() string { return "new" })
	out := httptest.NewRequest("GET", "/downstream", nil)
	inject(ctx, out)
	check(out.Header.Get("X-Trace-ID") == "t-123", "входящий trace ID проброшен в исходящий запрос")

	ctx = extract(httptest.NewRequest("GET", "/b", nil), func() string { return "gen-7" })
	out = httptest.NewRequest("GET", "/downstream", nil)
	inject(ctx, out)
	check(out.Header.Get("X-Trace-ID") == "gen-7", "без входящего ID генерируется новый")

	out = httptest.NewRequest("GET", "/downstream", nil)
	inject(context.Background(), out)
	check(out.Header.Get("X-Trace-ID") == "", "пустой context → заголовок не ставится")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
