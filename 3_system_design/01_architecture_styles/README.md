# Тема 01. Архитектура сервисов, сеть и API в Go (Architecture & API Styles)

## О чём эта тема (краткий контекст)

Этот модуль связывает фундаментальную архитектуру систем (монолит и
микросервисы, Clean Architecture, Hexagonal / Ports & Adapters, DDD, сетевой
стек и стили API REST/gRPC) с тем, как они идиоматично реализуются на языке Go.
Вы разберёте защиту границ модулей через компилятор (`internal/`), утиную
типизацию интерфейсов на стороне потребителя, работу сетевого поллера `netpoll`
и настройку промышленных HTTP/gRPC серверов и клиентов.

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [01. Монолит, модульный монолит и микросервисы](../../../5_devops_sre_knowledgebase/14_system_design/01_architecture_and_api_styles/01_theory/01_system_architecture_monolith_microservices.md)
> - [02. Внутренняя архитектура сервиса: Clean Architecture, Hexagonal и DDD](../../../5_devops_sre_knowledgebase/14_system_design/01_architecture_and_api_styles/01_theory/02_service_architecture_clean_ddd.md)
> - [03. Стили API (REST, gRPC, GraphQL), API Gateway и Service Discovery](../../../5_devops_sre_knowledgebase/14_system_design/01_architecture_and_api_styles/01_theory/03_api_styles_gateway_and_discovery.md)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — раскладка
Go-проекта, Clean Architecture и DI в Go, тюнинг `http.Server`/`http.Transport`,
REST vs gRPC, ловушки.

## Пререквизиты

- [2_oop](../../1_golang/2_oop/) (особенно интерфейсы и
  [06_solid_and_design_patterns](../../1_golang/2_oop/06_solid_and_design_patterns/)).
- [05_net_http_net](../../1_golang/4_standard_library/05_net_http_net/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория Go-специфики:
    *   [01_monolith_microservices_and_project_layout.md](./01_theory/01_monolith_microservices_and_project_layout.md)
        — модульный монолит vs микросервисы в Go, защита границ через
        `internal/`, статические бинарники и 12-Factor App.
    *   [02_clean_architecture_and_ddd_in_go.md](./01_theory/02_clean_architecture_and_ddd_in_go.md)
        — Clean Architecture, Ports & Adapters и DDD в Go: «accept interfaces,
        return structs», явный DI в `main.go` vs `wire`/`fx`, борьба с
        оверинжинирингом.
    *   [03_networking_and_api_styles_in_go.md](./01_theory/03_networking_and_api_styles_in_go.md)
        — рантайм `netpoll` (`epoll`/`kqueue`), таймауты `http.Server`, пул
        соединений `http.Transport`, утечки сокетов, REST (`encoding/json`)
        против `grpc-go` и Protobuf.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_dependency_injection](./02_examples/01_dependency_injection/main.go)
        — явный Dependency Injection через конструктор и подмена на фейк.
    *   [02_clean_layers](./02_examples/02_clean_layers/main.go) — три слоя
        Clean Architecture (`domain` → `usecase` → `adapter`), связанные через
        интерфейс у потребителя.
    *   [03_ports_adapters](./02_examples/03_ports_adapters/main.go) —
        Hexagonal (Ports & Adapters): порт `Notifier` и сменные адаптеры.
    *   [04_tcp_echo](./02_examples/04_tcp_echo/main.go) — TCP echo-сервер и
        клиент (`net.Listen`, `net.Dial`, горутина на соединение).
    *   [05_udp_datagram](./02_examples/05_udp_datagram/main.go) — обмен
        UDP-датаграммами (`net.ListenPacket`, `ReadFrom`/`WriteTo`).
    *   [06_http_server_client](./02_examples/06_http_server_client/main.go) —
        HTTP-сервер и клиент с заголовками, кодами ответа и JSON.
    *   [07_rest_crud](./02_examples/07_rest_crud/main.go) — REST CRUD-сервис
        (`POST`/`GET`/`DELETE`) с проверкой статус-кодов.
    *   [08_pagination](./02_examples/08_pagination/main.go) — сравнение
        offset- и cursor-пагинации на Go.
    *   [09_idempotency](./02_examples/09_idempotency/main.go) — идемпотентный
        `POST /payments` через заголовок `Idempotency-Key`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL`:
    *   `01_implement_repository` — реализация интерфейса репозитория для
        сервисного слоя.
    *   `02_constructor_injection` — внедрение зависимости через конструктор
        `NewOrderService`.
    *   `03_fix_dependency_direction` — разворот зависимости через интерфейс
        на стороне потребителя (DIP).
    *   `04_swap_adapter` — подмена адаптера порта без изменения ядра.
    *   `05_status_category` — классификация HTTP статус-кодов (`1xx`–`5xx`).
    *   `06_parse_url` — разбор URL на компоненты (`scheme`, `host`, `port`,
        `path`).
    *   `07_tcp_line_protocol` — текстовый протокол `UPPER <text>\n` поверх
        `net.Conn`.
    *   `08_http_handler` — реализация `http.Handler` с валидацией метода и
        параметров.
    *   `09_udp_sum_server` — UDP-сервер, суммирующий числа в датаграмме.
    *   `10_method_semantics` — проверка свойств HTTP-методов (`Safe`,
        `Idempotent`).
    *   `11_paginate` — безопасная пагинация слайса по `offset` и `limit`.
    *   `12_rest_status` — маппинг доменной ошибки и метода в HTTP-статус.
    *   `13_parse_filter` — разбор и валидация query-строки фильтрации.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения.

## 🚀 Как проходить тему

1. Изучите фундаментальную теорию в DevOps-базе по ссылкам в начале страницы,
   затем прочитайте все уроки в [01_theory/](./01_theory/).
2. Запустите примеры из папки раздела:
   ```bash
   go run ./02_examples/02_clean_layers/main.go
   go run ./02_examples/07_rest_crud/main.go
   ```
3. Решите упражнения из [03_exercises/](./03_exercises/) (список и описание —
   в [03_exercises/README.md](./03_exercises/README.md)):
   ```bash
   go run ./03_exercises/01_implement_repository/main.go
   ```
4. Сверьтесь с ответами в [04_exercises_answers/](./04_exercises_answers/).

## Когда двигаться дальше

Когда вы уверенно проектируете слои Go-сервиса без циклических импортов и
знаете, какие таймауты и лимиты пула соединений обязательно выставлять в
`http.Server` и `http.Transport`, переходите к
[02_performance_and_scaling](../02_performance_and_scaling/).
