# Части приложения — шпаргалка

Подробности и объяснения — в `01_theory/` каждой темы (ссылки в разделах). ⚠️ — частая ловушка.
Сторонние библиотеки (`chi`, `jwt`, `bcrypt`, `prometheus`, `otel`) только названы — код с ними не запускался.

## Слои: кто за что отвечает

| Слой | Делает | Не делает |
| :--- | :--- | :--- |
| Handler | decode → validate формат → вызов service → error → HTTP-статус → encode | бизнес-правила, SQL, транзакции |
| Middleware | cross-cutting: request_id, логи, recovery, метрики, трейсинг, timeout, CORS, auth, rate limit | бизнес-логику |
| Service | бизнес-правила, координация, границы транзакций, доменные ошибки | HTTP-статусы, SQL, глобальные синглтоны |
| Repository | SQL/драйвер, маппинг ошибок драйвера → доменные | бизнес-логику, side effects кроме записи |
| Config | один `Config` struct, загрузка + валидация в `main` | глобальная `var Cfg` |

Правило: «принимай интерфейсы, возвращай структуры»; `ctx context.Context` — первым параметром, `error` — последним результатом.

## 1. HTTP handlers

Подробно: [01_basics.md](./1_http_handlers/01_theory/01_basics.md), [02_pitfalls.md](./1_http_handlers/01_theory/02_pitfalls.md).

| Задача | Код |
| :--- | :--- |
| Интерфейс | `type Handler interface { ServeHTTP(w ResponseWriter, r *Request) }` |
| Функция как handler | `http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {...})` |
| Зависимости | factory-замыкание: `func createUser(svc UserService) http.HandlerFunc` |
| Роутинг с методом и параметром (1.22+) | `mux.HandleFunc("GET /users/{id}", h)`; в handler `r.PathValue("id")` |
| Декодировать JSON потоком | `if err := json.NewDecoder(r.Body).Decode(&in); err != nil { 400 }` |
| Лимит тела | `r.Body = http.MaxBytesReader(w, r.Body, 1<<20)` |
| Ответ JSON | `w.Header().Set("Content-Type", "application/json")` → `json.NewEncoder(w).Encode(resp)` |
| Отмена при отключении клиента | передавать `r.Context()` в service/БД |
| Таймауты сервера | `http.Server{ReadHeaderTimeout: 5s, ReadTimeout: 20s, WriteTimeout: 20s, IdleTimeout: 60s}` |

| Ситуация | Статус |
| :--- | :--- |
| Невалидное тело / параметры | 400 (422 — семантически неверные данные) |
| Нет / неверный токен | 401 |
| Аутентифицирован, но нельзя | 403 |
| Не нашли запись | 404, не 500 |
| Конфликт (дубликат) | 409 |
| Rate limit | 429 |
| Ошибка сервера | 500; апстрим: 502 / 503 / 504 |

## 2. Service layer

Подробно: [01_basics.md](./2_service_layer/01_theory/01_basics.md), [02_pitfalls.md](./2_service_layer/01_theory/02_pitfalls.md).

| Задача | Как |
| :--- | :--- |
| Форма | `type userService struct{ repo UserRepo; mailer Mailer; clock Clock }` + `NewUserService(...)` |
| Зависимости | узкие интерфейсы (`UserRepo`, `Mailer`, `Clock`), в тестах — fake |
| Бизнес-условия | sentinel / typed errors: `ErrInvalidEmail`, `ErrAlreadyExists` |
| Обёртка ошибок | `fmt.Errorf("save: %w", err)` |
| Транзакция | границы задаёт service: `s.tx.RunTx(ctx, func(ctx context.Context) error {...})` |
| Деление | по домену: `UserService`, `OrderService`, не один god-service |
| Цикл A → B → A | объединить / вынести общий пакет / события |

## 3. Repository layer

