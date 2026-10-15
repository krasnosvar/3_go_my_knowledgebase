// Упражнение 03: экспоненциальный backoff с потолком (Решение).
package main

import (
	"errors"
	"fmt"
	"time"
)

func backoff(attempt int, base, max time.Duration) time.Duration {
	d := base
	for i := 1; i < attempt; i++ {
		d *= 2
		if d >= max {
			return max
		}
	}
	return min(d, max)
}

func retry(fn func() error, attempts int, base, max time.Duration, sleep func(time.Duration)) error {
	var err error
	for n := 1; n <= attempts; n++ {
		if err = fn(); err == nil {
			return nil
		}
		if n < attempts {
			sleep(backoff(n, base, max))
		}
	}
	return fmt.Errorf("after %d attempts: %w", attempts, err)
}

func main() {
	fmt.Println("--- Проверка 03_retry_backoff (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	ms := time.Millisecond
	check(backoff(1, 100*ms, time.Second) == 100*ms, "attempt 1 → base")
	check(backoff(3, 100*ms, time.Second) == 400*ms, "attempt 3 → 4*base")
	check(backoff(10, 100*ms, time.Second) == time.Second, "attempt 10 → потолок max")
	check(backoff(200, ms, time.Minute) == time.Minute, "огромный attempt не переполняется")

	var sleeps []time.Duration
	sleep := func(d time.Duration) { sleeps = append(sleeps, d) }
	calls := 0
	err := retry(func() error {
		calls++
		if calls < 3 {
			return errors.New("temporary")
		}
		return nil
	}, 5, 10*ms, time.Second, sleep)
	check(err == nil && calls == 3 && fmt.Sprint(sleeps) == "[10ms 20ms]", "успех на 3-й попытке, паузы 10ms 20ms")

	boom := errors.New("down")
	sleeps = nil
	err = retry(func() error { return boom }, 3, ms, time.Second, sleep)
	check(errors.Is(err, boom) && len(sleeps) == 2, "после 3 неудач — обёрнутая последняя ошибка, 2 паузы")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
