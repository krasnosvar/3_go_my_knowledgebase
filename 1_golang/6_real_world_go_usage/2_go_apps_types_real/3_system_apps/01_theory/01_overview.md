# System apps / daemons

«Системное приложение» — программа, работающая фоном, без GUI, обычно как часть ОС: backup-агент, sync daemon, мониторинг, ETL-планировщик.

## Особенности

- **Long-running**: запущено как systemd-сервис или k8s-pod.
- **Без stdin/stdout interaction** — управляется сигналами и файлами/сокетами.
- **Логи в stderr / journal** — не в свой файл (системе виднее).
- **PID file** — обычно через systemd, не вручную.

## Скелет

```go
func main() {
    cfg := mustLoadConfig()
    setupLogger(cfg)

    ctx, stop := signal.NotifyContext(context.Background(),
        syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
    defer stop()

    app := newApp(cfg)
    if err := app.Run(ctx); err != nil {
        slog.Error("app", "err", err)
        os.Exit(1)
    }
}
```

`signal.NotifyContext` (Go 1.16+) — стандартный способ для graceful shutdown.

## Сигналы

- `SIGINT` / `SIGTERM` — выключение.
- `SIGHUP` — часто означает «перечитать конфиг» (можно реализовать).
- `SIGUSR1` / `SIGUSR2` — кастомные действия (например, dump diagnostic).

## Работа с файловой системой

- Используй `path/filepath` (не вручную через `/`).
- `fsnotify` — для watch'а изменений конфига.
- Atomic write: `tempfile + rename`, иначе rename atomic, partial write — нет.

```go
tmp, _ := os.CreateTemp(filepath.Dir(target), ".tmp-*")
tmp.Write(data)
tmp.Close()
os.Rename(tmp.Name(), target) // atomic на одном FS
```

## Запуск внешних процессов

```go
ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()

cmd := exec.CommandContext(ctx, "bash", "-c", "...")
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr

if err := cmd.Run(); err != nil { ... }
```

`CommandContext` важен — на отмене процесс убивается.

## Daemonization

В Linux традиционный «двойной fork» теперь не нужен — пусть systemd управляет. Запускай как foreground-процесс, systemd сам делает daemonize.

Если без systemd: `nohup app &`, `screen`, `tmux`.

## Автозапуск

- **systemd unit** в `/etc/systemd/system/myapp.service`:

```ini
[Unit]
Description=My App
After=network.target

[Service]
ExecStart=/usr/local/bin/myapp
Restart=on-failure
User=myapp

[Install]
WantedBy=multi-user.target
```

- **Docker / K8s** — современный путь, см. `2_backend`.

## Метрики и health

Даже системный демон выигрывает от наличия `/metrics` и `/healthz`. Минимум — http-listener на отдельном порту.

## Правило

- signal handling = `signal.NotifyContext` + ctx-aware горутины.
- Atomic writes для конфигов и state-файлов.
- Логи в stderr, systemd / docker их подхватит.
- exec через `CommandContext` всегда.
