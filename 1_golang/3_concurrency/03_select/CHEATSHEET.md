# select — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии Go указаны там, где поведение зависит от них.

## Семантика

Подробно: [01_basics.md](./01_theory/01_basics.md).

| Факт | Детали |
| :--- | :--- |
| Ветки | только операции с каналами: `v := <-ch`, `v, ok := <-ch`, `ch <- x` |
| Готовы несколько | выбирается **случайно**, приоритетов нет |
| Не готова ни одна, нет `default` | `select` блокируется |
| Не готова ни одна, есть `default` | сразу выполняется `default` (non-blocking) |
| `select {}` | блок навсегда |
| `break` внутри `select` | выходит из `select`, не из `for` → метка или `return` |
| `nil`-канал в ветке | ветка никогда не выбирается — способ её «отключить» |
| Таймер и контекст | тоже каналы: `<-time.After(d)`, `<-t.C`, `<-ctx.Done()` |

## Типовые формы

Подробно: [02_default.md](./01_theory/02_default.md), [03_timeout.md](./01_theory/03_timeout.md), [04_cancel.md](./01_theory/04_cancel.md).

| Задача | Код |
| :--- | :--- |
| Значение или таймаут | `case v := <-ch: ...` / `case <-time.After(2*time.Second): return ErrTimeout` |
| Значение или отмена | `case v := <-ch: ...` / `case <-ctx.Done(): return ctx.Err()` |
| Non-blocking приём | `select { case v := <-ch: ...; default: }` |
| Non-blocking отправка (drop) | `select { case ch <- v: default: dropped.Inc() }` |
| Долгоживущий цикл | `for { select { case <-ctx.Done(): return ctx.Err(); case j := <-jobs: process(j) } }` |
| Канал `jobs` может закрыться | `case j, ok := <-jobs: if !ok { return nil }` |
| Периодическая работа | `t := time.NewTicker(d); defer t.Stop()` + `case <-t.C:` |
| Выход из `for` | `Loop: for { select { case <-done: break Loop } }` |
| Graceful shutdown по сигналу | `case <-sigCh: cancel()` / `case err := <-serverErr: return err` |

## Таймаут: что выбрать

Подробно: [03_timeout.md](./01_theory/03_timeout.md).

| Ситуация | Инструмент |
| :--- | :--- |
| Разовый локальный таймаут | `time.After(d)` |
| Таймаут в горячем цикле | один `time.NewTimer(d)` + `Stop`/`Reset`, а не `time.After` на каждой итерации |
| Таймаут — часть общего бюджета запроса (handler → сервис → БД) | `context.WithTimeout` + `<-ctx.Done()` |
| Известна точка времени | `context.WithDeadline` |
| Отличить причину | `ctx.Err()`: `context.Canceled` (отменили) / `context.DeadlineExceeded` (истёк таймаут) |
| Утечка таймера `time.After` | до Go 1.23 таймер жил до срабатывания; с Go 1.23 (`go 1.23` в `go.mod`) несработавший таймер без ссылок собирает GC |

## Примеры

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md).

```go
package main

import (
	"fmt"
	"time"
)

func merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil // nil-канал отключает ветку
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

func gen(xs ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, x := range xs {
			ch <- x
		}
	}()
	return ch
}

func main() {
	sum := 0
	for v := range merge(gen(1, 2, 3), gen(10, 20)) {
		sum += v
	}
	fmt.Println(sum) // 36

	slow := make(chan int)
	t := time.NewTimer(50 * time.Millisecond)
	defer t.Stop()
	select {
	case v := <-slow:
		fmt.Println(v)
	case <-t.C:
		fmt.Println("timeout") // timeout
	}
}
```

| Use-case | Ветки `select` |
| :--- | :--- |
| «Первый из» | результат / `ctx.Done()` |
| Мультиплексор / event loop | `ctx.Done()` / команды / `ticker.C` / ошибки |
| Heartbeat воркера | `ctx.Done()` / `hb.C` → `status <- "alive"` / `jobs` |
| Fan-in двух каналов | см. `merge` выше |
| Простой rate limit | `<-time.Tick(100*time.Millisecond)` перед каждым запросом; для prod — `golang.org/x/time/rate` |
| `select` не нужен | одна операция с одним каналом — просто `<-ch` |

## Ловушки

Подробно: [05_pitfalls.md](./01_theory/05_pitfalls.md).

- ⚠️ `for { select { ...; default: } }` без паузы — busy loop, 100% CPU.
- ⚠️ `for { select { case j := <-jobs: } }` без `ctx.Done()` — горутину не остановить, утечка при остановке сервиса.
- ⚠️ `break` в `case` выходит только из `select`.
- ⚠️ Единственная ветка с `nil`-каналом → вечная блокировка.
- ⚠️ Нет приоритета веток: нужен приоритет — вложенный `select`.
- ⚠️ `default` в основной логике обработки — тихая потеря сообщений.
- ⚠️ `default + time.Sleep` вместо таймаута — используйте `time.After` / `context`.
- ⚠️ `time.After` в горячем цикле создаёт таймер на каждой итерации.
- ⚠️ Читатель ушёл по таймауту, писатель висит на отправке → результат-канал с буфером 1.
- ⚠️ Забыли `defer cancel()` после `WithTimeout`/`WithCancel` — ресурсы висят до срабатывания таймера.
