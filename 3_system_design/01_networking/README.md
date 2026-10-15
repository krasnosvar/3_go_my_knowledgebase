# 01. Networking (Сеть)

Базовый сетевой фундамент, на котором держатся все распределённые системы: как данные ходят между узлами, чем отличаются транспортные протоколы (TCP/UDP), как устроены модель OSI/TCP/IP, HTTP всех версий, DNS и TLS. Теория подкреплена runnable-примерами на стандартной библиотеке Go и упражнениями с самопроверкой.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория, разбитая по подтемам:
    *   [01_tcp_vs_udp.md](./01_theory/01_tcp_vs_udp.md) — TCP против UDP: соединение и handshake, гарантии доставки и порядка, границы сообщений, накладные расходы, когда что использовать.
    *   [02_osi_tcpip_model.md](./01_theory/02_osi_tcpip_model.md) — модель OSI (7 уровней) и стек TCP/IP (4 уровня), инкапсуляция, где живут HTTP, TCP, IP; L4 vs L7 балансировка.
    *   [03_http_https.md](./01_theory/03_http_https.md) — HTTP/1.1: структура запроса/ответа, методы и идемпотентность, коды статусов по категориям, ключевые заголовки, keep-alive; что добавляет HTTPS.
    *   [04_http2.md](./01_theory/04_http2.md) — HTTP/2: бинарные фреймы, мультиплексирование потоков, сжатие заголовков HPACK, server push (устарел), приоритизация; оставшийся TCP head-of-line blocking.
    *   [05_http3_quic.md](./01_theory/05_http3_quic.md) — HTTP/3 и QUIC поверх UDP: устранение HOL-блокировки, 0-RTT рукопожатие, встроенный TLS 1.3, connection migration.
    *   [06_dns.md](./01_theory/06_dns.md) — DNS: иерархия имён, рекурсивный резолвинг, типы записей (A/AAAA/CNAME/MX/NS/TXT/…), кэширование и TTL, почему UDP и когда TCP.
    *   [07_tls_handshake.md](./01_theory/07_tls_handshake.md) — TLS-хендшейк по шагам: сертификаты и цепочка доверия, TLS 1.3 (1-RTT), ECDHE и forward secrecy, зачем симметрика + асимметрика.
    *   [08_use_cases.md](./01_theory/08_use_cases.md) — практические применения: выбор транспорта и версии HTTP, correctness-паттерны API (retry/timeout/pooling), DNS и TLS в эксплуатации, диагностика по слоям.
*   **[02_examples/](./02_examples/)** — runnable-примеры (каждый поднимает сервер в горутине на `127.0.0.1:0`, обменивается данными и завершается):
    *   [01_tcp_echo](./02_examples/01_tcp_echo/main.go) — TCP echo: клиент шлёт строку, сервер возвращает её обратно; разграничение сообщений по `'\n'`.
    *   [02_udp_datagram](./02_examples/02_udp_datagram/main.go) — обмен одной UDP-датаграммой через `ListenPacket`/`ReadFrom`/`WriteTo`; сохранение границ сообщений.
    *   [03_http_server_client](./02_examples/03_http_server_client/main.go) — `net/http` + `httptest.NewServer`: GET-запрос клиентом с таймаутом, чтение статуса, заголовков и тела.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи с автопроверками `PASS`/`FAIL` (заготовки компилируются сразу):
    *   `01_status_category` — по коду HTTP-статуса вернуть категорию 1xx..5xx (или unknown).
    *   `02_parse_url` — разобрать URL на части (scheme/host/port/path/query) через `net/url`.
    *   `03_tcp_line_protocol` — строчный протокол поверх TCP: прочитать строку и вернуть её в верхнем регистре.
    *   `04_http_handler` — HTTP-хэндлер с корректным статусом (200/400) и заголовком `Content-Type`.
    *   `05_udp_sum_server` — **UDP-сервер, суммирующий числа из одной датаграммы (Хардкор)**.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач.

## 🚀 Как проходить тему

1. Прочитайте теорию по подтемам в [01_theory/](./01_theory/), начиная с [01_tcp_vs_udp.md](./01_theory/01_tcp_vs_udp.md).
2. Запустите и изучите примеры из папки `02_examples/` через `go run` (все офлайн, только stdlib; работает корневой `go.mod` репозитория):
   ```bash
   go run ./02_examples/01_tcp_echo/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/05_udp_sum_server/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.
