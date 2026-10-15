// Упражнение 04: TTL с jitter против одновременного протухания.
//
// Если тысячу ключей положить с одинаковым TTL, они протухнут в одну
// секунду и все запросы разом пойдут в базу. Лечится случайной добавкой.
//
// Задача: реализуйте ttlWithJitter(base time.Duration, pct int, rnd func(n int64) int64) time.Duration:
//   - jitter — случайная величина в диапазоне [0, base*pct/100), получаемая
//     как rnd(base*pct/100) (rnd(n) возвращает число из [0, n));
//   - вернуть base + jitter; при pct <= 0 вернуть base без вызова rnd.
//
// rnd передаётся параметром (в проде — rand.Int63n), тест подставляет фиксированный.
//
// Запуск: go run ./03_exercises/04_ttl_jitter
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ttlWithJitter(base time.Duration, pct int, rnd func(n int64) int64) time.Duration {
	// TODO: посчитайте ширину окна base*pct/100 и добавьте rnd(окно).
	return 0
}

func main() {
	fmt.Println("--- Проверка 04_ttl_jitter ---")

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
