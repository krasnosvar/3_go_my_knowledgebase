// Упражнение 07: защита cron-задачи от наложения запусков (Решение).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type guard struct {
	running atomic.Bool
	skipped atomic.Int32
}

func (g *guard) TryRun(fn func()) bool {
	if !g.running.CompareAndSwap(false, true) {
		g.skipped.Add(1)
		return false
	}
	defer g.running.Store(false)
	fn()
	return true
}

func main() {
	fmt.Println("--- Проверка 07_cron_guard (Решение) ---")

	allOK := true
	check := func(cond bool, msg string) {
		if !cond {
			allOK = false
			fmt.Println("FAIL:", msg)
		}
	}

	var g guard
	started := make(chan struct{})
	release := make(chan struct{})
	var runs atomic.Int32

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		g.TryRun(func() { runs.Add(1); close(started); <-release })
	}()
	<-started // первая задача точно идёт

	for i := 0; i < 3; i++ {
		ok := g.TryRun(func() { runs.Add(1) })
		check(!ok, "запуск поверх текущего отклонён")
	}
	close(release)
	wg.Wait()

	check(runs.Load() == 1, "выполнился только первый запуск")
	check(g.skipped.Load() == 3, "пропущено 3 запуска")
	check(g.TryRun(func() {}), "после завершения можно запускать снова")

	if allOK {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
