# Configuration: pitfalls

## 1. Глобальная переменная Config

```go
var Cfg Config
func init() { Cfg = loadConfig() }
```

Импорты получают конфиг через global. Тесты страдают (нельзя подменить), зависимости неявные. Передавай Config через конструкторы.

## 2. Configuration drift

«Дев работает, прод нет» — в 80% случаев из-за разных конфигов. Решение:

- Все варианты env описаны в `.env.example` или `config.example.yaml`.
- В CI на старте — проверяет, что все обязательные переменные определены.
- Конфиг-различия документированы.

## 3. Секреты в логах

```go
log.Printf("connecting to %+v", cfg.DB) // DSN с паролем в лог
```

Реализуй `String()` который маскирует секреты:

```go
func (c DBConfig) String() string {
    return fmt.Sprintf("DBConfig{Host:%s, User:%s, Password:***}", c.Host, c.User)
}
```

Или используй `slog.Replacer` чтобы автоматически замаскировать поля типа `password`.

## 4. Конфиг используется до валидации

```go
cfg := loadConfig()
db := openDB(cfg.DB.DSN) // пустой DSN → cryptic error
cfg.Validate()           // поздно
```

Валидация — первой. Только после неё — использование.

## 5. Бесконечные дефолты

```go
Port: 8080
Timeout: 30s
Workers: 10
MaxRetries: 3
RateLimit: 1000
```

Если каждое значение жёстко зашито в дефолт — продакшн настройки могут не учитываться, и ты не заметишь. Пометь явно: «if 0, then ...» — или требуй явное значение.

## 6. Несовпадение имён в коде и env

```go
type Cfg struct {
    HTTPAddr string `env:"HTTP_BIND_ADDR"`
}
```

В коде `HTTPAddr`, в env `HTTP_BIND_ADDR`. Кто-то ставит `HTTP_ADDR` и удивляется. Используй consistent naming, документируй.

## 7. Hot reload с глобалами

`Cfg = loadConfig()` в SIGHUP-обработчике. Параллельные горутины читают `Cfg` без синхронизации → race. Используй `atomic.Pointer[Config]`.

## 8. Время как config

```go
type Config struct {
    SomeTimeout int `env:"TIMEOUT_SEC"` // секунды как int
}
```

Используй `time.Duration`:

```go
SomeTimeout time.Duration `env:"TIMEOUT" default:"30s"`
```

`5m`, `100ms` — естественно. Голые секунды — частый источник «10000 — это что, ms или ns?».

## 9. Конфиг не в одной struct

```go
var Port = os.Getenv("PORT")
var DBHost = os.Getenv("DB_HOST")
```

Переменные раскиданы по пакетам, не видно полной картины. Один Config — единая точка правды.

## 10. Тесты с реальным env

`os.Getenv("DB_DSN")` в тесте → зависит от окружения CI. Используй `t.Setenv` или передавай конфиг в конструктор.

## Правило

- Один источник правды (Config struct).
- Валидация до использования.
- Секреты — отдельно, маскированы в логах.
- Тесты — с подменяемым конфигом, не глобальным.
