# 01. Горутины (goroutines)

Горутина — легковесная единица выполнения, которой управляет рантайм Go. С неё начинается вся конкурентность: запуск через `go`, ожидание завершения, ловушка с переменной цикла, утечки и корректная остановка.

**Зачем:** без горутин не написать ни HTTP-сервер, ни фоновый обработчик; ошибки здесь (утечки, гонки, «потерянные» горутины) — самые частые в продакшене и на собеседованиях.

**Пререквизиты:** [1_go_basics](../../1_go_basics/) целиком (функции, замыкания, структуры, срезы) и [2_oop](../../2_oop/) (методы, интерфейсы). В упражнениях заранее встречаются `sync.WaitGroup` и каналы в простейшем виде — подробно они разобраны в [02_channels](../02_channels/) и [04_sync_package](../04_sync_package/).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — запуск и модель GMP, ожидание (`WaitGroup`, `wg.Go`, errgroup), захват переменной цикла, утечки, отладка, ловушки.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_basics.md](./01_theory/01_basics.md) — запуск через `go`, планировщик, стоимость горутины.
    *   [02_waiting.md](./01_theory/02_waiting.md) — ожидание завершения: `WaitGroup`, каналы.
    *   [03_loop_capture.md](./01_theory/03_loop_capture.md) — захват переменной цикла и изменения Go 1.22.
    *   [04_leaks.md](./01_theory/04_leaks.md) — утечки горутин и как их избегать.
    *   [05_debugging.md](./01_theory/05_debugging.md) — отладка: `runtime.NumGoroutine`, стеки, pprof.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практические сценарии.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — параллельный запуск команд с `WaitGroup`.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — фоновые worker'ы и `runtime.NumGoroutine`.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — параллельная загрузка с записью по индексу.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_basic_go_keyword](./03_exercises/01_basic_go_keyword/main.go) — запуск n горутин через `go` без ожидания.
    *   [02_waitgroup](./03_exercises/02_waitgroup/main.go) — конкурентный запуск задач и ожидание через `WaitGroup`.
    *   [03_race_awareness](./03_exercises/03_race_awareness/main.go) — общий счётчик без data race (`Mutex`).
    *   [04_loop_capture](./03_exercises/04_loop_capture/main.go) — явная передача переменной цикла в горутину.
    *   [05_channel_signal](./03_exercises/05_channel_signal/main.go) — сигнал о завершении закрытием канала.
    *   [06_worker_basic](./03_exercises/06_worker_basic/main.go) — worker, читающий задачи из канала до закрытия.
    *   [07_leak_prevention](./03_exercises/07_leak_prevention/main.go) — генератор с путём остановки — без утечки горутины.
    *   [08_graceful_stop](./03_exercises/08_graceful_stop/main.go) — остановка нескольких горутин и ожидание через `WaitGroup`.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/01_goroutines`).

1. Прочитайте теорию по порядку, начиная с [01_basics.md](./01_theory/01_basics.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_basic_go_keyword/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_basic_go_keyword_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [02_channels](../02_channels/) — каналы как основной способ общения горутин.
