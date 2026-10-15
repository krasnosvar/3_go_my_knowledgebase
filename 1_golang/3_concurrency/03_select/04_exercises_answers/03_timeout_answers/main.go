// Упражнение 03: таймаут операции (Решение).
package main

import (
	"errors"
	"fmt"
	"time"
)

var errTimeout = errors.New("timeout")

func fetch(result <-chan string, timeout time.Duration) (string, error) {
	select {
	case s := <-result:
		return s, nil
	case <-time.After(timeout):
		return "", errTimeout
	}
}

func main() {
	fmt.Println("--- Проверка 03_timeout (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	fast := make(chan string, 1)
	fast <- "ok"
	s, err := fetch(fast, 50*time.Millisecond)
	check(err == nil && s == "ok", fmt.Sprintf("быстрый ответ: (%q, %v), ожидалось (\"ok\", nil)", s, err))

	slow := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		select {
		case slow <- "late":
		default:
		}
	}()
	start := time.Now()
	s, err = fetch(slow, 20*time.Millisecond)
	check(errors.Is(err, errTimeout), fmt.Sprintf("медленный ответ: ошибка %v, ожидалась errTimeout", err))
	check(s == "", "при таймауте строка должна быть пустой")
	check(time.Since(start) < 150*time.Millisecond, "fetch не вернулась по таймауту")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
