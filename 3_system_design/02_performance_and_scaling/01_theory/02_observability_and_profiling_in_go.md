# Профилирование (`pprof`, `trace`) и наблюдаемость в Go

## О чём эта тема (краткий контекст)

В распределённой системе под высокой нагрузкой невозможно угадать причину
деградации «на глаз»: тормозить может как сериализация JSON, так и ожидание
мьютекса или утечка горутин. Наблюдаемость (метрики, структурные логи,
распределённый трейсинг) отвечает на вопрос «что и где сломалось», а встроенные
в рантайм Go профилировщики (`pprof` и `runtime/trace`) показывают конкретную
функцию и строчку кода, потребляющую CPU, память или блокирующую потоки.

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [Узкие места, закон Литтла и методология профилирования](../../../../5_devops_sre_knowledgebase/14_system_design/02_performance_caching_and_scaling/01_theory/01_bottlenecks_and_profiling_methodology.md)
> - Три столпа наблюдаемости, Prometheus, Loki, OpenTelemetry и алертинг:
>   [8_observability](../../../../5_devops_sre_knowledgebase/8_observability/README.md)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

## 1. Боевое профилирование в продакшене: `net/http/pprof`

Одно из главных инженерных преимуществ Go — промышленный семплирующий
профилировщик, встроенный прямо в рантайм с накладными расходами `< 1–2%`, что
позволяет снимать профили прямо с живого продакшена.

⚠️ **Правило безопасности:** никогда не регистрируйте `net/http/pprof` на
публичном `http.DefaultServeMux`, который смотрит в интернет! Поднимайте для
метрик (`/metrics`), проб (`/healthz`, `/readyz`) и `/debug/pprof/` отдельный
внутренний `http.ServeMux` на закрытом порту (например, `:9090`).

| Профиль `pprof` | Команда снятия | Какую проблему в Go находит |
| :--- | :--- | :--- |
| **CPU (`profile`)** | `go tool pprof http://...:9090/debug/pprof/profile?seconds=30` | Функции, сжигающие процессор (парсинг, криптография, работа GC `runtime.mallocgc` / `runtime.gcBgMarkWorker`). |
| **Heap (`heap` — `inuse_space`)** | `go tool pprof -inuse_space http://...:9090/debug/pprof/heap` | Что прямо сейчас занимает оперативную память в куче (утечки в глобальных `map`, незакрытые ресурсы). |
| **Allocs (`allocs` — `alloc_objects`)** | `go tool pprof -alloc_objects http://...:9090/debug/pprof/allocs` | Функции, создающие миллионы временных объектов и разгоняющие сборщик мусора (даже если память тут же освобождается). |
| **Goroutine (`goroutine`)** | `curl http://...:9090/debug/pprof/goroutine?debug=2` | Утечки горутин: стек-трейсы всех висящих горутин и время их блокировки на канале или сокете. |
| **Mutex / Block (`mutex`, `block`)** | `go tool pprof http://...:9090/debug/pprof/mutex` | Конкуренция (lock contention) на `sync.Mutex` / `sync.RWMutex` и ожидание на каналах. (Требует включения `runtime.SetMutexProfileFraction(5)`). |

### Сравнение двух снимков (`-base`) для поиска утечек

Чтобы найти медленную утечку памяти или горутин, снимите базовый профиль и
профиль через 15 минут и передайте оба в `pprof` с флагом `-diff_base`:
`go tool pprof -diff_base=heap1.pb.gz heap2.pb.gz`.

## 2. Трассировка планировщика: `runtime/trace`

Когда CPU свободен, а латентность p99 всё равно высокая, `pprof` может ничего не
показать, потому что `pprof` видит только работу на CPU. В этом случае снимается
5-секундный трейс рантайма:

```bash
curl -o trace.out http://127.0.0.1:9090/debug/pprof/trace?seconds=5
go tool trace trace.out
```

`go tool trace` с точностью до микросекунды показывает на таймлайне каждого `P`:
фазы Stop-The-World и Mark Assist сборщика мусора, задержки в очереди
планировщика (`SchedWait`), блокировки на системных вызовах и сетевом поллере.

## 3. Экспорт метрик рантайма Go (`runtime/metrics`)

Помимо бизнес- и RED-метрик запросов (Rate, Errors, Duration — см. пример
`02_examples/02_inprocess_metrics`), каждый Go-сервис обязан экспортировать
метрики состояния рантайма (в Prometheus client_golang они включены в
`collectors.NewGoCollector()`):

- `/sched/goroutines:goroutines` — число живых горутин (монотонный рост =
  утечка);
- `/memory/classes/heap/objects:bytes` — размер живой кучи;
- `/gc/cycles/total:gc-cycles` и `/gc/pauses:seconds` — частота и длительность
  пауз GC;
- `/cpu/classes/gc/mark/assist:cpu-seconds` — сколько времени бизнес-горутины
  потратили на принудительную помощь сборщику мусора.

## 4. Структурное логирование без лишних аллокаций (`log/slog`)

Начиная с Go 1.21 стандартная библиотека включает пакет `log/slog` для
структурного логирования. На высоком RPS наивный вызов `slog.Info("msg", "id",
123)` упаковывает аргументы в `any` (`interface{}`), что может вызывать
аллокацию в куче. Для горячего пути используйте типизированные атрибуты
`slog.LogAttrs`:

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

// Zero-allocation вызов на горячем пути через типизированные slog.Attr:
logger.LogAttrs(ctx, slog.LevelInfo, "request completed",
	slog.String("method", r.Method),
	slog.Int("status", http.StatusOK),
	slog.Duration("latency", elapsed),
)
```

Сквозной `trace_id` и `request_id` передаются через `context.Context` и могут
автоматически извлекаться кастомным `slog.Handler`, оборачивающим стандартный
`JSONHandler`. См. рабочий пример `02_examples/01_structured_logging/main.go`.
