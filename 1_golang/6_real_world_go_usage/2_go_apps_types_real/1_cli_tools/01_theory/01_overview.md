# CLI tools на Go

Go хорошо подходит для CLI: один статический бинарник, быстрый старт, кросс-компиляция.

## Базовая структура

Маленький CLI:

```
cli/
├── go.mod
├── main.go
└── cmd.go
```

Большой CLI с подкомандами:

```
cli/
├── go.mod
├── main.go
├── cmd/
│   ├── root.go
│   ├── build.go
│   ├── deploy.go
│   └── version.go
├── internal/
│   └── ...
```

## Парсинг флагов

- **flag** (stdlib) — простой, без подкоманд.
- **cobra** — стандарт de-facto, поддержка subcommands, генерация completion, help.
- **urfave/cli** — компактнее cobra, тоже популярна.

```go
// cobra
var rootCmd = &cobra.Command{Use: "tool"}
var buildCmd = &cobra.Command{
    Use: "build",
    RunE: func(cmd *cobra.Command, args []string) error {
        return runBuild(cmd.Context())
    },
}

func init() { rootCmd.AddCommand(buildCmd) }

func main() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

## Конфиг

Приоритет (от низкого к высокому): defaults → file → env → flags. См. `1_app_parts/4_config_and_env`.

Стандарт: `--config /path/to/file` + переменные окружения с префиксом твоей утилиты.

## I/O и exit codes

- Данные — в stdout.
- Ошибки и логи — в stderr.
- Exit code 0 — успех, !=0 — ошибка.

```go
if err := run(); err != nil {
    fmt.Fprintln(os.Stderr, "error:", err)
    os.Exit(1)
}
```

UNIX-pipe'ы зависят от этого: `tool list | grep foo` работает только если данные в stdout.

## Прогресс и интерактив

- `progressbar/v3` — прогресс-бары.
- `survey/v2` — интерактивные prompts (выбор, текст, password).
- `pterm`, `lipgloss` — стилизация в TUI.

Интерактив помечай отдельно: `tool deploy --interactive`. Скрипты должны работать без TTY.

## Кросс-компиляция

```bash
GOOS=linux   GOARCH=amd64 go build -o tool-linux-amd64
GOOS=darwin  GOARCH=arm64 go build -o tool-darwin-arm64
GOOS=windows GOARCH=amd64 go build -o tool-windows-amd64.exe
```

Один codebase → бинарники под все платформы. Используй `goreleaser` для автоматизации релизов.

## Embedding ресурсов

Шаблоны, дефолтные конфиги — внутрь бинарника:

```go
import _ "embed"

//go:embed templates/default.yaml
var defaultConfig []byte
```

Не нужно тащить файлы рядом с бинарём.

## Тестирование

- Бизнес-логику — обычные unit-тесты.
- Команды cobra — `cobra.Command.SetArgs(...)` + `Execute()`.
- E2E — компилировать бинарник, запускать `os/exec`.

## Правило

- Один статический бинарник без внешних зависимостей.
- Stdin/stdout/stderr — UNIX-friendly.
- Конфиг с приоритетом источников.
- Кросс-компиляция через CI.