Подробно: [01_basics.md](./3_repository_layer/01_theory/01_basics.md), [02_pitfalls.md](./3_repository_layer/01_theory/02_pitfalls.md).

| Задача | Код |
| :--- | :--- |
| Одна строка | `row := r.db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)` → `row.Scan(&u.ID, &u.Email)` |
| Много строк | `rows, err := db.QueryContext(ctx, q, args...)`; `defer rows.Close()` |
| Запись | `db.ExecContext(ctx, q, args...)` |
| Не нашли | `if errors.Is(err, sql.ErrNoRows) { return User{}, ErrNotFound }` |
| Внутри и вне транзакции | метод принимает `DBTX` (общий интерфейс `*sql.DB` / `*sql.Tx`) или `TxRunner` кладёт tx в ctx |
| Транзакция | `ctx` с таймаутом → `tx, err := db.BeginTx(ctx, nil)` → `defer tx.Rollback()` → `tx.Commit()` |
| Миграции | `//go:embed migrations/*.sql` и применять на старте |

| Ошибка драйвера | Доменная |
| :--- | :--- |
| `sql.ErrNoRows` | `ErrNotFound` |
| pg `23505` unique violation | `ErrAlreadyExists` |
| pg `23503` FK violation | `ErrConflict` |

| Пагинация | Запрос | Когда |
| :--- | :--- | :--- |
| Offset | `LIMIT 20 OFFSET 40` | просто, медленнее с ростом offset |
| Cursor (keyset) | `WHERE id > $last ORDER BY id LIMIT 20` | большие списки, стабильно при вставках |

Инструменты: `database/sql` / `pgx` + `sqlx` или `sqlc` (codegen из SQL); тяжёлый ORM (`gorm`, `ent`) скрывает SQL.

## 4. Config and env

Подробно: [01_basics.md](./4_config_and_env/01_theory/01_basics.md), [02_pitfalls.md](./4_config_and_env/01_theory/02_pitfalls.md).

| Задача | Как |
| :--- | :--- |
| Приоритет источников | defaults → file → env → flags (каждый следующий переопределяет); секреты — secret manager |
| Структура | вложенные группы: `Config{HTTP HTTPConfig; DB DBConfig; Auth AuthConfig}` |
| Длительности | `time.Duration` + `"30s"`, `"100ms"`, не `int` секунд |
| Валидация | `cfg.Validate()` сразу после загрузки, до использования; ошибка → не стартовать |
| Маскировка секретов | свой `String()` (`Password:***`) или `slog.HandlerOptions.ReplaceAttr` |
| Hot reload | SIGHUP / `fsnotify` → новое значение в `atomic.Pointer[Config]` |
| Тесты | `t.Setenv("KEY", "v")` или конфиг через конструктор |
| Библиотеки | `flag` (stdlib), `kelseyhightower/envconfig`, `caarlos0/env`, `spf13/viper` |
| Шаблон env | `.env.example` в git, `.env` — в `.gitignore` |

## 5. Middleware chain

Подробно: [01_basics.md](./5_middleware_chain/01_theory/01_basics.md), [02_pitfalls.md](./5_middleware_chain/01_theory/02_pitfalls.md).

| Задача | Код |
| :--- | :--- |
| Тип | `func(next http.Handler) http.Handler` |
| Цепочка | `h := logging(recovery(auth(core)))` — выполняется снаружи внутрь, ответ — обратно |
| Значение в context | `type ctxKey int`; `ctx := context.WithValue(r.Context(), userKey, u)`; `next.ServeHTTP(w, r.WithContext(ctx))` |
| Прочитать в handler | `u, ok := r.Context().Value(userKey).(User)` |
| Перехватить статус | обёртка `statusWriter{http.ResponseWriter; status int}` с методом `WriteHeader` |
| Прочитать body и вернуть | `body, _ := io.ReadAll(r.Body)`; `r.Body = io.NopCloser(bytes.NewReader(body))` |
| Счётчик | `atomic.Int64` / Prometheus counter, не `counter++` |
| Группы маршрутов | `r := chi.NewRouter(); r.Use(recovery, logging, auth)` |
| Тест | `mw.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))` |

