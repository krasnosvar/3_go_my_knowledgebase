// Упражнение 07: политика ретраев клиента.
//
// Ретраить можно только временные ошибки (Unavailable, ResourceExhausted);
// InvalidArgument или NotFound повтор не исправит.
//
// Задача: реализуйте callWithRetry(call func() code, maxAttempts int, base time.Duration, sleep func(time.Duration)) (code, int):
//   - вызвать call; OK или неретраибельный код → вернуть (код, попыток);
//   - ретраибельный → sleep(base * 2^(n-1)) (n — номер неудачной попытки,
//     т.е. base, 2*base, 4*base...) и повторить, пока попыток < maxAttempts;
//   - после maxAttempts вернуть последний код. Последний sleep не нужен.
//
// Запуск: go run ./03_exercises/07_retry_policy
package main

import (
	"fmt"
	"time"
)

type code int

const (
	OK                code = 0
	InvalidArgument   code = 3
	ResourceExhausted code = 8
	Unavailable       code = 14
)

func retryable(c code) bool { return c == Unavailable || c == ResourceExhausted }

func callWithRetry(call func() code, maxAttempts int, base time.Duration, sleep func(time.Duration)) (code, int) {
	// TODO: цикл попыток с экспоненциальной паузой.
	return call(), 1
}

func main() {
	fmt.Println("--- Проверка 07_retry_policy ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	seq := func(codes ...code) func() code {
		i := 0
		return func() code { c := codes[min(i, len(codes)-1)]; i++; return c }
	}
	var sleeps []time.Duration
	sleep := func(d time.Duration) { sleeps = append(sleeps, d) }

	c, n := callWithRetry(seq(Unavailable, Unavailable, OK), 5, 10*time.Millisecond, sleep)
	check(c == OK && n == 3, "два Unavailable, затем OK → OK за 3 попытки")
	check(fmt.Sprint(sleeps) == "[10ms 20ms]", "паузы 10ms, 20ms: "+fmt.Sprint(sleeps))

	sleeps = nil
	c, n = callWithRetry(seq(InvalidArgument), 5, 10*time.Millisecond, sleep)
	check(c == InvalidArgument && n == 1 && len(sleeps) == 0, "InvalidArgument не ретраится")

	sleeps = nil
	c, n = callWithRetry(seq(Unavailable), 3, time.Millisecond, sleep)
	check(c == Unavailable && n == 3 && len(sleeps) == 2, "исчерпание попыток: 3 вызова, 2 паузы")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
