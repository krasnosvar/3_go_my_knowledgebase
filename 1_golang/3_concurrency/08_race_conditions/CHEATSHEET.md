# Гонки, deadlock, livelock, starvation — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии Go указаны там, где поведение зависит от них.

## Виды ошибок

Подробно: [01_data_race.md](./01_theory/01_data_race.md), [02_deadlock.md](./01_theory/02_deadlock.md), [03_livelock_starvation.md](./01_theory/03_livelock_starvation.md).

| Ошибка | Симптом | Кто ловит |
| :--- | :--- | :--- |
| Data race | ≥ 2 горутины трогают переменную, хотя бы одна пишет, без синхронизации → неопределённое поведение | `-race` (только если гонка случилась во время прогона) |
| Race condition | ошибка из-за порядка событий (двойная оплата), может быть без data race | ревью, тесты, транзакции/блокировки |
| Deadlock | горутины ждут друг друга, прогресса нет | рантайм — только если спят **все** горутины: `fatal error: all goroutines are asleep - deadlock!` |
| Livelock | горутины активны (CPU греется), состояние не меняется | мониторинг latency/QPS |
| Starvation | горутина никогда не получает ресурс | мониторинг, pprof goroutine |

Каждый data race — race condition; не каждый race condition — data race.

## Race detector

Подробно: [04_race_detector.md](./01_theory/04_race_detector.md).

| Задача | Команда |
| :--- | :--- |
| Запуск с детектором | `go run -race ./cmd/app` |
| Тесты | `go test -race ./...` (в CI: `go test -race -timeout 5m ./...`) |
| Сборка | `go build -race -o app ./cmd/app` |
| Отчёт | `WARNING: DATA RACE` + стеки обеих горутин в stderr; тест падает, программа завершается с кодом 66 (`exit status 66`) |
| Стоимость | CPU и память в 5–10 раз больше → не в проде, обязательно в CI |
| Исключить файл из `-race` | `//go:build !race` (держать ноль таких) |
| Не ловит | deadlock, livelock, логические race condition, гонки в непрогнанном коде |
| Дополнительно | `staticcheck`, `golangci-lint`, `go vet` (copylocks, loopclosure) |

## Где гонки встречаются → фикс

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md), [05_fix_strategies.md](./01_theory/05_fix_strategies.md).

| Место | Фикс |
| :--- | :--- |
| Счётчик `count++`, в том числе пакетная переменная в HTTP-handler | `atomic.Int64` или mutex |
| Общий `map` (чтение + запись) | `map` + `sync.Mutex`; `sync.Map` для особых случаев |
| Флаг «готово» `var done bool` | `atomic.Bool`, канал, `context` |
| Ленивая инициализация `if x == nil { x = New() }` | `sync.Once` / `atomic.Pointer` + CAS |
| Кэш «проверил → посчитал → записал» | mutex + double-check или `singleflight` |
| Переменная цикла в горутине (до Go 1.22) | параметр замыкания или `v := v` |
| Общий `bytes.Buffer`, срез, логгер | mutex или отдельные буферы → merge; логгер-горутина с каналом |
| Глобальные данные в тестах с `t.Parallel()` | локальные данные в каждом тесте |

## Стратегии исправления (по порядку)

Подробно: [05_fix_strategies.md](./01_theory/05_fix_strategies.md).

| Шаг | Что |
| :--- | :--- |
| 1. Убрать общее состояние | локальная переменная, копия вместо ссылки, шардирование данных по горутинам |
| 2. Правильный примитив | счётчик/флаг → atomic; несколько полей → `Mutex`; read-heavy → `RWMutex` (измерить); передача → канал; один раз → `Once` |
| 3. Владение через канал | у объекта один владелец; отправил в канал — больше не трогаешь |
| 4. Copy-on-write | `atomic.Pointer[T]`: читатели `Load()`, писатель собирает копию и `Store()` |
| 5. Deadlock на mutex'ах | единый глобальный порядок захвата замков |
| 6. Зависания | таймауты через `context` на всё, что может повиснуть |
| 7. Проверка | stress-тест (1000 горутин) + `-race` + `t.Parallel()` |
| 8. Документировать | какие методы типа безопасны для параллельного вызова |

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var racy int          // go run -race → WARNING: DATA RACE
	var safe atomic.Int64 // фикс
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			racy++
			safe.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(safe.Load()) // 1000 (racy — любое значение ≤ 1000)
}
```

## Deadlock, livelock, starvation: источники и лечение

Подробно: [02_deadlock.md](./01_theory/02_deadlock.md), [03_livelock_starvation.md](./01_theory/03_livelock_starvation.md).

| Источник | Лечение |
| :--- | :--- |
| Замки `a`, `b` захватываются в разном порядке | глобальный порядок захвата; избегать lock внутри lock |
| Повторный `Lock` того же `sync.Mutex` | не реентерабелен — перестроить код (внутренние методы без захвата) |
| Отправка/приём без пары | в `main` — fatal deadlock; в горутине — утечка без паники |
| `wg.Wait()` при забытом `Done()` | `defer wg.Done()` первой строкой |
| Mutex во время I/O / пользовательского кода | скопировать под замком, работать без замка |
| Livelock: все откатываются и повторяют синхронно | jitter в backoff, лимит попыток + fallback, явный приоритет |
| Starvation: писатель `RWMutex` под потоком читателей | короче удержание, тайм-аут ожидания, `x/sync/semaphore` по весу |
| «Жива, но стоит» из-за downstream | таймаут на каждый внешний вызов, ретраи с backoff, circuit breaker |

| Диагностика | Как |
| :--- | :--- |
| Сколько горутин ждут | `GODEBUG=schedtrace=1000` |
| Где именно висят | pprof goroutine-профиль (стек каждой) |
| В проде | таймауты, liveness probes, heartbeat воркеров, метрики latency |

## Ловушки

- ⚠️ «Работает же» — data race проявляется под нагрузкой и на другом железе.
- ⚠️ Одновременная запись и чтение `map` → `fatal error: concurrent map writes` / `concurrent map read and map write`, `recover` не спасает (обнаружение не гарантировано).
- ⚠️ `time.Sleep` для «синхронизации» в примерах и тестах прячет гонку.
- ⚠️ Детектор deadlock рантайма молчит, если жива хоть одна горутина (например, HTTP-сервер).
- ⚠️ `close` закрытого канала — panic, а не deadlock, но из того же класса ошибок.
- ⚠️ Отчёт `-race` в CI «перезапустить до зелёного» — нельзя: читать отчёт и чинить.
- ⚠️ `bytes.Buffer`, `map`, срезы не потокобезопасны.
