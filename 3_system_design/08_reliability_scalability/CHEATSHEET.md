# Надёжность и масштабируемость — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## Балансировка нагрузки

Подробно: [01_load_balancing.md](./01_theory/01_load_balancing.md).

| Алгоритм | Как выбирает | Когда |
| :--- | :--- | :--- |
| Round-robin | по кругу A, B, C, A… | равные инстансы, однородные короткие запросы |
| Weighted round-robin | по кругу с весами | инстансы разной мощности |
| Least connections | меньше всего активных соединений | запросы разной длительности |
| Least response time | соединения + измеренная задержка | разнородная нагрузка |
| Hash по ключу | `hash(ip / session / user) % N` | sticky-сессии, локальные кэши |
| Consistent hashing | ближайший узел на кольце | шардирование, распределённые кэши: при смене узлов переезжает малая доля ключей |

| Аспект | L4 | L7 |
| :--- | :--- | :--- |
| Уровень | TCP/UDP | HTTP/gRPC |
| Что видит | IP, порт | URL, заголовки, метод, cookie |
| Маршрутизация по пути | нет | да |
| Терминация TLS | нет (passthrough) | да |
| Накладные расходы | минимальные | выше (парсинг) |
| Примеры | AWS NLB, IPVS, MetalLB | AWS ALB, nginx, Envoy, Ingress |

| Health check | Суть |
| :--- | :--- |
| Active | LB сам опрашивает `GET /healthz`; N успехов подряд → healthy, M провалов → из ротации |
| Passive | outlier detection по ошибкам реального трафика |
| Liveness | «процесс жив?» — провал → рестарт |
| Readiness | «готов к трафику?» — провал → вывод из ротации без рестарта |
| Гистерезис | несколько ответов подряд, чтобы не было flapping |

## Масштабирование

Подробно: [02_scaling.md](./01_theory/02_scaling.md).

| Аспект | Вертикальное (scale up) | Горизонтальное (scale out) |
| :--- | :--- | :--- |
| Способ | мощнее узел | больше узлов за LB |
| Потолок | жёсткий | почти нет |
| Отказоустойчивость | SPOF, рестарт = простой | резервирование |
| Сложность | низкая | распределённая система |
| Стоимость наверху | нелинейна | линейна на commodity |
| Когда | быстрый временный шаг; плохо шардируемое (primary реляционной БД) | stateless-веб-слой — почти всегда |

| Шардирование | Плюс | Минус |
| :--- | :--- | :--- |
| Range (A–M, N–Z) | просто, диапазонные запросы | hotspot при неравномерных ключах |
| Hash: `hash(key) % N` | равномерно | смена `N` переносит почти все ключи |
| Consistent hashing (+ виртуальные ноды) | при добавлении узла переезжает ~`1/N` ключей | сложнее реализовать |

Stateless: сессии → Redis/valkey, данные → БД, файлы → S3; любой инстанс обслуживает любой запрос.

## Rate limiting

Подробно: [03_rate_limiting.md](./01_theory/03_rate_limiting.md).

| Алгоритм | Механика | Всплески | Память | Граница окна |
| :--- | :--- | :--- | :--- | :--- |
| Token bucket | корзина `capacity`, +`rate` токенов/с, запрос берёт токен | разрешает burst | O(1) | — |
| Leaky bucket | очередь, вытекает ровно `rate`/с | сглаживает | O(размер очереди) | равномерный выход |
| Fixed window | счётчик на календарное окно | до 2×limit на стыке | O(1) | плохая |
| Sliding log | метки всех запросов за окно | точно | O(число запросов) | идеальная |
| Sliding window counter | `curr + prev × доля_prev` | сглаживает | O(1) | хорошая (аппроксимация) |

| Где | Как |
| :--- | :--- |
| Локально (per-instance) | in-memory; суммарный лимит = `N × локальный` |
| Глобально | общий счётчик в Redis/valkey: атомарный `INCR` + TTL |
| Ключ | IP, API-ключ, пользователь, эндпоинт |
| Ответ при превышении | `429 Too Many Requests` + `Retry-After` |

## Circuit breaker

Подробно: [04_circuit_breaker.md](./01_theory/04_circuit_breaker.md).

| Состояние | Поведение | Переход |
| :--- | :--- | :--- |
| Closed | запросы идут, считаем сбои | ошибок ≥ порога → Open |
| Open | fail fast, к зависимости не ходим | истёк open timeout → Half-Open |
| Half-Open | пропускаем несколько пробных | успех → Closed (сброс счётчиков); провал → Open |

| Параметр | Значение |
| :--- | :--- |
| Порог | N ошибок подряд или доля ошибок за окно |
| Что считать сбоем | таймауты, 5xx, ошибки соединения; **не** 4xx и бизнес-ошибки |
| Соседи | timeout (иначе зависание не видно), retry под breaker'ом, fallback (кэш/дефолт), bulkhead (изоляция пулов) |

