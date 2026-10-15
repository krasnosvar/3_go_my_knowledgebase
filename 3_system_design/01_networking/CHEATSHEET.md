# Networking — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## TCP vs UDP

Подробно: [01_tcp_vs_udp.md](./01_theory/01_tcp_vs_udp.md).

| Характеристика | TCP | UDP |
| :--- | :--- | :--- |
| Соединение | да: `SYN → SYN-ACK → ACK` | нет |
| Доставка | гарантирована (retransmission) | best-effort |
| Порядок | сохраняется (sequence numbers) | не гарантируется, возможны дубли |
| Границы сообщений | нет — поток байтов | есть — одна отправка = одна датаграмма |
| Flow / congestion control | receive window / slow start | нет (делает приложение) |
| Заголовок | 20+ байт | 8 байт |
| Задержка | выше (ACK, retrans) | ниже |

| Задача | Транспорт |
| :--- | :--- |
| HTTP/1.1, HTTP/2, gRPC, БД, файлы | TCP |
| DNS-запросы, VoIP, видео, игры, statsd-метрики | UDP |
| HTTP/3, мобильные клиенты с нестабильной сетью | QUIC поверх UDP |

## Модель OSI и TCP/IP

Подробно: [02_osi_tcpip_model.md](./01_theory/02_osi_tcpip_model.md).

| OSI | TCP/IP | Протоколы | Единица |
| :--- | :--- | :--- | :--- |
| L7 Application / L6 Presentation / L5 Session | Application | HTTP, DNS, gRPC, TLS (условно) | данные |
| L4 Transport | Transport | TCP, UDP — порты | сегмент / датаграмма |
| L3 Network | Internet | IP, ICMP — адреса, маршрутизация | пакет |
| L2 Data Link / L1 Physical | Link | Ethernet, Wi-Fi, ARP, MAC | кадр (frame) / биты |

| Балансировщик | Видит | Умеет |
| :--- | :--- | :--- |
| L4 | IP:port, соединения TCP/UDP | распределять соединения, не глядя в содержимое |
| L7 | HTTP | маршрутизировать по URL, заголовкам, cookie |

## HTTP/1.1 и HTTPS

Подробно: [03_http_https.md](./01_theory/03_http_https.md).

| Метод | Safe | Идемпотентный | Тело запроса |
| :--- | :--- | :--- | :--- |
| `GET` | да | да | нет |
| `HEAD` | да | да | нет (ответ без тела) |
| `OPTIONS` | да | да | — (методы, CORS) |
| `PUT` | нет | да | замена целиком |
| `DELETE` | нет | да | — |
| `PATCH` | нет | не гарантируется | частичное изменение |
| `POST` | нет | нет | создание / действие |

| Код | Смысл |
| :--- | :--- |
| 1xx | `100 Continue`, `101 Switching Protocols` |
| 2xx | `200 OK`, `201 Created`, `204 No Content` |
| 3xx | `301 Moved Permanently`, `302 Found`, `304 Not Modified` |
| `400` / `404` / `429` | невалидный запрос / нет ресурса / слишком много запросов |
| `401` vs `403` | не аутентифицирован vs аутентифицирован, но нет прав |
| `500` | ошибка сервера |
| `502` | прокси получил невалидный ответ от upstream |
| `503` | сервис временно недоступен / перегружен |
| `504` | upstream не ответил вовремя |

| Заголовок | Назначение |
| :--- | :--- |
| `Host` | виртуальный хост, обязателен в HTTP/1.1 |
| `Content-Type` / `Content-Length` | MIME-тип / размер тела |
| `Transfer-Encoding: chunked` | тело частями без известной длины |
| `Authorization: Bearer ...` | креды |
| `Cache-Control`, `ETag`, `If-None-Match` | кэширование |
| `Connection: keep-alive` | переиспользовать TCP (в 1.1 — по умолчанию) |

HTTPS = HTTP поверх TLS, порт 443: конфиденциальность, целостность, аутентификация сервера.

## HTTP/2 и HTTP/3

Подробно: [04_http2.md](./01_theory/04_http2.md), [05_http3_quic.md](./01_theory/05_http3_quic.md).

| Версия | Транспорт | HOL-блокировка | Рукопожатие | Фишки |
| :--- | :--- | :--- | :--- | :--- |
| HTTP/1.1 | TCP | на уровне приложения | TCP + TLS отдельно | текст, keep-alive, 6+ соединений на домен в браузере |
| HTTP/2 | TCP | только на уровне TCP | TCP + TLS | бинарные фреймы, мультиплексирование streams, HPACK, приоритеты; server push устарел → `103 Early Hints` |
| HTTP/3 | QUIC/UDP | нет | 1-RTT, повторно 0-RTT; TLS 1.3 встроен | независимые потоки, connection migration по Connection ID |

