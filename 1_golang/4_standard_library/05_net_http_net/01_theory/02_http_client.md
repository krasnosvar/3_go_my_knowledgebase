# net/http: client

## Не использовать http.DefaultClient

`http.DefaultClient` не имеет таймаута. Запрос может висеть вечно. Всегда свой клиент.

## Минимальный клиент

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}

resp, err := client.Get("https://example.com")
if err != nil {
    return err
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
```

`Timeout` у `http.Client` — общий на весь запрос (dial + headers + body). Это грубый верхний лимит.

## Полный Request

Для POST, PUT, headers — используй `http.NewRequestWithContext`:

```go
req, err := http.NewRequestWithContext(ctx, http.MethodPost, url,
    bytes.NewReader(body))
if err != nil { return err }
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer "+token)

resp, err := client.Do(req)
if err != nil { return err }
defer resp.Body.Close()
```

`ctx` из сервиса передаётся через `NewRequestWithContext` — если клиент ушёл, исходящий запрос отменяется.

## Проверка статуса

```go
if resp.StatusCode < 200 || resp.StatusCode >= 300 {
    b, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, b)
}
```

Важно: `err == nil` после `Do` не значит «200 OK». Это значит «HTTP-ответ получен, смотри статус».

## Transport

`http.Client` использует `http.Transport` — пул соединений, TLS, proxy. Настраивай для prod:

```go
tr := &http.Transport{
    MaxIdleConns:        100,
    MaxConnsPerHost:     20,
    MaxIdleConnsPerHost: 20,
    IdleConnTimeout:     90 * time.Second,
}
client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
```

Общий клиент на приложение, а не новый на каждый запрос — чтобы работал connection pool.

## Ретраи

В stdlib нет. Пиши руками или используй библиотеки (`hashicorp/go-retryablehttp`). Ретраить **только идемпотентные** запросы (`GET`, `PUT`, `DELETE`) — не `POST` без проверки.

## Основные ошибки

- Не закрыть `resp.Body` — leak connections, пул забьётся.
- `http.Get` без таймаута — зависший запрос.
- Перезаписывать `req.Context()` вручную — используй `NewRequestWithContext`.
- Парсить `resp.Body` через `json.Unmarshal(io.ReadAll(...))` — нормально; на потоковых — `json.NewDecoder(resp.Body).Decode(&v)`.
- Не проверять `resp.StatusCode` — работает «всё зелёное», а на самом деле 500.

## Правило

- Один `http.Client` в сервисе.
- `NewRequestWithContext` — всегда.
- Таймаут — обязателен.
- `resp.Body.Close()` — в `defer`.
- Проверка `StatusCode` — явно.
