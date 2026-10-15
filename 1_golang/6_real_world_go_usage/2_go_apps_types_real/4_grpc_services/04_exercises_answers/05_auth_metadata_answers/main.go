// Упражнение 05: аутентификация по metadata (Решение).
package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type mdKey struct{}
type userKey struct{}

type md map[string][]string

type handler func(ctx context.Context, req any) (any, error)
type interceptor func(ctx context.Context, req any, next handler) (any, error)

var errUnauthenticated = errors.New("unauthenticated")

func authInterceptor(valid map[string]string) interceptor {
	return func(ctx context.Context, req any, next handler) (any, error) {
		m, _ := ctx.Value(mdKey{}).(md)
		vals := m["authorization"]
		if len(vals) == 0 {
			return nil, errUnauthenticated
		}
		tok, ok := strings.CutPrefix(vals[0], "Bearer ")
		uid, known := valid[tok]
		if !ok || !known {
			return nil, errUnauthenticated
		}
		return next(context.WithValue(ctx, userKey{}, uid), req)
	}
}

func main() {
	fmt.Println("--- Проверка 05_auth_metadata (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ic := authInterceptor(map[string]string{"t1": "user-1"})
	calls := 0
	final := func(ctx context.Context, req any) (any, error) {
		calls++
		return ctx.Value(userKey{}), nil
	}
	withMD := func(m md) context.Context { return context.WithValue(context.Background(), mdKey{}, m) }

	resp, err := ic(withMD(md{"authorization": {"Bearer t1"}}), nil, final)
	check(err == nil && resp == "user-1", "валидный токен → userID в ctx")
	_, err = ic(context.Background(), nil, final)
	check(errors.Is(err, errUnauthenticated), "без metadata → unauthenticated")
	_, err = ic(withMD(md{"authorization": {"Bearer bad"}}), nil, final)
	check(errors.Is(err, errUnauthenticated), "чужой токен → unauthenticated")
	_, err = ic(withMD(md{"authorization": {"t1"}}), nil, final)
	check(errors.Is(err, errUnauthenticated), "без префикса Bearer → unauthenticated")
	check(calls == 1, "handler вызван только для валидного запроса")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
