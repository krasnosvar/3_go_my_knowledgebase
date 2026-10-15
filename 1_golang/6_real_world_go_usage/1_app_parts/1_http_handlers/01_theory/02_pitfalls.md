# HTTP handlers: pitfalls

## 1. Бизнес-логика в handler

```go
func createOrder(w http.ResponseWriter, r *http.Request) {
    // декодим, валидируем, считаем скидки, проверяем склад, пишем в БД
    // 250 строк
}
```

Handler становится unmaintainable. Бизнес-логика → service. Handler декодирует, вызывает, кодирует.

## 2. Игнорирование ошибок Encode/Decode

```go
json.NewDecoder(r.Body).Decode(&in) // err игнорируется
json.NewEncoder(w).Encode(resp)     // тоже
```

Decode без проверки → silently nil поля. Encode без проверки — отдадим клиенту сломанный JSON.

## 3. Не закрывать body клиента

В прод HTTP-сервере Go закрывает `r.Body` сам. Но если ты `io.Copy` или `io.ReadAll` без проверки — можешь съесть всю память.

```go
r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MiB лимит
```

## 4. WriteHeader дважды

```go
w.WriteHeader(http.StatusOK)
// ... что-то пошло не так ...
w.WriteHeader(http.StatusInternalServerError) // ignored, в логе warning
```

После первого `Write` или `WriteHeader` заголовки уже отправлены. Решай статус заранее.

## 5. Не выставлены таймауты на сервере

`http.Server` без `ReadTimeout`/`WriteTimeout`/`IdleTimeout` — уязвим к slowloris, висящим коннектам.

```go
srv := &http.Server{
    ReadHeaderTimeout: 5 * time.Second,
    ReadTimeout:       20 * time.Second,
    WriteTimeout:      20 * time.Second,
    IdleTimeout:       60 * time.Second,
}
```

## 6. Возврат внутренних ошибок клиенту

```go
http.Error(w, err.Error(), 500)
```

Раскрывает имена БД-таблиц, путей файлов, IP, версии библиотек. На прод — generic-сообщение, детали — в логи.

## 7. JSON в hot path без потокового decoder

`io.ReadAll` + `json.Unmarshal` на больших body — двойная аллокация. Лучше:

```go
err := json.NewDecoder(r.Body).Decode(&in)
```

Стримит без хранения всего тела в памяти.

## 8. panic в handler

`panic` в handler не роняет процесс: `http.Server` сам ловит её для этого запроса, пишет stack trace в error log и обрывает соединение (HTTP/2 — `RST_STREAM`). Клиент при этом не получает 500, а в логе нет контекста запроса. Явный `recovery` middleware — must:

```go
func recovery(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if rec := recover(); rec != nil {
                slog.Error("panic", "err", rec)
                http.Error(w, "internal error", 500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

## 9. Игнорировать r.Context()

Запрос отменён клиентом → handler продолжает делать тяжёлую работу. Передавай ctx в сервисы и БД, проверяй `ctx.Err()` в долгих операциях.

## 10. CORS, CSRF, security headers — забыты

Добавляются middleware'ом, но часто забываются. На production-сервисе нужны: CORS (если фронт на другом домене), CSP, HSTS, X-Frame-Options и т.п. Используй `chi/middleware`, `gorilla/handlers` или ручной middleware.

## Правило

Handler — короткий, тонкий, явный. Каждая ошибка обработана, каждый таймаут на месте.
