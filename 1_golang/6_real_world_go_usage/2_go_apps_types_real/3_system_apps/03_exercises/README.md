# Упражнения: System apps / daemons

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_signal_handling
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_signal_handling_answers
```

## Список задач (от простого к сложному)

1. [`01_signal_handling`](./01_signal_handling/main.go) — обработка сигналов демоном; решение: [`01_signal_handling_answers`](../04_exercises_answers/01_signal_handling_answers/main.go).
2. [`02_graceful_stop`](./02_graceful_stop/main.go) — упорядоченная остановка компонентов; решение: [`02_graceful_stop_answers`](../04_exercises_answers/02_graceful_stop_answers/main.go).
3. [`03_pid_file`](./03_pid_file/main.go) — PID-файл против двойного запуска; решение: [`03_pid_file_answers`](../04_exercises_answers/03_pid_file_answers/main.go).
4. [`04_config_path`](./04_config_path/main.go) — поиск пути к конфигу; решение: [`04_config_path_answers`](../04_exercises_answers/04_config_path_answers/main.go).
5. [`05_stderr_logging`](./05_stderr_logging/main.go) — логгер демона в stderr; решение: [`05_stderr_logging_answers`](../04_exercises_answers/05_stderr_logging_answers/main.go).
6. [`06_disk_check`](./06_disk_check/main.go) — проверка размера каталога; решение: [`06_disk_check_answers`](../04_exercises_answers/06_disk_check_answers/main.go).
7. [`07_exec_timeout`](./07_exec_timeout/main.go) — запуск внешней команды с таймаутом; решение: [`07_exec_timeout_answers`](../04_exercises_answers/07_exec_timeout_answers/main.go).
8. [`08_health_probe`](./08_health_probe/main.go) — HTTP health probe с таймаутом; решение: [`08_health_probe_answers`](../04_exercises_answers/08_health_probe_answers/main.go).
