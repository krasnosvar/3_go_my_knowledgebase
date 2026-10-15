// Упражнение 05: аутентификация по metadata.
//
// В gRPC заголовки — это metadata (map[string][]string) во входящем
// context. Здесь metadata кладётся в context по ключу mdKey{}.
//
// Задача: реализуйте authInterceptor(valid map[string]string) interceptor:
//   - взять md из ctx; нет md или нет ключа "authorization" → errUnauthenticated;
//   - значение должно быть "Bearer <token>" и token есть в valid → положить
//     userID (valid[token]) в ctx по ключу userKey{} и вызвать next;
//   - иначе errUnauthenticated, next не вызывать.
//
// Запуск: go run ./03_exercises/05_auth_metadata
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

var _ = strings.CutPrefix

func authInterceptor(valid map[string]string) interceptor {
	return func(ctx context.Context, req any, next handler) (any, error) {
		// TODO: достаньте md, проверьте "authorization", положите userID в ctx.
		return nil, errUnauthenticated
	}
}

func main() {
	fmt.Println("--- Проверка 05_auth_metadata ---")

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