| Выбор | Когда |
| :--- | :--- |
| HTTP/1.1 | простые сервисы, ручная отладка, низкая конкурентность |
| HTTP/2 | много параллельных запросов к одному хосту, gRPC |
| HTTP/3 | публичные сервисы с мобильной аудиторией, критична latency |

## DNS

Подробно: [06_dns.md](./01_theory/06_dns.md).

| Запись | Назначение |
| :--- | :--- |
| `A` / `AAAA` | имя → IPv4 / IPv6 |
| `CNAME` | имя → другое имя (не на apex зоны) |
| `MX` | почтовый сервер с приоритетом |
| `NS` | authoritative-серверы зоны |
| `TXT` | SPF, DKIM, верификация |
| `SOA` | параметры зоны (serial, TTL) |
| `PTR` | IP → имя |
| `SRV` | адрес + порт сервиса |

| Факт | Детали |
| :--- | :--- |
| Резолвинг | клиент → рекурсивный резолвер → root → TLD → authoritative; ответ кэшируется |
| Транспорт | UDP/53; TCP при TC=1 (truncated) и для AXFR; EDNS0 — большие UDP-ответы; DoT/DoH — приватность |
| TTL большой / маленький | меньше нагрузка, медленные изменения / быстрый failover, больше запросов |
| Миграция | снизить TTL (≈60 с) заранее → переключить → дождаться → вернуть TTL |
| Kubernetes | `svc.namespace.svc.cluster.local` (CoreDNS) |

## TLS

Подробно: [07_tls_handshake.md](./01_theory/07_tls_handshake.md).

| Шаг TLS 1.3 (1-RTT) | Что передаётся |
| :--- | :--- |
| `ClientHello` | версии, cipher suites, `key_share` клиента |
| `ServerHello` | выбранный cipher, `key_share` сервера, `{Certificate}`, `{CertificateVerify}`, `{Finished}` |
| клиент `{Finished}` | далее — данные, зашифрованные симметричным ключом |

| Понятие | Суть |
| :--- | :--- |
| Проверка сертификата | цепочка до доверенного root CA; не истёк/не отозван (CRL/OCSP); SAN совпадает с доменом |
| ECDHE | обе стороны вычисляют общий секрет, приватные части не передаются |
| Forward secrecy | ephemeral-ключи на сессию: утечка ключа сервера не расшифрует старый трафик |
| TLS 1.2 vs 1.3 | 2 RTT, устаревшие шифры, RSA без FS vs 1 RTT (0-RTT при возобновлении), только AEAD, FS обязателен |
| Асимметрика + симметрика | асимметрика — один раз в хендшейке; трафик — быстрый симметричный ключ |
| Эксплуатация | терминация на LB/ingress; mTLS между сервисами; ротация через ACME / cert-manager |

## Диагностика (снизу вверх)

Подробно: [08_use_cases.md](./01_theory/08_use_cases.md).

| Вопрос | Слой | Команда |
| :--- | :--- | :--- |
| Есть IP-связность? | L3 | `ping host`, `traceroute host` |
| Порт открыт? | L4 | `nc -vz host port` / `telnet host port` |
| TLS валиден? | TLS | `openssl s_client -connect host:443` |
| HTTP отвечает? | L7 | `curl -v https://host/path` |
| Имя резолвится? | DNS | `dig host` / `nslookup host` |

| Go-идиома | Код |
| :--- | :--- |
| Таймаут запроса | `context.WithTimeout` + `http.Client{Timeout: ...}` |
| Пул keep-alive соединений | `http.Transport{MaxIdleConnsPerHost: N}` |
| Retry для POST | заголовок `Idempotency-Key` |

## Ловушки

- ⚠️ TCP не сохраняет границы: один `Write` ≠ один `Read` — нужен framing (длина/разделитель).
- ⚠️ HTTP/2 убирает HOL только на уровне HTTP: потеря TCP-сегмента тормозит все потоки.
- ⚠️ После смены IP старый адрес резолвится до истечения TTL в кэшах.
- ⚠️ CNAME на apex зоны запрещён.
- ⚠️ Ретраить бездумно 4xx бессмысленно; кандидаты на retry с backoff — 5xx и 429.
- ⚠️ Клиент без таймаута → висящие соединения исчерпают пул.
- ⚠️ UDP/QUIC режут корпоративные фаерволы → нужен fallback на HTTP/2.
