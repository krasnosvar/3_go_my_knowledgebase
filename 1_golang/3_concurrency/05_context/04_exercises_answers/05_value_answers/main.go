// Упражнение 05: request-id в context.WithValue (Решение).
package main

import (
	"context"
	"fmt"
)

type ctxKey struct{}

var requestIDKey ctxKey

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func requestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

func main() {
	fmt.Println("--- Проверка 05_value (Решение) ---")

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