Типичный порядок снаружи внутрь: request_id → logging → recovery → metrics → tracing → timeout → CORS → auth → rate limit → handler.

```go
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("panic", "err", rec)
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		fmt.Println(r.Method, r.URL.Path, sw.status)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("id") == "0" {
			panic("boom")
		}
		fmt.Fprintf(w, "user %s", r.PathValue("id"))
	})
	h := logging(recovery(mux)) // снаружи внутрь: logging → recovery → mux

	for _, path := range []string{"/users/42", "/users/0"} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("GET", path, nil))
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("POST", "/users/42", nil))
}

// GET /users/42 200
// GET /users/0 500        (+ в stderr: ERROR panic err=boom)
// POST /users/42 405      — метод не совпал с паттерном
```

## 6. Auth and permissions

Подробно: [01_basics.md](./6_auth_and_permissions/01_theory/01_basics.md), [02_pitfalls.md](./6_auth_and_permissions/01_theory/02_pitfalls.md).

| Способ authn | Где уместен |
| :--- | :--- |
| Session cookie | web-фронт; нужна защита от CSRF (`SameSite`, CSRF-токен) |
| JWT Bearer | API, микросервисы; сервер не хранит сессию |
| OAuth 2.0 / OIDC | вход через внешнего провайдера |
| API key | service-to-service |
| mTLS | строгая межсервисная аутентификация |

| Модель authz | Суть |
| :--- | :--- |
| Roles | admin / editor / viewer — просто, негибко |
| Permissions | роль = набор строковых прав |
| RBAC + scopes | `scope: read:users write:orders` в токене |
| ABAC | политики по атрибутам (`owner == user.id`); casbin / OPA |

| Задача | Код |
| :--- | :--- |
| Хеш пароля | `bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)`; argon2id — рекомендация OWASP |
| Проверить пароль | `bcrypt.CompareHashAndPassword(hash, []byte(input))` |
| Сравнить токены | `subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1` |
| JWT: проверить алгоритм | в keyfunc: `if t.Method.Alg() != "HS256" { return nil, ErrUnexpectedAlg }` |
| Политика | одна функция `canEdit(u User, o Order) bool`, вызывается везде |
| Сроки токенов | access ~15 мин, refresh — дни, в HttpOnly cookie, с отзывом |
| Brute-force | rate limit per-IP и per-user на login, CAPTCHA, аудит неудачных входов |

## 7. Cache layer

Подробно: [01_basics.md](./7_cache_layer/01_theory/01_basics.md), [02_pitfalls.md](./7_cache_layer/01_theory/02_pitfalls.md).

| Паттерн | Как | Риск |
| :--- | :--- | :--- |
| Cache-aside (default) | get → miss → БД → set | stale до TTL / invalidate |
| Read-through / write-through | кеш сам читает / пишет в БД | нужна обёртка-библиотека |
| Write-back | пишем в кеш, в БД — асинхронно | потеря данных при падении |

| Инвалидация | Как |
| :--- | :--- |
| TTL | всегда; ограничивает stale-окно |
| Explicit | `delete` после записи в БД |
| Versioning | ключ `user:42:v3`, старые версии протухают сами |

| Задача | Инструмент |
| :--- | :--- |
| In-process | `map` + `sync.RWMutex`, `sync.Map`, `bigcache`, `freecache` |
| Ограничение размера | LRU `hashicorp/golang-lru/v2` + фоновая очистка просроченных |
| Distributed | Redis / Memcached (L1 in-process + L2 Redis) |
| Stampede | `singleflight`: `v, err, _ := group.Do(key, func() (any, error) { return load(key) })` |
| Метрики | hit_rate = hits / (hits + misses), latency Get, size, evictions/s |

