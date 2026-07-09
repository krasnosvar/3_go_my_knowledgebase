# Use-cases ошибок

## 1. Валидация входа

Вернуть `*ValidationError` с полем и причиной. Сервис передаёт наверх, HTTP-слой отдаёт 400 с деталями.

```go
if name == "" {
    return &ValidationError{Field: "name", Message: "required"}
}
```

## 2. Внешние API (HTTP/gRPC)

Заверни network-ошибку вместе с endpoint-ом и кодом статуса:

```go
if resp.StatusCode >= 500 {
    return fmt.Errorf("calling %s: status %d: %w", url, resp.StatusCode, errRemote)
}
```

Это даёт контекст и для ретраев (`Is(err, errRemote)`), и для алертов.

## 3. Файловая система и БД

Мапим типовые ошибки нижнего уровня в свои sentinel-ошибки на границе репозитория:

```go
if errors.Is(err, sql.ErrNoRows) {
    return nil, ErrNotFound
}
```

Выше по стеку никто не должен знать про `database/sql`.

## 4. Retry policy

Решение «ретраить или нет» должно быть явным. Частый приём — интерфейс с методом `Temporary() bool` или проверка `errors.Is(err, ErrTransient)`:

```go
if isRetriable(err) && attempt < 3 {
    time.Sleep(backoff(attempt))
    continue
}
```

## 5. Контекст отмены

`context.Canceled` и `context.DeadlineExceeded` — это нормальные завершения, не «ошибки для логов»:

```go
if errors.Is(err, context.Canceled) {
    return nil // это ок
}
```

Иначе обычный `log.Error` забьёт лог при каждом графовом shutdown.

## 6. Пользовательские сообщения

Формируются на границе. В ошибке — технический контекст, в ответе пользователю — человеческий текст:

```go
// внутри
return fmt.Errorf("validate email: invalid format: %w", err)
// на транспорте
return 400, "некорректный email"
```