## Таймауты и ретраи

Подробно: [05_retries_timeouts.md](./01_theory/05_retries_timeouts.md).

| Что | Как в Go / правило |
| :--- | :--- |
| Общий дедлайн | `context.WithTimeout(ctx, d)`; дочерний берёт минимум из своего и унаследованного |
| HTTP-клиент | `http.Client{Timeout: ...}`, `net.Dialer{Timeout: ...}` |
| HTTP-сервер | `http.Server{ReadTimeout: ..., WriteTimeout: ...}` |
| Значение таймаута | от реального p99 зависимости |
| Deadline propagation | `ctx` первым аргументом; между сервисами — в HTTP/gRPC-заголовках |
| Что ретраить | таймауты, сброс соединения, `5xx`, `429` |
| Что не ретраить | `4xx` (кроме `429`) |
| Число попыток | фиксированный max (например, 3) |
| Backoff | `base × 2^attempt` с потолком `max` |
| Full jitter | `sleep = random(0, min(max, base × 2^attempt))` |
| Безопасный повтор | только идемпотентное: `GET`, `PUT`, `DELETE`; `POST` — с `Idempotency-Key` |

```go
// Full jitter (Go 1.22+: math/rand/v2; встроенный min — Go 1.21+).
func backoff(attempt int, base, max time.Duration) time.Duration {
	return rand.N(min(max, base<<attempt) + 1)
}
```

## Graceful shutdown

Подробно: [06_graceful_shutdown.md](./01_theory/06_graceful_shutdown.md).

| Шаг | Действие |
| :--- | :--- |
| 1 | получили SIGTERM (k8s/systemd; Ctrl+C = SIGINT) |
| 2 | провалили readiness → LB выводит из ротации (нужен зазор) |
| 3 | перестали принимать новые соединения |
| 4 | дренаж активных запросов в пределах таймаута |
| 5 | закрыли пулы БД, брокеры, файлы; flush логов/метрик |
| 6 | exit 0; дренаж короче grace-периода (k8s: 30 с), иначе SIGKILL |

```go
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{Addr: "127.0.0.1:0", Handler: http.NewServeMux()}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	go func() { time.Sleep(100 * time.Millisecond); syscall.Kill(os.Getpid(), syscall.SIGTERM) }() // имитация k8s

	<-ctx.Done() // пришёл SIGTERM
	shCtx, cancel := context.WithTimeout(context.Background(), 25*time.Second) // < grace-периода (30 с)
	defer cancel()
	if err := srv.Shutdown(shCtx); err != nil { // стоп accept + дренаж активных
		log.Printf("shutdown: %v", err)
	}
	log.Println("остановлен")
}
```

| Что дренировать | Как |
| :--- | :--- |
| HTTP | `Server.Shutdown(ctx)`: закрывает listener, ждёт активные и idle keep-alive |
| WebSocket, стримы, gRPC-стримы | закрывать явно — `Shutdown` их не завершит |
| Воркеры | не брать новые задачи, доделать взятые или вернуть в очередь |
| Консьюмеры очередей | остановить чтение, дообработать, закоммитить/отклонить offset |

## Цепочка обороны

Подробно: [07_use_cases.md](./01_theory/07_use_cases.md).

```
клиент
  │  timeout на каждый вызов        (не висеть вечно)
  │  retry + backoff + jitter       (пережить транзиентный сбой, без волны)
  │  circuit breaker                (fail fast, не добивать больную зависимость)
  ▼
сервис
  │  rate limiting                  (не дать себя перегрузить)
  │  bulkhead / пулы                (изолировать ресурсы по зависимостям)
  ▼
зависимость (БД, соседний сервис)
```

## Ловушки

- ⚠️ Retry без backoff/jitter и без лимита попыток → retry storm, thundering herd.
- ⚠️ Retry неидемпотентной операции (`POST /charge`) → двойное списание.
- ⚠️ Нет таймаутов → исчерпание пулов, каскадный отказ; retry без таймаута вообще не сработает.
- ⚠️ Health check, дёргающий БД на каждый опрос → проверки сами валят здоровые инстансы.
- ⚠️ Liveness, зависящий от внешних зависимостей → рестарты по кругу вместо вывода из ротации.
- ⚠️ `hash(key) % N` для кэша → при добавлении узла почти все ключи мимо (массовый промах).
- ⚠️ Бизнес-ошибки (4xx) считаются сбоями breaker'а → цепь размыкается на здоровой зависимости.
- ⚠️ Stateful-веб-слой (локальные сессии) → нужны sticky-сессии, горизонтальное масштабирование ломается.
- ⚠️ Игнор SIGTERM или дренаж дольше grace-периода → SIGKILL, оборванные запросы на каждом деплое.
- ⚠️ Fixed window → до двойного лимита на стыке окон.
