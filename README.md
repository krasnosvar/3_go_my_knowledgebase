# Go Backend Knowledge Base

Учебная база по Go backend-разработке: язык, стандартная библиотека,
конкурентность, runtime, базы данных, брокеры сообщений, прикладной system
design и базовые темы Computer Science.

Материалы собраны как практический roadmap: краткая теория, runnable-примеры,
упражнения и решения для самопроверки.

## 0. Подготовка

1. Установите Go **1.22 или новее** (корневой `go.mod` требует `go 1.22`;
   часть теории описывает изменения до Go 1.25+): <https://go.dev/doc/install>.
   Проверка: `go version`.
2. Клонируйте репозиторий. В корне лежит единый `go.mod`, поэтому любой
   пример или упражнение запускается через `go run` по пути к папке —
   из корня или из папки темы:

   ```bash
   # из корня репозитория
   go run ./1_golang/1_go_basics/01_variables_constants/03_exercises/01_var_decl
   # или из папки темы
   cd 1_golang/1_go_basics/01_variables_constants
   go run ./03_exercises/01_var_decl
   go run ./04_exercises_answers/01_var_decl_answers
   ```

3. Упражнение — это `main.go` с `TODO`-заготовкой и встроенной самопроверкой:
   в стартовом состоянии оно компилируется и печатает `FAIL`, ваша цель —
   довести до `PASS`. Решение лежит в `04_exercises_answers/NN_name_answers/`.
4. Раздел `5_testing` учит писать тесты: там упражнения — `*_test.go`-заготовки,
   которые падают (`go test` → FAIL), пока тест не написан; примеры и ответы
   проходят:
   `go test ./1_golang/5_testing/01_unit_tests/02_examples/...`.
5. Для части 2 ([2_db_kafka_redis](./2_db_kafka_redis/)) нужны **Docker** и
   **Docker Compose v2** (`docker compose ...`): PostgreSQL, Kafka, Redis,
   ClickHouse и MongoDB поднимаются локальными стендами. Go-примеры там —
   отдельные модули со своими `go.mod` (запуск из их папки, первый раз
   скачиваются зависимости).
6. Проверить всё разом: `go build ./... && go vet ./...` из корня.

## Структура

- [1_golang](./1_golang/) — Go от базового синтаксиса и ООП (включая SOLID и
  паттерны проектирования) до конкурентности, runtime, testing, real-world
  backend-паттернов и алгоритмов.
- [2_db_kafka_redis](./2_db_kafka_redis/) — PostgreSQL, Kafka, Redis,
  ClickHouse, MongoDB, Docker Compose стенды и практические задания.
- [3_system_design](./3_system_design/) — прикладной System Design для
  Go-бэкенда (в связке с канонической базой DevOps/SRE): архитектура сервисов
  и API, производительность и кэширование, распределённые компоненты и сквозные
  кейсы собеседований.
- [other](./other/) — небольшие заметки по общим темам: кодировки, endianness,
  сериализация, процессы, сигналы, время и часовые пояса.

## Формат разделов

Большинство тем устроены по одному шаблону:

- `01_theory/` — краткая теория без лишней воды;
- `02_examples/` — рабочие примеры, которые можно запускать;
- `03_exercises/` — упражнения от простого к сложному;
- `04_exercises_answers/` — решения для самопроверки (`NN_name_answers/`);
- `README.md` — навигация по теме.

## Шпаргалки

Уже знакомы с темой и нужен быстрый справочник? У тем есть `CHEATSHEET.md` —
таблицы «задача → синтаксис/команда», «когда что использовать», сложность,
типичные ошибки. Подробности — по ссылкам на теорию внутри. Основы Go,
ООП, конкурентность, стандартная библиотека и system design — по шпаргалке
на тему; небольшие темы собраны в одну шпаргалку на раздел (`##` на тему).

**1. Go — основы**
- [01. Переменные и константы](./1_golang/1_go_basics/01_variables_constants/CHEATSHEET.md)
- [02. Типы данных](./1_golang/1_go_basics/02_data_types/CHEATSHEET.md)
- [03. Форматированный вывод (fmt)](./1_golang/1_go_basics/03_fmt_and_output/CHEATSHEET.md)
- [04. Операторы и побитовые операции](./1_golang/1_go_basics/04_operators_bitwise/CHEATSHEET.md)
- [05. Управляющие конструкции](./1_golang/1_go_basics/05_control_flow/CHEATSHEET.md)
- [06. Указатели](./1_golang/1_go_basics/06_pointers/CHEATSHEET.md)
- [07. Функции](./1_golang/1_go_basics/07_functions/CHEATSHEET.md)
- [08. Структуры](./1_golang/1_go_basics/08_structs/CHEATSHEET.md)
- [09. Слайсы](./1_golang/1_go_basics/09_slices/CHEATSHEET.md)
- [10. Строки](./1_golang/1_go_basics/10_strings/CHEATSHEET.md)
- [11. Maps](./1_golang/1_go_basics/11_maps/CHEATSHEET.md)
- [12. Ошибки](./1_golang/1_go_basics/12_errors/CHEATSHEET.md)

