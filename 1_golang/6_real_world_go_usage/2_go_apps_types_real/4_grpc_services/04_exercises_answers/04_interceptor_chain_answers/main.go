// Упражнение 04: цепочка unary-интерсепторов (Решение).
package main

import (
	"context"
	"fmt"
	"strings"
)

type handler func(ctx context.Context, req any) (any, error)
type interceptor func(ctx context.Context, req any, next handler) (any, error)

func chain(interceptors ...interceptor) interceptor {
	return func(ctx context.Context, req any, final handler) (any, error) {
		h := final
		for i := len(interceptors) - 1; i >= 0; i-- {
			ic, next := interceptors[i], h
			h = func(ctx context.Context, req any) (any, error) { return ic(ctx, req, next) }
		}
		return h(ctx, req)
	}
}

func main() {
	fmt.Println("--- Проверка 04_interceptor_chain (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var trace []string
	mark := func(name string) interceptor {
		return func(ctx context.Context, req any, next handler) (any, error) {
			trace = append(trace, name+">")
			resp, err := next(ctx, req)
			trace = append(trace, "<"+name)
			return resp, err
		}
	}
	final := func(ctx context.Context, req any) (any, error) {
		trace = append(trace, "h")
		return "pong:" + req.(string), nil
	}

	resp, err := chain(mark("log"), mark("auth"), mark("metrics"))(context.Background(), "ping", final)
	check(err == nil && resp == "pong:ping", "ответ handler'а проходит насквозь")
	got := strings.Join(trace, " ")
	check(got == "log> auth> metrics> h <metrics <auth <log", "порядок: "+got)

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
