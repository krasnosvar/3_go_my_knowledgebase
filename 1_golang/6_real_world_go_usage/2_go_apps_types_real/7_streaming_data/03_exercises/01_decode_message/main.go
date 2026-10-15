// Упражнение 01: строгий разбор сообщения из топика.
//
// Задача: реализуйте decode(raw []byte) (orderEvent, error):
//   - json.NewDecoder(bytes.NewReader(raw)) + DisallowUnknownFields
//     (неизвестное поле — признак смены схемы, лучше упасть явно);
//   - Version должен быть 1 → иначе errUnsupportedVersion;
//   - OrderID непустой и Amount > 0 → иначе errInvalid;
//   - ошибка JSON → обернуть: fmt.Errorf("decode: %w", err).
//
// Запуск: go run ./03_exercises/01_decode_message
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type orderEvent struct {
	Version int    `json:"version"`
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
}

var (
	errUnsupportedVersion = errors.New("unsupported version")
	errInvalid            = errors.New("invalid event")
)

var _ = bytes.NewReader
var _ = json.NewDecoder

func decode(raw []byte) (orderEvent, error) {
	// TODO: строгий декодер + проверки версии и полей.
	return orderEvent{}, nil
}

func main() {
	fmt.Println("--- Проверка 01_decode_message ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	e, err := decode([]byte(`{"version":1,"order_id":"o-1","amount":500}`))
	check(err == nil && e.OrderID == "o-1" && e.Amount == 500, "валидное событие разобрано")
	_, err = decode([]byte(`{"version":2,"order_id":"o-1","amount":5}`))
	check(errors.Is(err, errUnsupportedVersion), "version=2 → errUnsupportedVersion")
	_, err = decode([]byte(`{"version":1,"order_id":"","amount":5}`))
	check(errors.Is(err, errInvalid), "пустой order_id → errInvalid")
	_, err = decode([]byte(`{"version":1,"order_id":"o","amount":5,"currency":"EUR"}`))
	check(err != nil && !errors.Is(err, errInvalid), "неизвестное поле → ошибка декодера")
	_, err = decode([]byte(`not json`))
	check(err != nil, "мусор → ошибка")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
