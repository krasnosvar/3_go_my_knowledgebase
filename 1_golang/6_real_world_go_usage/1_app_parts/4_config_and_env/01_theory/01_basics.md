# Configuration

Конфиг — то, что отличает прод от стейджа от лока. Хранилища, креды, лимиты, фичи. Правильный конфиг сильно облегчает жизнь.

## Источники

В порядке приоритета (более поздний переопределяет более ранний):

1. **Defaults** в коде — sane fallback'ы.
2. **File** — `.yaml` / `.toml` / `.json`. Удобно для прод-настроек.
3. **Env** — `APP_PORT=8080`. Стандарт для контейнеров.
4. **Flags** — `--port=8080`. Локальная отладка, операторские утилиты.
5. **Secret manager** — Vault, AWS Secrets Manager. Только для секретов.

Каждый следующий уровень переопределяет предыдущий.

## Структура

```go
type Config struct {
    HTTP HTTPConfig
    DB   DBConfig
    Auth AuthConfig
}

type HTTPConfig struct {
    Addr            string        `env:"HTTP_ADDR" default:":8080"`
    ReadTimeout     time.Duration `env:"HTTP_READ_TIMEOUT" default:"5s"`
    WriteTimeout    time.Duration `env:"HTTP_WRITE_TIMEOUT" default:"10s"`
}

type DBConfig struct {
    DSN          string `env:"DB_DSN" required:"true"`
    MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" default:"20"`
}
```

Группируй родственные настройки. Один большой плоский конфиг — нечитаемо.

## Библиотеки

- `kelseyhightower/envconfig` — простой, через теги.
- `spf13/viper` — мульти-источники, мощный, тяжёлый.
- `caarlos0/env` — лёгкий конкурент envconfig.
- `flag` (stdlib) — для CLI-флагов.

В микросервисе на 5 настроек stdlib хватает. В монолите — что-то типа envconfig.

## Валидация

```go
func (c *Config) Validate() error {
    if c.HTTP.Addr == "" {
        return errors.New("HTTP_ADDR required")
    }
    if c.DB.MaxOpenConns < 1 {
        return errors.New("DB_MAX_OPEN_CONNS must be > 0")
    }
    return nil
}
```

Загрузил → валидируй сразу. Стартуй только если всё ОК. Лучше упасть на старте, чем через час под нагрузкой.

## Секреты

- НЕ в код / git: `.env` в `.gitignore`, шаблон `.env.example` коммитим.
- Не в логи: `Config.String()` маскирует поля типа `password`, `token`.
- В контейнерах — env через секрет-менеджер, не через `-e PASSWORD=...`.

## 12-factor

Идея: конфиг — это окружение, не код. Один и тот же бинарник работает в dev/staging/prod, отличается только env. Это база для контейнеризации и CI/CD.

## Hot reload

Иногда нужно: фича-флаги, лимиты. Реализуется:

- `SIGHUP` → reread file.
- Watcher на файл (`fsnotify`).
- Внешний config-server (Consul, etcd).

Сложно правильно: атомарно подменить, рассказать слушателям, обновить активные операции. Чаще проще: задеплоить новую версию.

## Правило

- Defaults в коде, env переопределяет, flags — для отладки.
- Валидация на старте.
- Секреты — никогда в git.
- Один Config struct, явная инициализация в main.
