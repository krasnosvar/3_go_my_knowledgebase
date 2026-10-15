# net/http, net — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Паттерны `ServeMux` с методом и `{id}` — Go 1.22+.

## HTTP-сервер

Подробно: [01_http_server.md](./01_theory/01_http_server.md).

| Задача | Код |
| :--- | :--- |
| Роутер | `mux := http.NewServeMux()` |
| Обработчик-функция | `mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {...})` |
| Метод + параметр пути (1.22+) | `mux.HandleFunc("GET /users/{id}", h)` → `r.PathValue("id")`; другой метод → `405` |
| Интерфейс обработчика | `ServeHTTP(w http.ResponseWriter, r *http.Request)`; функцию → `http.HandlerFunc(f)` |
| Сервер с таймаутами | `srv := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}` |
| Запуск | `log.Fatal(srv.ListenAndServe())` |
| Заголовок → статус → тело | `w.Header().Set("Content-Type", "application/json")`; `w.WriteHeader(http.StatusCreated)`; `json.NewEncoder(w).Encode(v)` |
| Контекст запроса | `ctx := r.Context()` — отменяется, когда клиент ушёл; передавать в БД/внешние вызовы |
| Middleware | `func logging(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { start := time.Now(); next.ServeHTTP(w, r); log.Println(r.URL.Path, time.Since(start)) }) }` |
| Цепочка | `srv.Handler = auth(recovery(logging(mux)))` |
| Остановка | `srv.Shutdown(ctx)` — не принимает новые, ждёт текущие |

| Таймаут `http.Server` | Что ограничивает | Ориентир из теории |
| :--- | :--- | :--- |
| `ReadHeaderTimeout` | чтение заголовков (slowloris) | `5s` |
| `ReadTimeout` | чтение всего запроса | `20s` |
| `WriteTimeout` | запись ответа | `20s` |
| `IdleTimeout` | keep-alive между запросами | `60s` |

```go
ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
defer stop()
go func() { _ = srv.ListenAndServe() }()
<-ctx.Done()
sdCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
_ = srv.Shutdown(sdCtx)
```

## HTTP-клиент

Подробно: [02_http_client.md](./01_theory/02_http_client.md).

| Задача | Код |
| :--- | :--- |
| Клиент (один на приложение) | `client := &http.Client{Timeout: 10 * time.Second}` — общий лимит: dial + заголовки + body |
| Запрос с контекстом | `req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))` |
| Заголовки | `req.Header.Set("Authorization", "Bearer "+token)` |
| Выполнить | `resp, err := client.Do(req)`; затем `defer resp.Body.Close()` |
| Проверить статус | `if resp.StatusCode < 200 \|\| resp.StatusCode >= 300 { ... }` |
| JSON-ответ | `json.NewDecoder(resp.Body).Decode(&v)` |
| Пул соединений | `&http.Transport{MaxIdleConns: 100, MaxConnsPerHost: 20, MaxIdleConnsPerHost: 20, IdleConnTimeout: 90 * time.Second}` |
| Ретраи | в stdlib нет; только идемпотентные `GET`/`PUT`/`DELETE` |

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) { // Go 1.22+
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"id": r.PathValue("id")})
	})
	srv := httptest.NewServer(mux) // 127.0.0.1:случайный порт
	defer srv.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for _, m := range []string{http.MethodGet, http.MethodPost} {
		req, _ := http.NewRequestWithContext(ctx, m, srv.URL+"/users/42", nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s %d %q\n", m, resp.StatusCode, body)
	}
	// GET 200 "{\"id\":\"42\"}\n"
	// POST 405 "Method Not Allowed\n"
}
```

## net: TCP/UDP

Подробно: [03_net_basics.md](./01_theory/03_net_basics.md).

| Задача | Код |
| :--- | :--- |
| TCP-клиент | `conn, err := net.Dial("tcp", "example.com:80")` |
| С таймаутом и контекстом | `d := net.Dialer{Timeout: 5 * time.Second}; d.DialContext(ctx, "tcp", addr)` |
| TCP-сервер | `ln, err := net.Listen("tcp", ":8080")`; `conn, err := ln.Accept()`; `go handle(conn)` |
| Дедлайны | `conn.SetReadDeadline(time.Now().Add(2 * time.Second))`, `SetWriteDeadline` (абсолютное время) |
| DNS | `net.ResolveTCPAddr("tcp", "example.com:80")`; свой — `net.Resolver` |
| UDP | `pc, _ := net.ListenPacket("udp", ":1053")`; `n, addr, err := pc.ReadFrom(buf)`; `pc.WriteTo(resp, addr)` |
| Unix socket | `net.Listen("unix", "/tmp/app.sock")` |

| Нужно | Уровень |
| :--- | :--- |
| REST/HTTP API | `net/http` |
| БД, gRPC | `database/sql`, `grpc` |
| Свой бинарный протокол, прокси/туннель, load-тестер | `net` |

## Ловушки

- ⚠️ `http.DefaultClient` и `http.Get` без таймаута — запрос может висеть вечно.
- ⚠️ Не закрыт `resp.Body` — соединения не возвращаются в пул.
- ⚠️ `err == nil` после `Do` не значит 200 — проверяйте `StatusCode`.
- ⚠️ Новый `http.Client`/`Transport` на каждый запрос — пул соединений не работает.
- ⚠️ Сервер без таймаутов — slowloris и утечки соединений.
- ⚠️ `WriteHeader` после первого `Write` не действует — заголовки уже ушли.
- ⚠️ `panic` в handler: `net/http` гасит её для соединения, но состояние может быть битым — нужен recovery middleware.
- ⚠️ Ретрай `POST` без идемпотентности — двойные операции.
- ⚠️ `conn.Read` без дедлайна — зависание на «молчащем» соединении.
- ⚠️ `Read` может вернуть `n > 0` и `err` одновременно — сначала обработайте `n` байт.
