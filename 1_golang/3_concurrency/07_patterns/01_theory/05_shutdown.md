# Graceful shutdown

Корректное завершение сервиса — это не «убил процесс». Это: «перестал принимать новое, доделал текущее, закрыл ресурсы, сообщил системе».

## Базовая схема

```go
func main() {
    ctx, cancel := signal.NotifyContext(context.Background(),
        syscall.SIGINT, syscall.SIGTERM)
    defer cancel()

    srv := &http.Server{Addr: ":8080", Handler: handler()}

    go func() {
        _ = srv.ListenAndServe()
    }()

    <-ctx.Done()

    shutdownCtx, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel2()
    _ = srv.Shutdown(shutdownCtx)
}
```

## Что должно быть в shutdown

1. **Перестать принимать новое**: `srv.Shutdown` для HTTP, close входного канала для воркеров.
2. **Дождаться текущих операций**: WaitGroup, errgroup.Wait, http.Server.Shutdown.
3. **Закрыть ресурсы**: БД-соединения, клиенты к внешним сервисам, файлы.
4. **Флашить буферы**: логи, метрики, traces.

## Сигналы

- `SIGINT` (Ctrl+C) — дев/локал.
- `SIGTERM` — оркестратор просит закончиться (Kubernetes, systemd).
- `SIGKILL` — нельзя поймать, процесс убит жёстко.

## Таймаут shutdown

`Shutdown(ctx)` с контекстом отдельного таймаута. Если не успели за N секунд — возвращаем ошибку и выходим. Это важно, чтобы не подвиснуть навсегда.

## Worker shutdown

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return
        case j, ok := <-jobs:
            if !ok {
                return
            }
            process(ctx, j)
        }
    }
}
```

Два канала выхода: ctx и закрытие jobs. В зависимости от семантики — закрыть jobs чтобы «доработать очередь» или отменить ctx чтобы «бросить всё».

## Порядок в большом сервисе

Типично это выглядит так:

1. Получили SIGTERM → cancel().
2. HTTP-сервер уходит в Shutdown: принимает текущие соединения, отказывает новым.
3. Background workers уходят по ctx.Done().
4. Ждём все WaitGroup.Wait() / errgroup.Wait().
5. Закрываем БД, MQ-клиентов, логгер.
6. Выходим.

Порядок важен: нельзя закрыть БД, пока handler'ы ещё её дергают.

## Ошибки

- Не ловить SIGTERM — kubernetes/systemd решат «не отвечает» и пошлют SIGKILL.
- Зависать в shutdown дольше grace period — процесс получит SIGKILL, ресурсы потеряются.
- Не закрывать ресурсы в defer → утечки между перезапусками.

## Readiness/liveness

- В момент shutdown правильно **сразу** отвечать 503 на readiness — чтобы орk's перестал слать трафик.
- Liveness при этом остаётся ок — мы жив, но «собираемся».
