// Упражнение 06: requeue с per-item backoff (Решение).
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
	n := r.failures[key]
	r.failures[key] = n + 1
	d := r.base
	for i := 0; i < n; i++ {
		d *= 2
		if d >= r.max {
			return r.max
		}
	}
	return d
}

func (r *rateLimiter) Forget(key string) {
	delete(r.failures, key)
}

func (r *rateLimiter) NumRequeues(key string) int { return r.failures[key] }

func main() {
	fmt.Println("--- Проверка 06_error_retry (Решение) ---")

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
