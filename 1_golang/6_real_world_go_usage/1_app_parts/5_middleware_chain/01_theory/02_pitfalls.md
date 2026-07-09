# Middleware: pitfalls

## 1. Неправильный порядок

```go
h := auth(recovery(logging(handler)))
```

`auth` снаружи `recovery` → если auth паникует, recovery его не поймает. Recovery **первым (снаружи)** или хотя бы до того, что может упасть.

Логирование тоже близко к внешнему слою — иначе пропустишь логи запросов, упавших в auth.

## 2. Middleware с побочкой между запросами

```go
var counter int
func metrics(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        counter++ // race
        next.ServeHTTP(w, r)
    })
}
```

`counter++` без atomic — data race. Используй `sync/atomic` или Prometheus counter.

## 3. WriteHeader в middleware после next

```go
next.ServeHTTP(w, r)
w.WriteHeader(http.StatusOK) // заголовки уже отправлены, ничего не сделает
```

После `next.ServeHTTP` ответ уже формируется/отправлен. Изменения headers в middleware — **до** `next`, не после.

Если нужно перехватывать статус для логов — оборачивай `ResponseWriter`:

```go
type statusWriter struct{ http.ResponseWriter; status int }
func (sw *statusWriter) WriteHeader(code int) { sw.status = code; sw.ResponseWriter.WriteHeader(code) }
```

## 4. Утечка значений в context

```go
ctx := context.WithValue(r.Context(), "user", u)
```

`string` ключ — коллизии между пакетами. Используй свой тип:

```go
type ctxKey struct{ name string }
var userKey = ctxKey{"user"}
```

## 5. Middleware блокирует запрос

```go
func slowLog(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // синхронно пишем в БД лог
        db.Insert(...)
        next.ServeHTTP(w, r)
    })
}
```

I/O в middleware → каждый запрос ждёт. Логи пиши асинхронно (buffered channel, отдельная горутина).

## 6. Middleware видит body после handler

```go
body, _ := io.ReadAll(r.Body) // забрали body
next.ServeHTTP(w, r)           // handler видит пустое body
```

Если нужно прочитать body в middleware (логирование, валидация) — заверни обратно:

```go
body, _ := io.ReadAll(r.Body)
r.Body = io.NopCloser(bytes.NewReader(body))
next.ServeHTTP(w, r)
```

И учти лимит размера, чтобы не съесть память.

## 7. CORS-заголовки в каждом handler

```go
func handler(w, r) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // ...
}
```

CORS — middleware, не дело каждого handler'а. Используй `gorilla/handlers.CORS` или ручной.

## 8. Авторизация до auth

```go
h := authorize(auth(handler))
```

`authorize` снаружи `auth` → user ещё не положен в context. Либо переставь, либо объедини в один middleware.

## 9. Цепочка собирается ad-hoc

```go
http.Handle("/users", logging(auth(usersHandler)))
http.Handle("/orders", recovery(logging(auth(ordersHandler)))) // другой порядок
```

Несогласованные цепочки → разное поведение для разных endpoint'ов. Используй роутер с поддержкой group'ов (`chi`, `gin`):

```go
r := chi.NewRouter()
r.Use(recovery, logging, auth)
r.Get("/users", usersHandler)
```

## 10. Middleware без тестов

Каждый middleware — самостоятельная единица. Тестируй:

```go
called := false
mw := logging(http.HandlerFunc(func(w, r) { called = true }))
mw.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
assert.True(t, called)
```

## Правило

- Recovery — внешний слой.
- Cross-cutting — middleware.
- Один middleware — одна задача.
- Цепочка определяется централизованно.
