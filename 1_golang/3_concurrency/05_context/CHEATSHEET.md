# Пакет context — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии Go указаны там, где поведение зависит от них.

## API

Подробно: [01_basics.md](./01_theory/01_basics.md).

| Задача | Код |
| :--- | :--- |
| Корневой контекст (main, init, тесты) | `context.Background()` |
| «Пока не знаю, какой» | `context.TODO()` — вместо `nil` |
| Явная отмена | `ctx, cancel := context.WithCancel(parent); defer cancel()` |
| Отмена по таймауту | `ctx, cancel := context.WithTimeout(parent, 2*time.Second); defer cancel()` |
| Отмена в момент времени | `ctx, cancel := context.WithDeadline(parent, t); defer cancel()` |
| Значение запроса | `ctx = context.WithValue(parent, key, v)` |
| Канал отмены | `<-ctx.Done()` (закрывается при отмене) |
| Причина | `ctx.Err()`: `nil` / `context.Canceled` / `context.DeadlineExceeded` |
| Дедлайн и остаток | `if dl, ok := ctx.Deadline(); ok { left := time.Until(dl) }` |
| Значение | `ctx.Value(key)` — ищет вверх по цепочке |
| Отмена по SIGINT/SIGTERM (Go 1.16+) | `ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)` |
| Отмена со своей причиной (Go 1.20+) | `ctx, cancel := context.WithCancelCause(parent)`; `cancel(err)`; `context.Cause(ctx)` |
| Таймаут со своей причиной (Go 1.21+) | `context.WithTimeoutCause(parent, d, err)` |
| Колбэк при отмене (Go 1.21+) | `stop := context.AfterFunc(ctx, f)` |
| Значения без отмены родителя (Go 1.21+) | `context.WithoutCancel(ctx)` — для фоновой работы после ответа |

## Правила

Подробно: [01_basics.md](./01_theory/01_basics.md), [02_cancel.md](./01_theory/02_cancel.md).

| Правило | Почему |
| :--- | :--- |
| `ctx` — первый параметр, имя `ctx` | `func F(ctx context.Context, ...)` — единый стиль |
| `defer cancel()` сразу после `WithCancel`/`WithTimeout`/`WithDeadline` | иначе дочерний контекст и таймер живут до отмены родителя; `cancel()` идемпотентен |
| Не хранить `ctx` в полях структур | контекст — область жизни операции; исключение — `http.Request.Context()` |
| Не передавать `nil` | `nil.Done()` → panic; используйте `context.TODO()` |
| Отмена родителя отменяет всех потомков | распространение только сверху вниз |
| Контекст не останавливает код сам | функция обязана проверять `ctx.Done()` / `ctx.Err()` |

## Таймауты и дедлайны

Подробно: [03_timeout_deadline.md](./01_theory/03_timeout_deadline.md).

| Факт | Детали |
| :--- | :--- |
| `WithTimeout(p, d)` | ≡ `WithDeadline(p, time.Now().Add(d))` |
| Дочерний дедлайн позже родительского | не продлевает: действует более ранний |
| Бюджет запроса | handler задаёт общий таймаут, нижние слои нарезают дочерние `WithTimeout` |
| Внешний вызов | таймаут обязателен; передавать через `ctx`, а не `time.After` поверх |
| HTTP-статус по ошибке | `errors.Is(err, context.DeadlineExceeded)` → `504 Gateway Timeout` |

| `context.Canceled` | `context.DeadlineExceeded` |
| :--- | :--- |
| кто-то вызвал `cancel()` (клиент ушёл, сервис останавливается) | истёк таймаут |
| штатное завершение — не логировать как error | сигнал о проблеме ниже по стеку или в бюджете |

## Значения

Подробно: [04_values.md](./01_theory/04_values.md).

| Класть | Не класть |
| :--- | :--- |
| request ID, trace ID / span | бизнес-параметры функции |
| аутентифицированный пользователь, роль, tenant | опции, меняющие поведение |
| логгер с привязанными полями | большие или изменяемые объекты |
| | зависимости: БД, HTTP-клиенты (это поля сервиса) |

| Приём | Код |
| :--- | :--- |
| Свой тип ключа (нет коллизий между пакетами) | `type ctxKey struct{ name string }`; `var userKey = ctxKey{"user"}` |
| Типизированный аксессор | `func UserFrom(ctx context.Context) (User, bool) { u, ok := ctx.Value(userKey).(User); return u, ok }` |
| Middleware | `next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey, u)))` |
| Стоимость `Value` | линейна по длине цепочки — один `WithValue` в начале запроса, не десятки |

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type ctxKey struct{ name string }

var reqIDKey = ctxKey{"req_id"}

func RequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(reqIDKey).(string)
	return id, ok
}

func slowOp(ctx context.Context) error {
	select {
	case <-time.After(time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx := context.WithValue(context.Background(), reqIDKey, "r-42")

	tctx, cancel := context.WithTimeout(ctx, 20*time.Millisecond)
	defer cancel()
	err := slowOp(tctx)
	fmt.Println(errors.Is(err, context.DeadlineExceeded)) // true

	id, _ := RequestID(tctx) // значения наследуются дочерними
	fmt.Println(id)          // r-42

	// дочерний не может пережить родителя
	long, cancel2 := context.WithTimeout(tctx, time.Hour)
	defer cancel2()
	d1, _ := tctx.Deadline()
	d2, _ := long.Deadline()
	fmt.Println(d1.Equal(d2)) // true

	cctx, cancel3 := context.WithCancelCause(ctx) // Go 1.20+
	cancel3(errors.New("shutdown"))
	fmt.Println(cctx.Err(), "|", context.Cause(cctx)) // context canceled | shutdown
}
```

## Где применять

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md).

| Место | Как |
| :--- | :--- |
| HTTP handler | `ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)` → передать вниз |
| БД | `db.QueryContext(ctx, ...)` — отмена прерывает запрос (если драйвер умеет) |
| Исходящий HTTP | `httpClient.Do(req.WithContext(ctx))` |
| Воркеры | `for { select { case <-ctx.Done(): return; case j := <-jobs: ... } }` |
| Параллельные вызовы | `errgroup.WithContext(ctx)` — первая ошибка отменяет остальные |
| Не нужен | чистые функции без I/O и долгих вычислений |

## Ловушки

Подробно: [05_common_pitfalls.md](./01_theory/05_common_pitfalls.md).

- ⚠️ `ctx, _ := context.WithTimeout(...)` — потерян `cancel`; `go vet` (lostcancel) ругается.
- ⚠️ `doWork(nil)` → panic при обращении к `ctx.Done()`.
- ⚠️ `type Service struct{ ctx context.Context }` — антипаттерн.
- ⚠️ Долгий цикл без проверки `ctx.Done()` не реагирует на отмену.
- ⚠️ `context.WithValue(ctx, "amount", 100)` — бизнес-данные и строковый ключ.
- ⚠️ Ожидание, что `WithTimeout(parent, 10*time.Minute)` продлит дедлайн родителя.
- ⚠️ `cancel` как общее состояние между горутинами без правил — держите `defer cancel()` там, где создан контекст.
- ⚠️ Логировать `context.Canceled` как ошибку — заливает метрики; `if errors.Is(err, context.Canceled) { return nil }`.
- ⚠️ `var ctx = context.Background()` на уровне пакета — контекст живёт в рамках операции.
