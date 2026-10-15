# Context: основы

`context.Context` — стандартный способ пробрасывать через функции три вещи:

- **отмену** (cancellation),
- **таймаут/дедлайн**,
- **scoped-значения** (request ID, auth, tracing).

## Почему он нужен

Без контекста у библиотек нет единого языка «эй, мы больше не ждём». Сетевые клиенты, БД, очереди, gRPC — все ожидают `ctx` и умеют прерывать операцию.

## API

```go
ctx := context.Background()        // корневой, чаще всего в main
ctx := context.TODO()              // «пока не знаю, какой»

ctx, cancel := context.WithCancel(ctx)
defer cancel()

ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()

ctx, cancel := context.WithDeadline(ctx, t)
defer cancel()

ctx := context.WithValue(ctx, key, value)
```

## Семантика

- `ctx.Done()` → канал, который закрывается при отмене.
- `ctx.Err()` → `nil`, `context.Canceled` или `context.DeadlineExceeded`.
- `ctx.Deadline()` → время дедлайна, если есть.
- `ctx.Value(key)` → значение из цепочки.

## Базовые правила

- Context — **первый аргумент** функции: `func F(ctx context.Context, ...)`.
- Имя переменной всегда `ctx`.
- Не клади `ctx` в поля structs. Исключения редки (см. `http.Request.Context()`).
- Не передавай `nil`. Если «пока нечего передать» — используй `context.TODO()`.
- `Background()` — только на верхнем уровне (main, init, тесты).

## Цепочка

Каждый `WithX` возвращает **дочерний** контекст. Отмена родителя каскадно отменяет всех детей. Это механизм распространения отмены «сверху вниз» по всему стеку.

## Чего context НЕ делает

- Не останавливает выполняющийся код сам по себе. Программист обязан проверять `ctx.Done()` / `ctx.Err()` и уметь прерваться.
- Не заменяет обычные параметры. В `Value` кладут cross-cutting concerns, не бизнес-данные.
- Не «легковесен» к злоупотреблениям: `ctx.Value(key)` ищет ключ линейным проходом вверх по цепочке контекстов — чем длиннее цепочка `WithValue`, тем медленнее.

## Рекомендация

- Передавай `ctx` везде, где функция делает сетевые/БД-вызовы или может быть долгой.
- Делай `cancel()` в `defer` сразу после `WithCancel/WithTimeout`.
- На границе (HTTP handler, consumer, cron job) проверяй `ctx.Err()` и выходи аккуратно.
