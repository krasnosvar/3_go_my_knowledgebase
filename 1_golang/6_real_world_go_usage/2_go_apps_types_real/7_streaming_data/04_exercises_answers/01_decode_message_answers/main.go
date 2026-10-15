// Упражнение 01: строгий разбор сообщения из топика (Решение).
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

func decode(raw []byte) (orderEvent, error) {
	var e orderEvent
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&e); err != nil {
		return orderEvent{}, fmt.Errorf("decode: %w", err)
	}
	if e.Version != 1 {
		return orderEvent{}, errUnsupportedVersion
	}
	if e.OrderID == "" || e.Amount <= 0 {
		return orderEvent{}, errInvalid
	}
	return e, nil
}

func main() {
	fmt.Println("--- Проверка 01_decode_message (Решение) ---")

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
