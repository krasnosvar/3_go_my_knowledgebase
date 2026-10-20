# Производительность, наблюдаемость и кэширование в Go — шпаргалка

Подробности — в [01_theory/](./01_theory/) и каноническом разделе
[DevOps System Design](../../../5_devops_sre_knowledgebase/14_system_design/02_performance_caching_and_scaling/README.md).
⚠️ — частая ловушка.

## Тюнинг рантайма Go в контейнерах

Подробно:
[01_go_scheduler_memory_and_scaling.md](./01_theory/01_go_scheduler_memory_and_scaling.md).

| Переменная / Флаг | Назначение и рекомендация |
| :--- | :--- |
| `GOMAXPROCS` | Число логических процессоров `P`. В контейнерах с CPU limit задавать по лимиту (`automaxprocs` или Downward API; в Go 1.25+ учитывается автоматически). |
| `GOMEMLIMIT` | Мягкий предел памяти Go (Go 1.19+). Ставить **80–90% от лимита контейнера** (например, `850MiB` при `1GiB`), чтобы избежать `OOMKilled`. |
| `GOGC` | Процент роста живой кучи до следующего цикла GC (по умолчанию `100`). |
| `go build -gcflags="-m"` | Проверка escape analysis: какие переменные убегают со стека в кучу (`escapes to heap`). |

## Профилирование (`pprof`, `trace`) и логи (`log/slog`)

Подробно:
[02_observability_and_profiling_in_go.md](./01_theory/02_observability_and_profiling_in_go.md).

| Задача | Команда / Пакет |
| :--- | :--- |
| Снять CPU-профиль за 30 с | `go tool pprof http://127.0.0.1:9090/debug/pprof/profile?seconds=30` |
| Найти утечку памяти в куче | `go tool pprof -inuse_space http://127.0.0.1:9090/debug/pprof/heap` |
| Найти генератор мусора (аллокации) | `go tool pprof -alloc_objects http://127.0.0.1:9090/debug/pprof/allocs` |
| Посмотреть стеки всех горутин | `curl http://127.0.0.1:9090/debug/pprof/goroutine?debug=2` |
| Найти борьбу за мьютексы | `go tool pprof http://127.0.0.1:9090/debug/pprof/mutex` |
| Записать трейс планировщика | `curl -o trace.out http://127.0.0.1:9090/debug/pprof/trace?seconds=5 && go tool trace trace.out` |
| Быстрое JSON-логирование (Go 1.21+) | `logger.LogAttrs(ctx, slog.LevelInfo, "msg", slog.Int("id", 42))` |

## Потокобезопасный кэш в Go

Подробно: [03_caching_in_go.md](./01_theory/03_caching_in_go.md).

| Конструкция в Go | Когда применять |
| :--- | :--- |
| `map[K]V` + `sync.RWMutex` | Простой L1-кэш с умеренной конкуренцией за лок. |
| Шардированный `[256]shard` (`RWMutex` на бакет) | Высокий RPS на многоядерных машинах — устраняет lock contention. |
| `sync.Map` | Только write-once / read-many ключи или непересекающиеся ключи по ядрам. |
| `golang.org/x/sync/singleflight` | Подавление cache stampede (схлопывание параллельных промахов по одному ключу). |
| `[]byte` арена + `map[uint64]uint32` | Огромные кэши (миллионы записей) без указателей, чтобы GC не тратил время на сканирование (`noscan`). |

```go
package main

import (
	"fmt"
	"sync"
)

type SafeCache struct {
	mu sync.RWMutex
	m  map[string]string
}

func (c *SafeCache) Get(k string) (string, bool) {
	c.mu.RLock()
	v, ok := c.m[k]
	c.mu.RUnlock()
	return v, ok
}

func (c *SafeCache) Set(k, v string) {
	c.mu.Lock()
	c.m[k] = v
	c.mu.Unlock()
}

func main() {
	c := &SafeCache{m: make(map[string]string)}
	c.Set("lang", "Go")
	v, ok := c.Get("lang")
	fmt.Println(v, ok) // Go true
}
```

## Ловушки

- ⚠️ Запустить Go-сервис в контейнере с лимитом `1 CPU` на 64-ядерной ноде без
  настройки `GOMAXPROCS` (до Go 1.25) — жёсткий CFS throttling и деградация p99.
- ⚠️ Поставить `GOMEMLIMIT` равным 100% лимита контейнера без запаса под стек
  горутин и память ОС — приведёт к `OOMKilled` или бесконечному циклу GC.
- ⚠️ Выставить `/debug/pprof/` на внешний публичный HTTP-порт — утечка данных и
  DoS-уязвимость.
- ⚠️ Использовать `RLock()` в методе `Get()` у LRU-кэша, который внутри
  переставляет элемент в двусвязном списке (мутация списка под `RLock` вызывает
  гонку данных!).
- ⚠️ Передать клиентский `r.Context()` внутрь `singleflight.Do`: отмена первого
  запроса уронит все ожидающие горутины.
