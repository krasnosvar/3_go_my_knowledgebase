# Timeout в select

Самый частый use-case `select` — ограничить время ожидания.

## time.After

```go
select {
case v := <-ch:
    use(v)
case <-time.After(2 * time.Second):
    return ErrTimeout
}
```

`time.After(d)` возвращает канал, который «сработает» через `d` и передаст значение `time.Time`. Мы это значение не используем — достаточно получить что-то из канала.

### Нюанс: утечка таймера

`time.After` создаёт таймер, который уничтожается только после срабатывания (или после GC). Если `select` сработал по другой ветке, таймер всё ещё жив.

В hot path предпочти `time.NewTimer` + `Stop`:

```go
t := time.NewTimer(2 * time.Second)
defer t.Stop()

select {
case v := <-ch:
    use(v)
case <-t.C:
    return ErrTimeout
}
```

## context.WithTimeout

Идиоматичнее и удобнее, если время ожидания — часть общего дедлайна операции:

```go
ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
defer cancel()

select {
case v := <-ch:
    use(v)
case <-ctx.Done():
    return ctx.Err()
}
```

`ctx.Err()` — либо `context.DeadlineExceeded`, либо `context.Canceled`. Это полезно различать.

## Deadline

`context.WithDeadline` — если известна точка времени, а не длительность. Чаще удобен timeout.

## Несколько таймаутов одновременно

Можно добавить несколько веток с разными таймаутами:

```go
select {
case v := <-ch:
    use(v)
case <-time.After(soft):
    degradeToFallback()
case <-time.After(hard):
    return ErrTimeout
}
```

На практике вернее — через вложенные контексты.

## Общая рекомендация

- Разовый таймаут, локально — `time.After` / `time.NewTimer`.
- Таймауты, которые должны каскадироваться (HTTP handler → сервис → репо → БД) — `context.WithTimeout`.
