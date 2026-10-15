// Упражнение 06: requeue с per-item backoff.
//
// controller-runtime при ошибке reconcile возвращает ключ в очередь с
// экспоненциальной задержкой для ЭТОГО ключа (workqueue rate limiter).
//
// Задача: реализуйте rateLimiter:
//   - When(key) time.Duration: base * 2^failures[key], не больше max;
//     после вызова failures[key]++;
//   - Forget(key): сбросить счётчик ключа (вызывается после успешного reconcile);
//   - NumRequeues(key) int: текущее число неудач.
//
// Запуск: go run ./03_exercises/06_error_retry
package main

import (
	"fmt"
	"time"
)

type rateLimiter struct {
	base, max time.Duration
	failures  map[string]int
}

func (r *rateLimiter) When(key string) time.Duration {
	// TODO: задержка по числу неудач ключа, затем инкремент.
	return 0
}

func (r *rateLimiter) Forget(key string) {
	// TODO: удалите счётчик ключа.
}

func (r *rateLimiter) NumRequeues(key string) int { return r.failures[key] }

func main() {
	fmt.Println("--- Проверка 06_error_retry ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	rl := &rateLimiter{base: 5 * time.Millisecond, max: 100 * time.Millisecond, failures: map[string]int{}}
	var got []time.Duration
	for i := 0; i < 6; i++ {
		got = append(got, rl.When("ns/a"))
	}
	check(fmt.Sprint(got) == "[5ms 10ms 20ms 40ms 80ms 100ms]", "экспонента с потолком: "+fmt.Sprint(got))
	check(rl.When("ns/b") == 5*time.Millisecond, "у другого ключа свой счётчик")
	check(rl.NumRequeues("ns/a") == 6, "NumRequeues(a) = 6")
	rl.Forget("ns/a")
	check(rl.NumRequeues("ns/a") == 0 && rl.When("ns/a") == 5*time.Millisecond, "после Forget задержка снова base")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
