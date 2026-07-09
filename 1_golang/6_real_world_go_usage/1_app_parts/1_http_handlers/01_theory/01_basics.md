# HTTP handlers: основы

Handler — это функция, превращающая HTTP-запрос в HTTP-ответ. В Go базовый интерфейс:

```go
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
```

`http.HandlerFunc` — обёртка для обычных функций.

## Обязанности handler'а

1. **Decode input**: query, path, body, headers.
2. **Validate**: формат, диапазоны, обязательные поля.
3. **Call service / use-case**: основная логика — НЕ в handler.
4. **Map result/error → HTTP**: статус, тело ответа.
5. **Encode output**: JSON, plain text, файл.

Handler — тонкий слой. Если он разрастается до 200 строк — что-то делает не своё.

## Анатомия

```go
func createUser(svc UserService) http.HandlerFunc {
    type req struct {
        Email string `json:"email"`
    }
    type resp struct {
        ID int `json:"id"`
    }
    return func(w http.ResponseWriter, r *http.Request) {
        var in req
        if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
            httpError(w, http.StatusBadRequest, "invalid body")
            return
        }
        u, err := svc.Create(r.Context(), in.Email)
        if err != nil {
            httpError(w, mapStatus(err), err.Error())
            return
        }
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(resp{ID: u.ID})
    }
}
```

Зависимости handler'а получают через замыкание (factory) или DI-конструктор.

## Routing

- Стандартный `http.ServeMux` (Go 1.22+ поддерживает методы и pattern'ы): `mux.HandleFunc("GET /users/{id}", h)`.
- Сторонние: `chi`, `gin`, `echo` — больше синтаксического сахара, удобный middleware-стек.

Для маленьких сервисов stdlib хватает.

## Status codes

- 2xx — успех.
- 4xx — ошибка клиента (400, 401, 403, 404, 409, 422, 429).
- 5xx — ошибка сервера (500, 502, 503, 504).

Не возвращай 200 при ошибке. Не возвращай 500 на «не нашли запись» — это 404.

## Контекст

`r.Context()` отменяется, если клиент отключился. Передавай его в сервисы и БД — пусть остановят работу.

## Минимум для чистого handler'а

- Не больше 30-40 строк.
- Никакой бизнес-логики (всё через service).
- Один success-путь, явные error-маппинги.
- Все таймауты — на уровне сервера/контекста, не в handler'е.
