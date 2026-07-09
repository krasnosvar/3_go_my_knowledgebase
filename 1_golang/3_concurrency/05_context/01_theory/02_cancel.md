# context.WithCancel

`WithCancel` создаёт дочерний контекст, который можно явно отменить, вызвав `cancel()`.

## Синтаксис

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```

- `ctx` — дочерний контекст.
- `cancel` — функция отмены. Всегда вызывай её (обычно через `defer`), даже если операция завершилась успешно.

## Зачем `cancel()` обязательно

Без `cancel()` рантайм хранит ссылку на дочерний контекст до момента отмены родителя. Это потенциальная утечка памяти и утечка горутин, которые подписаны на `ctx.Done()`.

## Типовой сценарий

```go
func search(ctx context.Context, queries []string) ([]Result, error) {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    results := make(chan Result, len(queries))
    for _, q := range queries {
        q := q
        go func() {
            r, err := fetch(ctx, q)
            if err == nil {
                results <- r
            }
        }()
    }

    select {
    case r := <-results:
        cancel() // отменяем остальные горутины
        return []Result{r}, nil
    case <-ctx.Done():
        return nil, ctx.Err()
    }
}
```

«Первый вернул результат — остальных отменяем».

## Distiguish Canceled vs DeadlineExceeded

```go
select {
case <-ctx.Done():
    switch ctx.Err() {
    case context.Canceled:
        return errUserAbort
    case context.DeadlineExceeded:
        return errTimeout
    }
}
```

Полезно для метрик и логов.

## Cancel может быть вызван много раз

Безопасно. Следующие вызовы — no-op. Это важно: не бойся писать `defer cancel()`, даже если внутри уже могут вызывать.

## Распространение

Когда родитель отменяется — все дочерние контексты тоже. Это происходит автоматически.

```go
ctx1, cancel1 := context.WithCancel(root)
ctx2, cancel2 := context.WithCancel(ctx1)
defer cancel1()
defer cancel2()

cancel1() // ctx2 тоже закрылся
```

## Антипаттерны

- Не вызывать `cancel()` в надежде «оно само завершится» — это утечка.
- Игнорировать `cancel()` в отложенной горутине — отмены не случится.
- Делать `cancel()` из другого места, а не defer: теряешь гарантию.
