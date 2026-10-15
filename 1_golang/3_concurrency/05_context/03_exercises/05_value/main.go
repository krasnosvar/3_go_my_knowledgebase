// Упражнение 05: request-id в context.WithValue.
//
// Задача:
//   - withRequestID(ctx, id) кладёт id в контекст через context.WithValue;
//   - requestID(ctx) достаёт его: (id, true) или ("", false).
//
// Ключ — собственный неэкспортируемый тип ctxKey: это исключает коллизии
// с ключами других пакетов. Значение доставайте через type assertion
// с проверкой: v, ok := ctx.Value(requestIDKey).(string).
//
// Запуск: go run ./03_exercises/05_value/main.go
package main

import (
	"context"
	"fmt"
)

type ctxKey struct{}

var requestIDKey ctxKey

// withRequestID возвращает контекст с request-id.
func withRequestID(ctx context.Context, id string) context.Context {
	// TODO: context.WithValue(ctx, requestIDKey, id)
	return ctx
}

// requestID достаёт request-id из контекста.
func requestID(ctx context.Context) (string, bool) {
	// TODO: type assertion с ok.
	return "", false
}

func main() {
	fmt.Println("--- Проверка 05_value ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ctx := withRequestID(context.Background(), "req-42")
	id, ok := requestID(ctx)
	check(ok && id == "req-42", fmt.Sprintf("requestID = (%q, %t), ожидалось (\"req-42\", true)", id, ok))

	// значение видно и в производных контекстах
	child, cancel := context.WithCancel(ctx)
	defer cancel()
	id, ok = requestID(child)
	check(ok && id == "req-42", "request-id должен быть виден в дочернем контексте")

	// чужой ключ-строка с тем же текстом не должен совпасть
	_, ok = requestID(context.WithValue(context.Background(), "requestID", "x"))
	check(!ok, "строковый ключ не должен совпадать с типизированным ключом")

	_, ok = requestID(context.Background())
	check(!ok, "в пустом контексте request-id отсутствует")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
