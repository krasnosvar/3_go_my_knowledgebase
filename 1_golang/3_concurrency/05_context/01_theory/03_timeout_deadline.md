# WithTimeout и WithDeadline

Две функции для ограничения времени жизни контекста.

## WithTimeout

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()
```

Контекст отменится либо когда родитель отменится, либо через 2 секунды.

## WithDeadline

```go
ctx, cancel := context.WithDeadline(parent, time.Now().Add(2*time.Second))
defer cancel()
```

Разница только в удобстве. `WithTimeout(parent, d)` ≡ `WithDeadline(parent, time.Now().Add(d))`.

## Наследование дедлайна

Если в цепочке уже есть дедлайн, новый не может его «расширить»:

```go
// у parent дедлайн через 100 мс
ctx, cancel := context.WithTimeout(parent, 10*time.Second)
```

`ctx.Deadline()` вернёт время от 100 мс (то, что раньше). Дочерний контекст не может жить дольше родителя.

## Cascading timeouts

Типовая практика: сервис задаёт общий таймаут, каждый слой сокращает его под себя.

```go
// handler: total 2s
ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
defer cancel()

// service: db.Query получает оставшееся время
res, err := db.QueryContext(ctx, ...)
```

Без бюджета «нарезки» каждый слой может съесть всё время.

## Получение остатка времени

```go
if dl, ok := ctx.Deadline(); ok {
    remaining := time.Until(dl)
    // ...
}
```

Полезно для rate limit или решения «есть ли смысл стартовать эту операцию вообще».

## WithTimeout без cancel

```go
ctx, _ := context.WithTimeout(parent, time.Second)
```

Компилятор не запретит, но это баг: таймер не отменяется сразу после завершения операции, ресурсы висят до срабатывания таймаута. В современных `go vet` это ловят, но лучше всегда `defer cancel()`.

## Разница Canceled / DeadlineExceeded

```go
if errors.Is(err, context.DeadlineExceeded) {
    return http.StatusGatewayTimeout
}
```

- `Canceled` — кто-то явно отменил (юзер закрыл вкладку, сервис выключается).
- `DeadlineExceeded` — истёк таймаут. Сигнал о проблеме в нижней системе или в твоём бюджете.

## Рекомендация

- Таймаут на внешний вызов — **обязателен**. Всегда.
- Таймаут пропусти через контекст, не через `time.After` поверх.
- Один запрос — один общий `ctx` с дедлайном, нарезка — дочерние WithTimeout.
