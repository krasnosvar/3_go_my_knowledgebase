// Упражнение 04: at-least-once — коммит только после успеха.
//
// Задача: реализуйте consume(msgs []message, handle func(message) error, retries int) (committed int64, err error):
//   - сообщения обрабатываются по порядку;
//   - каждое — до 1+retries попыток handle;
//   - после успеха сообщения committed = его Offset;
//   - если сообщение так и не обработалось — остановиться и вернуть
//     (последний закоммиченный оффсет, ошибку): после рестарта чтение
//     продолжится с него, сообщение не потеряется;
//   - committed изначально -1.
//
// Запуск: go run ./03_exercises/04_retry_on_fail
package main

import (
	"errors"
	"fmt"
)

type message struct {
	Offset int64
	Value  string
}

func consume(msgs []message, handle func(message) error, retries int) (committed int64, err error) {
	committed = -1
	// TODO: цикл по сообщениям, внутренний цикл попыток, коммит после успеха.
	return committed, nil
}

func main() {
	fmt.Println("--- Проверка 04_retry_on_fail ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	msgs := []message{{0, "a"}, {1, "flaky"}, {2, "b"}, {3, "broken"}, {4, "c"}}
	attempts := map[string]int{}
	var processed []string
	handle := func(m message) error {
		attempts[m.Value]++
		if m.Value == "flaky" && attempts[m.Value] < 3 {
			return errors.New("timeout")
		}
		if m.Value == "broken" {
			return errors.New("db down")
		}
		processed = append(processed, m.Value)
		return nil
	}

	committed, err := consume(msgs, handle, 2)
	check(err != nil, "на broken возвращается ошибка")
	check(committed == 2, fmt.Sprintf("закоммичен оффсет 2, не дальше (получено %d)", committed))
	check(attempts["flaky"] == 3 && attempts["broken"] == 3, "flaky и broken получили 1+2 попытки")
	check(fmt.Sprint(processed) == "[a flaky b]", "c не обработано после остановки")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
