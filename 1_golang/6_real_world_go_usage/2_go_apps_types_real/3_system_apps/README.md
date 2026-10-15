# System apps / daemons

Системные утилиты и демоны: сигналы, упорядоченная остановка, PID-файл, пути к конфигу, логи в stderr, работа с ФС, внешние процессы и probe.

**Уже знакомы с темой?** → [CHEATSHEET.md](../CHEATSHEET.md#3-system-apps) — раздел «3. System apps» общей шпаргалки типов Go-приложений.

## Пререквизиты

[1_cli_tools](../1_cli_tools/), [3_concurrency/05_context](../../../3_concurrency/05_context/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_overview.md](./01_theory/01_overview.md) — System apps / daemons.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib, завершаются сами):
    *   [01_signal_shutdown](./02_examples/01_signal_shutdown/main.go) — остановка по сигналу через `signal.NotifyContext`.
    *   [02_file_scan](./02_examples/02_file_scan/main.go) — обход файлов каталога.
    *   [03_exec_timeout_shape](./02_examples/03_exec_timeout_shape/main.go) — запуск внешней команды с таймаутом.
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL` (заготовки компилируются сразу):
    *   [01_signal_handling](./03_exercises/01_signal_handling/main.go) — обработка сигналов демоном.
    *   [02_graceful_stop](./03_exercises/02_graceful_stop/main.go) — упорядоченная остановка компонентов.
    *   [03_pid_file](./03_exercises/03_pid_file/main.go) — PID-файл против двойного запуска.
    *   [04_config_path](./03_exercises/04_config_path/main.go) — поиск пути к конфигу.
    *   [05_stderr_logging](./03_exercises/05_stderr_logging/main.go) — логгер демона в stderr.
    *   [06_disk_check](./03_exercises/06_disk_check/main.go) — проверка размера каталога.
    *   [07_exec_timeout](./03_exercises/07_exec_timeout/main.go) — запуск внешней команды с таймаутом.
    *   [08_health_probe](./03_exercises/08_health_probe/main.go) — HTTP health probe с таймаутом.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения (`<имя>_answers`).

## 🚀 Как проходить тему

Команды запускаются из корня репозитория (там лежит `go.mod`) или из каталога темы `1_golang/6_real_world_go_usage/2_go_apps_types_real/3_system_apps` с относительными путями `./02_examples/...`.

1. Прочитайте теорию: [01_overview.md](./01_theory/01_overview.md).
2. Запустите примеры:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/3_system_apps/02_examples/01_signal_shutdown
   ```
3. Решайте задачи по порядку; заготовка печатает `FAIL`, пока задача не решена:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/3_system_apps/03_exercises/01_signal_handling
   ```
4. Сверьтесь с решением:
   ```bash
   go run ./1_golang/6_real_world_go_usage/2_go_apps_types_real/3_system_apps/04_exercises_answers/01_signal_handling_answers
   ```

## Когда дальше

Когда все задачи печатают `PASS`: [4_grpc_services](../4_grpc_services/) — gRPC services.
