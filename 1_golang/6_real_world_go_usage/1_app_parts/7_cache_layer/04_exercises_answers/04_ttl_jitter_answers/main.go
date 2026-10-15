// Упражнение 04: TTL с jitter (Решение).
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ttlWithJitter(base time.Duration, pct int, rnd func(n int64) int64) time.Duration {
	if pct <= 0 {
		return base
	}
	window := int64(base) * int64(pct) / 100
	return base + time.Duration(rnd(window))
}

func main() {
	fmt.Println("--- Проверка 04_ttl_jitter (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var asked int64
	maxRnd := func(n int64) int64 { asked = n; return n - 1 }
	zeroRnd := func(n int64) int64 { return 0 }

	got := ttlWithJitter(10*time.Second, 10, maxRnd)
	check(asked == int64(time.Second), "окно jitter = 10% от 10s = 1s")
	check(got == 10*time.Second+time.Second-1, "максимальный jitter < base+окно")
	check(ttlWithJitter(10*time.Second, 10, zeroRnd) == 10*time.Second, "нулевой jitter → base")
	check(ttlWithJitter(5*time.Second, 0, func(int64) int64 { panic("не вызывать") }) == 5*time.Second, "pct=0 → base без rnd")

	r := rand.New(rand.NewSource(1))
	for i := 0; i < 100; i++ {
		d := ttlWithJitter(time.Minute, 20, r.Int63n)
		if d < time.Minute || d >= time.Minute+12*time.Second {
			check(false, "случайный TTL вне [base, base+20%)")
			break
		}
	}

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
