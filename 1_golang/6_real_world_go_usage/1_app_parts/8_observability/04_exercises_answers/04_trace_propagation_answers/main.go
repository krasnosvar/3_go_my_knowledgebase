// Упражнение 04: проброс trace ID между сервисами (Решение).
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type traceKey struct{}

func extract(r *http.Request, gen func() string) context.Context {
	id := r.Header.Get("X-Trace-ID")
	if id == "" {
		id = gen()
	}
	return context.WithValue(r.Context(), traceKey{}, id)
}

func inject(ctx context.Context, req *http.Request) {
	if id, ok := ctx.Value(traceKey{}).(string); ok {
		req.Header.Set("X-Trace-ID", id)
	}
}

func main() {
	fmt.Println("--- Проверка 04_trace_propagation (Решение) ---")

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
