# Table-driven tests: паттерн

Table-driven test — идиоматичный для Go способ покрыть много кейсов в одном тесте: список входов и ожиданий, цикл, общая проверка.

## Базовая форма

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        want     int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Fatalf("Add(%d,%d) = %d; want %d",
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

`t.Run(tt.name, ...)` — каждое имя становится подтестом. Вывод:

```
=== RUN   TestAdd/positive
=== RUN   TestAdd/negative
=== RUN   TestAdd/zero
--- PASS: TestAdd (0.00s)
```

## Зачем

- **Меньше шаблона** — общая логика теста в одном месте.
- **Легко добавить кейс** — одна строка в slice.
- **Селективный запуск**: `go test -run TestAdd/zero`.
- **Лучшие отчёты**: каждый кейс — отдельная строка PASS/FAIL.

## Поля кейса

Минимум: `name`, входы, ожидание. Полезные расширения:

- `wantErr error` или `wantErrIs error` — если функция может возвращать ошибки.
- `setup func(t *testing.T)` — нестандартная подготовка.
- `skip string` — причина пропуска кейса.

## С ошибками

```go
tests := []struct{
    name string
    in   string
    want int
    wantErr error
}{
    {"ok", "42", 42, nil},
    {"empty", "", 0, ErrEmpty},
    {"not number", "abc", 0, ErrFormat},
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        got, err := Parse(tt.in)
        if !errors.Is(err, tt.wantErr) {
            t.Fatalf("err = %v; want %v", err, tt.wantErr)
        }
        if got != tt.want {
            t.Fatalf("got = %d; want %d", got, tt.want)
        }
    })
}
```

`errors.Is` сравнивает по цепочке, не по сообщению.

## Когда НЕ нужен table

- Только 1–2 случая → просто два теста с понятными именами.
- Каждый кейс требует уникальной подготовки → table вырождается в `if name=="x" {...}`.
- Поведение функции изменяется кардинально по входу → разные тесты лучше отражают намерение.

## Правило

Table-driven — default для функций с несколькими «нормальными + краевыми» кейсами. Когда ловишь себя за copy-paste тестов с одинаковой логикой, это ровно тот случай.