## 8. Observability

Подробно: [01_basics.md](./8_observability/01_theory/01_basics.md), [02_pitfalls.md](./8_observability/01_theory/02_pitfalls.md).

| Сигнал | Отвечает на | Инструмент |
| :--- | :--- | :--- |
| Logs | что случилось | `log/slog` (1.21+), JSON в stdout |
| Metrics | сколько / как часто / как быстро | `prometheus/client_golang`, endpoint `/metrics` |
| Traces | где запрос провёл время | OpenTelemetry: `ctx, span := tracer.Start(ctx, "svc.Get")`; `defer span.End()` |

| Тип метрики | Пример |
| :--- | :--- |
| Counter | `http_requests_total` — только растёт |
| Gauge | активные соединения, память |
| Histogram | latency, размер запроса; задавать свои `Buckets` |
| Summary | предрассчитанные квантили |

| Задача | Код |
| :--- | :--- |
| Структурный лог | `slog.Info("user registered", "user_id", u.ID)` |
| Ошибка | `slog.Error("save failed", "err", err, "user_id", id)` |
| JSON-логгер | `slog.New(slog.NewJSONHandler(os.Stdout, nil))` |
| Алерт на долю 5xx | `rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m]) > 0.05` |
| Liveness | `/healthz` — простой 200, без проверки зависимостей |
| Readiness | `/readyz` — БД доступна, миграции применены |

## Ловушки

- ⚠️ Бизнес-логика в handler или в SQL → service пустеет, правила расползаются.
- ⚠️ Ошибки `Decode`/`Encode` проигнорированы → нулевые поля / битый JSON.
- ⚠️ `WriteHeader` после первого `Write` или после `next.ServeHTTP` — игнорируется (warning в логе).
- ⚠️ `http.Error(w, err.Error(), 500)` наружу — утечка имён таблиц, путей, версий.
- ⚠️ `http.Server` без таймаутов → slowloris и висящие соединения.
- ⚠️ net/http сам ловит panic handler'а (логирует, рвёт соединение), но без своего recovery клиент не получит 500, а лог — контекста.
- ⚠️ `recovery` внутри `auth` → паника в auth не поймана; `authorize` снаружи `auth` → user ещё не в context.
- ⚠️ `string`-ключ в `context.WithValue` → коллизии между пакетами; нужен свой тип.
- ⚠️ Middleware прочитал `r.Body` и не вернул → handler видит пустое тело.
- ⚠️ `sql.ErrNoRows` утёк из repo → service зависит от `database/sql`.
- ⚠️ `SELECT *`, конкатенация SQL (инъекция), `Query` без `Context`, забытый `rows.Close()`, N+1 в цикле.
- ⚠️ Транзакция без таймаута и без `defer tx.Rollback()` → висящие row-locks.
- ⚠️ Глобальная `var Cfg` + reload из SIGHUP → data race; нужен `atomic.Pointer[Config]`.
- ⚠️ DSN с паролем в `log.Printf("%+v", cfg.DB)`.
- ⚠️ Пароли в SHA-256/MD5; JWT без проверки `alg` (`none`); секреты в payload JWT (он подписан, не зашифрован).
- ⚠️ Проверка прав только во фронте; `token == expected` (timing attack).
- ⚠️ Кеш без TTL и без лимита размера → утечка памяти / вечные ключи в Redis.
- ⚠️ Ошибка Redis роняет запрос → кеш должен деградировать к источнику.
- ⚠️ Полный flush кеша → stampede на БД.
- ⚠️ Label `user_id` / `request_id` в метрике → миллионы time series.
- ⚠️ Метрика создаётся в handler на каждый запрос → регистрируйте один раз глобально.
- ⚠️ `go func() { fetch(context.Background()) }()` → trace разорван.
- ⚠️ Тяжёлый liveness-probe с проверкой БД → ложные рестарты pod'а.
