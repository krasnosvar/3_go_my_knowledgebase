# Упражнения: Service layer

Каждая задача — `main.go` с заготовкой (`TODO`) и встроенной самопроверкой, печатающей `PASS` или `FAIL`. Заготовки компилируются сразу (в стартовом состоянии печатают `FAIL`) — доведите их до `PASS`. Только стандартная библиотека, без сети и внешних сервисов.

Запуск шаблона (из каталога темы):

```bash
go run ./03_exercises/01_validate_input
```

Запуск эталонного решения:

```bash
go run ./04_exercises_answers/01_validate_input_answers
```

## Список задач (от простого к сложному)

1. [`01_validate_input`](./01_validate_input/main.go) — валидация входа в service layer; решение: [`01_validate_input_answers`](../04_exercises_answers/01_validate_input_answers/main.go).
2. [`02_translate_errors`](./02_translate_errors/main.go) — сервис зависит от интерфейса репозитория и переводит его ошибки; решение: [`02_translate_errors_answers`](../04_exercises_answers/02_translate_errors_answers/main.go).
3. [`03_business_rule`](./03_business_rule/main.go) — бизнес-правило перевода денег; решение: [`03_business_rule_answers`](../04_exercises_answers/03_business_rule_answers/main.go).
4. [`04_unique_email`](./04_unique_email/main.go) — правило уникальности и зависимость от часов; решение: [`04_unique_email_answers`](../04_exercises_answers/04_unique_email_answers/main.go).