**1. Go — ООП**
- [01. Методы](./1_golang/2_oop/01_methods/CHEATSHEET.md)
- [02. Интерфейсы](./1_golang/2_oop/02_interfaces/CHEATSHEET.md)
- [03. Встраивание](./1_golang/2_oop/03_embedding/CHEATSHEET.md)
- [04. Type assertions и type switches](./1_golang/2_oop/04_type_assertions_switches/CHEATSHEET.md)
- [05. Generics](./1_golang/2_oop/05_generics/CHEATSHEET.md)
- [06. SOLID и паттерны проектирования](./1_golang/2_oop/06_solid_and_design_patterns/CHEATSHEET.md)

**1. Go — конкурентность**
- [01. Горутины](./1_golang/3_concurrency/01_goroutines/CHEATSHEET.md)
- [02. Каналы](./1_golang/3_concurrency/02_channels/CHEATSHEET.md)
- [03. select](./1_golang/3_concurrency/03_select/CHEATSHEET.md)
- [04. Пакет sync](./1_golang/3_concurrency/04_sync_package/CHEATSHEET.md)
- [05. Пакет context](./1_golang/3_concurrency/05_context/CHEATSHEET.md)
- [06. sync/atomic](./1_golang/3_concurrency/06_atomic/CHEATSHEET.md)
- [07. Паттерны конкурентности](./1_golang/3_concurrency/07_patterns/CHEATSHEET.md)
- [08. Гонки, deadlock, livelock, starvation](./1_golang/3_concurrency/08_race_conditions/CHEATSHEET.md)

**1. Go — стандартная библиотека, тесты, практика, runtime, алгоритмы**
- [fmt, strconv, strings](./1_golang/4_standard_library/01_fmt_strconv_strings/CHEATSHEET.md)
- [regexp, sort, time](./1_golang/4_standard_library/02_regexp_sort_time/CHEATSHEET.md)
- [encoding/json, csv, base64](./1_golang/4_standard_library/03_encoding_json_csv_base64/CHEATSHEET.md)
- [io, bufio, os, path/filepath](./1_golang/4_standard_library/04_io_bufio_os_filepath/CHEATSHEET.md)
- [net/http, net](./1_golang/4_standard_library/05_net_http_net/CHEATSHEET.md)
- [Testing (все темы раздела)](./1_golang/5_testing/CHEATSHEET.md)
- [Real world: части приложения](./1_golang/6_real_world_go_usage/1_app_parts/CHEATSHEET.md)
- [Real world: типы Go-приложений](./1_golang/6_real_world_go_usage/2_go_apps_types_real/CHEATSHEET.md)
- [Go под капотом (runtime, GC, компилятор, профилирование, reflection)](./1_golang/7_under_the_hood/CHEATSHEET.md)
- [Алгоритмы и структуры данных](./1_golang/8_algorithms_dsa/CHEATSHEET.md)

**2. Базы данных и брокеры**
- [PostgreSQL](./2_db_kafka_redis/1_postgres/CHEATSHEET.md)
- [Kafka](./2_db_kafka_redis/2_kafka_rabbit_nats_redis/1_kafka/CHEATSHEET.md)
- [Redis](./2_db_kafka_redis/2_kafka_rabbit_nats_redis/4_redis/CHEATSHEET.md)
- [ClickHouse](./2_db_kafka_redis/4_clickhouse/CHEATSHEET.md)
- [MongoDB](./2_db_kafka_redis/5_mongodb/CHEATSHEET.md)

**3. System design (Go)**
- [01. Архитектура сервисов, сеть и API](./3_system_design/01_architecture_styles/CHEATSHEET.md)
- [02. Производительность, наблюдаемость и кэширование](./3_system_design/02_performance_and_scaling/CHEATSHEET.md)
- [03. Распределённые компоненты, очереди и отказоустойчивость](./3_system_design/03_distributed_components/CHEATSHEET.md)
- [04. System Design интервью и сквозные кейсы](./3_system_design/04_case_studies/CHEATSHEET.md)

**Разное**
- [Computer Science basics (все заметки other/)](./other/CHEATSHEET.md)

## Как проходить

1. Начать с [1_golang](./1_golang/): основы языка, ООП-подход в Go,
   конкурентность, стандартная библиотека и testing.
2. После базового Go перейти к [2_db_kafka_redis](./2_db_kafka_redis/):
   SQL, индексы, транзакции, брокеры сообщений и кэши.
3. Параллельно изучать [3_system_design](./3_system_design/), чтобы связывать
   Go-код с архитектурными решениями: API, данные, очереди, отказоустойчивость,
   наблюдаемость.
4. Использовать [other](./other/) как справочник по темам, которые часто
   всплывают на собеседованиях и в системной разработке.

## Проверка примеров

Весь Go-код части 1 и `3_system_design` входит в корневой модуль:

```bash
go build ./... && go vet ./...
go test ./1_golang/5_testing/.../02_examples/... ./1_golang/5_testing/.../04_exercises_answers/...
```

Упражнения специально содержат `TODO`: это стартовое состояние задачи
(печатает `FAIL`), а готовые варианты лежат рядом в `04_exercises_answers/`.
Go-проекты в `2_db_kafka_redis` — отдельные модули, проверяются из своих папок.

## Полезные ресурсы

- [Go.dev](https://go.dev) — официальная документация Go.
- [Go by Example](https://gobyexample.com) — короткие примеры по стандартным
  темам.
- [Effective Go](https://go.dev/doc/effective_go) — идиомы и практики Go.
- [The Go Programming Language](https://www.gopl.io/) — книга Донована и
  Кернигана.
- [Exercism Go Track](https://exercism.org/tracks/go) — практические
  упражнения.
