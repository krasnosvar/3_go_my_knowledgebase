# 02. Каналы (channels)

Канал — типизированная труба между горутинами: передаёт значения и одновременно синхронизирует отправителя и получателя. Здесь — буферизованные и небуферизованные каналы, закрытие и `range`, направленные каналы, fan-out/fan-in и простой pipeline.

**Зачем:** «Не общайтесь через общую память — делитесь памятью через общение»: каналы — идиоматичный способ строить конкурентный код в Go.

**Пререквизиты:** [01_goroutines](../01_goroutines/).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — синтаксис, поведение nil/открытого/закрытого канала, выбор буфера, направленные каналы, правила `close`, канал или mutex, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — создание, отправка/чтение, блокировки.
    *   [02_buffered.md](./01_theory/02_buffered.md) — буферизованные каналы, `len`/`cap`.
    *   [03_close_and_range.md](./01_theory/03_close_and_range.md) — закрытие канала и чтение через `range`.
    *   [04_directional.md](./01_theory/04_directional.md) — направленные каналы `chan<-` / `<-chan`.
    *   [05_common_pitfalls.md](./01_theory/05_common_pitfalls.md) — частые ошибки: deadlock, паника при отправке в закрытый канал, nil-каналы.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — передача результата из горутины.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — worker над буферизованной очередью задач.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — producer, закрывающий канал, и `range`.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_unbuffered](./03_exercises/01_unbuffered/main.go) — передача значения через небуферизованный канал.
    *   [02_buffered](./03_exercises/02_buffered/main.go) — заполнение буферизованного канала без читателя.
    *   [03_close_range](./03_exercises/03_close_range/main.go) — producer закрывает канал, consumer читает через `range`.
    *   [04_directional](./03_exercises/04_directional/main.go) — sender/receiver с направленными каналами.
    *   [05_fan_out](./03_exercises/05_fan_out/main.go) — несколько worker'ов читают один канал, порядок результатов сохранён.
    *   [06_fan_in](./03_exercises/06_fan_in/main.go) — слияние двух каналов с закрытием выхода.
    *   [07_timeout](./03_exercises/07_timeout/main.go) — чтение с таймаутом через `time.After`.
    *   [08_pipeline](./03_exercises/08_pipeline/main.go) — конвейер из стадий «умножить» → «отфильтровать».
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/02_channels`).

1. Прочитайте теорию по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_unbuffered/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_unbuffered_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [03_select](../03_select/) — ожидание нескольких каналов сразу.
