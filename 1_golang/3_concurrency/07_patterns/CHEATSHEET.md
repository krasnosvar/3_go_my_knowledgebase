# Паттерны конкурентности — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии Go указаны там, где поведение зависит от них.

## Задача → паттерн

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md).

| Задача | Паттерн | Инструмент |
| :--- | :--- | :--- |
| Несколько независимых запросов параллельно, первая ошибка всё отменяет | fan-out с ошибками | `errgroup.WithContext` |
| Обработка очереди с ограниченным параллелизмом | worker pool | N горутин + канал `jobs` |
| Большой набор → параллельно → собрать | fan-out / fan-in | worker'ы + `merge` |
| ETL, потоковая обработка шагами | pipeline | стадии на каналах |
| Не больше M одновременных обращений к ресурсу | semaphore | `make(chan struct{}, M)` / `x/sync/semaphore` |
| Ограничить RPS | rate limiter | `golang.org/x/time/rate` |
| Downstream падает — перестать его дёргать | circuit breaker | `sony/gobreaker` или своя обёртка |
| Отправлять реже, пачками | batching | ticker + буфер + flush по размеру/времени/`ctx.Done()` |
| Остановка по SIGTERM | graceful shutdown | `signal.NotifyContext` + `srv.Shutdown(ctx)` |
| Показать «воркер жив» | heartbeat | ticker в `select` |

## Worker pool

Подробно: [01_worker_pool.md](./01_theory/01_worker_pool.md).

| Аспект | Как |
| :--- | :--- |
| Размер для CPU-bound | `runtime.NumCPU()` |
| Размер для I/O-bound | десятки–сотни, но не больше пула соединений БД / лимитов API |
| «Доделать очередь и выйти» | `close(jobs)` → `for j := range jobs` завершается |
| «Бросить всё сейчас» | `cancel()` → ветка `<-ctx.Done()` |
| Закрыть канал результатов | `go func() { wg.Wait(); close(out) }()` |
| Паника в job | `recover` в обработчике, иначе падает весь процесс |

```go
package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
)

func workerPool(ctx context.Context, n int, jobs <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				select {
				case out <- j * j:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	go func() { wg.Wait(); close(out) }() // закрыть out, когда все worker'ы вышли
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)
	go func() {
		defer close(jobs) // закрыть jobs → worker'ы доработают и выйдут
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
	}()

	var res []int
	for r := range workerPool(ctx, 3, jobs) {
		res = append(res, r)
	}
	sort.Ints(res) // порядок результатов не гарантирован
	fmt.Println(res) // [1 4 9 16 25]

	// семафор на buffered-канале: не больше 2 одновременно
	sem := make(chan struct{}, 2)
	var mu sync.Mutex
	active, peak := 0, 0
	var wg sync.WaitGroup
	for range 6 {
		sem <- struct{}{} // acquire
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }() // release
			mu.Lock()
			active++
			peak = max(peak, active) // max — builtin Go 1.21+
			mu.Unlock()
			mu.Lock()
			active--
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(peak <= 2) // true
}
```

## Fan-out / fan-in и pipeline

Подробно: [02_fan_out_fan_in.md](./01_theory/02_fan_out_fan_in.md), [03_pipeline.md](./01_theory/03_pipeline.md).

| Правило | Детали |
| :--- | :--- |
| Fan-out | N горутин читают один входной канал |
| Fan-in (`merge`) | горутина на каждый вход пишет в общий `out`; `WaitGroup` + закрывающая горутина |
| Стадия pipeline | `func stage(ctx, in <-chan T) <-chan U`; внутри `defer close(out)` и `for v := range in` |
| Каждая стадия закрывает свой выход | закрытие «протекает» по цепочке до потребителя |
| Отмена стадии | отправка через `select { case out <- v: case <-ctx.Done(): return }` |
| Тяжёлая стадия | распараллелить: fan-out внутри стадии + merge |
| Ошибки стадий | отдельный `errCh` или `errgroup` |
| Fan-out для I/O | обычно 2–8 горутин; больше — перегруз downstream |

## Semaphore

Подробно: [04_semaphore.md](./01_theory/04_semaphore.md).

| Задача | Код |
| :--- | :--- |
| Захват / освобождение | `sem <- struct{}{}` / `defer func() { <-sem }()` |
| Захват с отменой | `select { case sem <- struct{}{}: case <-ctx.Done(): return ctx.Err() }` |
| Взвешенный (несколько слотов) | `s := semaphore.NewWeighted(10)`; `s.Acquire(ctx, 4)`; `s.Release(4)`; `s.TryAcquire(1)` |
| Semaphore или worker pool | фиксированный набор однотипных задач → pool; разовые горутины с лимитом → semaphore |

## Graceful shutdown

Подробно: [05_shutdown.md](./01_theory/05_shutdown.md).

| Шаг | Что делать |
| :--- | :--- |
| 1 | SIGTERM/SIGINT → отмена `ctx` (`signal.NotifyContext`) |
| 2 | readiness сразу отдаёт 503; liveness — ок |
| 3 | `srv.Shutdown(shutdownCtx)` с отдельным таймаутом: новые не принимаем, текущие доделываем |
| 4 | воркеры выходят по `ctx.Done()` или по закрытию `jobs` |
| 5 | `wg.Wait()` / `g.Wait()` |
| 6 | закрыть БД, MQ-клиенты, сбросить логи/метрики/трейсы, выйти |

```go
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{Addr: "127.0.0.1:8080", Handler: http.NotFoundHandler()}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-ctx.Done() // SIGINT/SIGTERM

	shCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shCtx); err != nil { // новые не принимаем, ждём текущие
		log.Println("shutdown:", err)
	}
	log.Println("stopped")
}
```

## Ловушки

- ⚠️ Канал результатов без читателя → worker'ы виснут на отправке.
- ⚠️ Забыли закрыть `out` в `merge`/pool → потребитель висит на `range`.
- ⚠️ `wg.Add` внутри горутины fan-out → `Wait` вернётся раньше.
- ⚠️ «Горутина на каждое сообщение» без лимита → рост горутин и утечки при медленном downstream.
- ⚠️ Стадия без `ctx.Done()` в отправке → при отмене блокируется навсегда.
- ⚠️ Нет `release` на error-path → семафор постепенно «затягивается» до полной блокировки.
- ⚠️ Ёмкость семафора 0 или 1 без причины → последовательное выполнение.
- ⚠️ Каждый модуль создаёт свой «глобальный» семафор → лимит не соблюдается.
- ⚠️ Не ловить SIGTERM или зависнуть дольше grace period → SIGKILL, потерянные данные.
- ⚠️ Закрыть БД до остановки handler'ов — порядок shutdown важен.
- ⚠️ `SIGKILL` перехватить нельзя.
- ⚠️ Для очень коротких операций накладные расходы каналов съедают выигрыш.
