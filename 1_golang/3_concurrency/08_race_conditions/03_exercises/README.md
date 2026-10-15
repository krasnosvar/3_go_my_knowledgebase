# Упражнения: race_conditions

Каждая задача — это `main.go` с функцией/методом-заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии тесты печатают `FAIL` и не зависают — ожидания ограничены таймаутами) — ваша цель довести их до `PASS`. Только стандартная библиотека.

Запускайте с флагом `-race`: решение засчитывается, только если печатает `PASS` и детектор гонок молчит.

Запуск шаблона (из папки темы):

```bash
go run -race ./03_exercises/01_detect_race/main.go
```

Запуск эталонного решения:

```bash
go run -race ./04_exercises_answers/01_detect_race_answers/main.go
```

## Список задач (от простого к сложному)

1. [01_detect_race](./01_detect_race/main.go) — воспроизвести гонку с `-race` и исправить через atomic. Решение: [01_detect_race_answers](../04_exercises_answers/01_detect_race_answers/main.go).
2. [02_fix_with_mutex](./02_fix_with_mutex/main.go) — инвариант баланса: проверка и списание под одним мьютексом. Решение: [02_fix_with_mutex_answers](../04_exercises_answers/02_fix_with_mutex_answers/main.go).
3. [03_fix_with_channels](./03_fix_with_channels/main.go) — единственная горутина-владелец состояния. Решение: [03_fix_with_channels_answers](../04_exercises_answers/03_fix_with_channels_answers/main.go).
4. [04_deadlock](./04_deadlock/main.go) — встречные переводы: единый порядок захвата мьютексов. Решение: [04_deadlock_answers](../04_exercises_answers/04_deadlock_answers/main.go).
5. [05_livelock](./05_livelock/main.go) — «вежливый» захват с `TryLock` и случайной паузой. Решение: [05_livelock_answers](../04_exercises_answers/05_livelock_answers/main.go).
6. [06_starvation](./06_starvation/main.go) — приоритетная очередь без голодания низкого приоритета. Решение: [06_starvation_answers](../04_exercises_answers/06_starvation_answers/main.go).
7. [07_ordering](./07_ordering/main.go) — happens-before через закрытие канала. Решение: [07_ordering_answers](../04_exercises_answers/07_ordering_answers/main.go).
8. [08_review_checklist](./08_review_checklist/main.go) — квиз по ревью конкурентного кода. Решение: [08_review_checklist_answers](../04_exercises_answers/08_review_checklist_answers/main.go).
