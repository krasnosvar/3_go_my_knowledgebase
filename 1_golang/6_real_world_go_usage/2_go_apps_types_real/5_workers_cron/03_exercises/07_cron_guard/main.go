// Упражнение 07: защита cron-задачи от наложения запусков.
//
// Если задача по расписанию «каждую минуту» иногда идёт 3 минуты,
// новые запуски не должны стартовать поверх текущего.
//
// Задача: реализуйте (g *guard) TryRun(fn func()) bool:
//   - если fn уже выполняется — сразу вернуть false и увеличить g.skipped;
//   - иначе выполнить fn синхронно и вернуть true;
//   - используйте atomic.Bool (CompareAndSwap) — без блокировок ожидания.
//
// Запуск: go run ./03_exercises/07_cron_guard
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
	// TODO: CompareAndSwap(false, true), defer Store(false).
	fn()
	return true
}

func main() {
	fmt.Println("--- Проверка 07_cron_guard ---")

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
