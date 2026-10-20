# System Design интервью и сквозные кейсы в Go — шпаргалка

Подробности — в [01_theory/](./01_theory/) и каноническом разделе
[DevOps System Design](../../../5_devops_sre_knowledgebase/14_system_design/07_case_studies_and_interview/README.md).
⚠️ — частая ловушка.

## Опорные числа Go для расчётов на салфетке

Подробно:
[01_estimations_and_go_in_interviews.md](./01_theory/01_estimations_and_go_in_interviews.md).

| Ресурс / Операция в Go | Оценка для интервью |
| :--- | :--- |
| Сутки в секундах | $86\,400 \approx 10^5$ с ($\implies 1\text{ млн запросов/день} \approx 12\text{ QPS}$) |
| Стартовый стек 1 горутины | **2–4 КБ** ($100\,000$ горутин $pprox 200\text{–}400\text{ МБ}$) |
| Память на 1 idle TCP/WebSocket соединение | **~10–30 КБ** ($100\,000$ соединений $pprox 1.5\text{–}3\text{ ГБ RAM}$) |
| Пропускная способность 1 канала (`chan T`) | **~5–15 млн опер./с** (при более высокой нагрузке шардировать по бакетам) |
| Пропускная способность HTTP/gRPC на 1 vCPU Go | **~15 000 – 40 000 RPS** (для лёгких I/O запросов) |

```go
package main

import "fmt"

func EstimateQPS(dau int64, actionsPerUser int64, peakFactor float64) (avg float64, peak float64) {
	avg = float64(dau*actionsPerUser) / 86400.0
	peak = avg * peakFactor
	return avg, peak
}

func main() {
	avg, peak := EstimateQPS(10_000_000, 10, 3.0)
	fmt.Printf("avg=%.0f QPS, peak=%.0f QPS
", avg, peak) // avg=1157 QPS, peak=3472 QPS
}
```

## Ключевые узлы 4 классических кейсов на Go

Подробно: [02_case_studies_in_go.md](./01_theory/02_case_studies_in_go.md).

| Кейс | Главные технические решения в Go |
| :--- | :--- |
| **1. Денежные переводы** | Суммы строго в `int64` (копейки/центы) или `decimal`; в одной SQL-транзакции (`*sql.Tx`): `INSERT idempotency_key ON CONFLICT DO NOTHING` + `SELECT ... FOR UPDATE` в порядке `min(fromID, toID)` + запись события в таблицу `outbox`. |
| **2. Галерея / Медиа** | Загрузка файлов клиентом напрямую в S3 по Presigned URL; если через Go — потоковый `io.Copy` с `http.MaxBytesReader` (без `io.ReadAll`!). Ресайз в фоновом Worker Pool. |
| **3. URL Shortener** | Генерация ID через Snowflake или батч-диапазоны в памяти (`atomic.Uint64`) → Base62 (`[0-9a-zA-Z]`); чтение через L1 sharded map + L2 Redis под защитой `singleflight.Group`. |
| **4. News Feed** | Гибридный Fan-out: для обычных авторов пуш `post_id` в Redis через `errgroup` с `SetLimit(N)` и Redis Pipeline; для знаменитостей — слияние в памяти при чтении. |

## Ловушки на интервью

- ⚠️ Считать память на 100 000 соединений в Go по размеру стека потока ОС (8 МБ)
  — завышение оценки памяти в ~300 раз.
- ⚠️ Использовать `float64` для хранения баланса или суммы платежа.
- ⚠️ Блокировать счета в `SELECT ... FOR UPDATE` в произвольном порядке (`from`,
  затем `to`) — гарантированный взаимный клинч (deadlock) при встречных
  переводах.
- ⚠️ Вызывать `io.ReadAll(r.Body)` при загрузке пользовательских медиафайлов —
  мгновенный `OOMKilled` при параллельных загрузках.
- ⚠️ Ходить в Redis/PostgreSQL на каждый клик по вирусной короткой ссылке без
  схлопывания параллельных запросов через `singleflight`.
