# 08. Гонки, deadlock, livelock, starvation

Типовые ошибки конкурентного кода: data race, deadlock, livelock и starvation, а также race detector (`-race`) и стратегии исправления.

**Зачем:** эти ошибки не ловятся компилятором, плохо воспроизводятся и проявляются под нагрузкой — их нужно уметь узнавать по симптомам и предотвращать на ревью.

**Пререквизиты:** темы [01](../01_goroutines/)–[07](../07_patterns/) этого раздела.

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_data_race.md](./01_theory/01_data_race.md) — data race и почему это неопределённое поведение.
    *   [02_deadlock.md](./01_theory/02_deadlock.md) — deadlock и порядок захвата блокировок.
    *   [03_livelock_starvation.md](./01_theory/03_livelock_starvation.md) — livelock и starvation.
    *   [04_race_detector.md](./01_theory/04_race_detector.md) — race detector: `go run/test -race`.
    *   [05_fix_strategies.md](./01_theory/05_fix_strategies.md) — стратегии исправления гонок.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — где гонки встречаются чаще всего.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — напоминание о запуске с `-race`.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — запись в map под `Mutex`.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — чистая остановка горутины через закрытие канала.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу, подробности — в [03_exercises/README.md](./03_exercises/README.md)):
    *   [01_detect_race](./03_exercises/01_detect_race/main.go) — воспроизвести гонку с `-race` и исправить через atomic.
    *   [02_fix_with_mutex](./03_exercises/02_fix_with_mutex/main.go) — инвариант баланса: проверка и списание под одним мьютексом.
    *   [03_fix_with_channels](./03_exercises/03_fix_with_channels/main.go) — единственная горутина-владелец состояния.
    *   [04_deadlock](./03_exercises/04_deadlock/main.go) — встречные переводы: единый порядок захвата мьютексов.
    *   [05_livelock](./03_exercises/05_livelock/main.go) — «вежливый» захват с `TryLock` и случайной паузой.
    *   [06_starvation](./03_exercises/06_starvation/main.go) — приоритетная очередь без голодания низкого приоритета.
    *   [07_ordering](./03_exercises/07_ordering/main.go) — happens-before через закрытие канала.
    *   [08_review_checklist](./03_exercises/08_review_checklist/main.go) — квиз по ревью конкурентного кода.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения тех же задач (`NN_name_answers`).

## 🚀 Как проходить тему

Все команды выполняются из папки темы (`1_golang/3_concurrency/08_race_conditions`).

1. Прочитайте теорию по порядку, начиная с [01_data_race.md](./01_theory/01_data_race.md).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_cli_utilities/main.go
   ```
3. Решайте задачи в `03_exercises/` по порядку и запускайте самопроверку (с детектором гонок):
   ```bash
   go run -race ./03_exercises/01_detect_race/main.go
   ```
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`:
   ```bash
   go run -race ./04_exercises_answers/01_detect_race_answers/main.go
   ```

## ➡️ Что дальше

Когда все задачи печатают `PASS`: [4_standard_library](../../4_standard_library/), а для углубления — [7_under_the_hood](../../7_under_the_hood/) (планировщик и модель памяти).
