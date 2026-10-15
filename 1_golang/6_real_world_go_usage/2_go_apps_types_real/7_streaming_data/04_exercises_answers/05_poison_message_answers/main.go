// Упражнение 05: poison message → DLQ (Решение).
package main

import (
	"errors"
	"fmt"
)

type message struct {
	Offset int64
	Value  string
}

func consume(msgs []message, handle func(message) error, maxAttempts int) (committed int64, dlq []message) {
	committed = -1
	for _, m := range msgs {
		ok := false
		for a := 0; a < maxAttempts && !ok; a++ {
			ok = handle(m) == nil
		}
		if !ok {
			dlq = append(dlq, m)
		}
		committed = m.Offset
	}
	return committed, dlq
}

func main() {
	fmt.Println("--- Проверка 05_poison_message (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	calls := map[string]int{}
	handle := func(m message) error {
		calls[m.Value]++
		if m.Value == "{broken" {
			return errors.New("invalid json")
		}
		return nil
	}
	msgs := []message{{10, "ok1"}, {11, "{broken"}, {12, "ok2"}}

	committed, dlq := consume(msgs, handle, 3)
	check(committed == 12, "поток не застрял: закоммичен последний оффсет 12")
	check(len(dlq) == 1 && dlq[0].Offset == 11, "битое сообщение 11 в DLQ")
	check(calls["{broken"] == 3 && calls["ok2"] == 1, "битое получило 3 попытки, следующие обработаны")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
