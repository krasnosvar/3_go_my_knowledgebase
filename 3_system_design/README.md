# Раздел 3. Прикладной System Design в Go (System Design)

## О чём эта тема (краткий контекст)

System Design отвечает на вопрос, как из отдельных сервисов, баз данных, кэшей
и очередей собрать систему, которая выдерживает высокую нагрузку, переживает
отказы узлов и остаётся поддерживаемой. В этом разделе мы не дублируем общую
платформенную теорию, а фокусируемся на том, как архитектурные паттерны
реализуются и эксплуатируются непосредственно в коде и рантайме **Go**.

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [Раздел 14. System Design и Высоконагруженные Системы (DevOps / SRE)](../../5_devops_sre_knowledgebase/14_system_design/README.md)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

## Формат подразделов

Каждая тема устроена по единому шаблону:

- `01_theory/` — теория Go-специфики со ссылками на фундаментальную базу DevOps;
- `02_examples/` — самодостаточные runnable-примеры на стандартной библиотеке;
- `03_exercises/` — практические задачи с автопроверкой `PASS`/`FAIL`;
- `04_exercises_answers/` — эталонные решения (`NN_name_answers/`);
- `CHEATSHEET.md` — плотная шпаргалка для быстрого повторения.

## Пререквизиты

- Пройдены [1_go_basics](../1_golang/1_go_basics/),
  [2_oop](../1_golang/2_oop/) (включая
  [06_solid_and_design_patterns](../1_golang/2_oop/06_solid_and_design_patterns/))
  и [3_concurrency](../1_golang/3_concurrency/).
- Примеры и упражнения запускаются из папки подтемы или из корня репозитория:
  `go run ./3_system_design/01_architecture_styles/02_examples/01_dependency_injection/main.go`.

## Подразделы

1. [01_architecture_styles](./01_architecture_styles/) — модульный монолит и
   микросервисы в Go (`cmd/`, `internal/`), Clean Architecture, Hexagonal
   (Ports & Adapters), тактический DDD, утиная типизация интерфейсов, рантайм
   `netpoll`, тюнинг `http.Server`/`http.Transport`, REST (`encoding/json`) vs
   `grpc-go` и Protobuf.
2. [02_performance_and_scaling](./02_performance_and_scaling/) — планировщик
   `G-M-P`, настройка `GOMAXPROCS` (защита от CFS CPU throttling в контейнерах)
   и `GOMEMLIMIT`, escape analysis, боевое профилирование (`net/http/pprof`,
   `runtime/trace`), структурные логи `log/slog`, потокобезопасный L1-кэш,
   `singleflight` и обход пауз GC на больших кэшах.
3. [03_distributed_components](./03_distributed_components/) — пул соединений
   `database/sql` (`MaxOpenConns`, `MaxIdleConns`), роутинг на реплики и
   Consistent Hashing, Transactional Outbox, идемпотентные консьюмеры,
   естественный backpressure на каналах и `errgroup`, Rate Limiting, Circuit
   Breaker, Retries с jitter и Graceful Shutdown по `SIGTERM`.
4. [04_case_studies](./04_case_studies/) — опорные числа рантайма Go для
   расчётов ёмкости на System Design интервью и реализация ключевых узлов
   четырёх сквозных кейсов: Денежные переводы, Медиа-галерея, URL Shortener и
   News Feed.

## 🚀 Как проходить

Идите последовательно от `01_architecture_styles` к `04_case_studies`: сначала
устройство отдельного Go-сервиса и его сетевого слоя → затем производительность
рантайма, профилирование и кэширование → затем распределённые данные, очереди и
отказоустойчивость → и в финале сборка сквозных систем.
