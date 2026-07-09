# Subtests

`t.Run(name, func(t *testing.T) { ... })` создаёт подтест. Это не отдельный тест-функция, а вложенная подсекция внутри одного `Test*`.

## Базовый пример

```go
func TestStatus(t *testing.T) {
    t.Run("server", func(t *testing.T) {
        if got := Classify(500); got != "server" {
            t.Errorf("got %s", got)
        }
    })
    t.Run("client", func(t *testing.T) {
        if got := Classify(404); got != "client" {
            t.Errorf("got %s", got)
        }
    })
}
```

Запуск конкретного подтеста:

```
go test -run TestStatus/server
```

## Что даёт

- **Изоляция**: каждый `t.Run` — свой `t`, свои `Cleanup`, свой `Parallel`.
- **Понятный вывод**: `--- PASS: TestStatus/server (0.00s)`.
- **Группировка** связанных проверок без 10 отдельных функций.

## С table-driven

Связка `for + t.Run` — самый частый use-case подтестов. См. `02_table_driven_tests`.

## Параллельные подтесты

```go
for _, tt := range tests {
    tt := tt
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()
        // ...
    })
}
```

Все подтесты в этом цикле запустятся параллельно — отлично для тестов с I/O / sleep.

Особенность: внешний тест (`Test*`) ждёт **завершения всех параллельных подтестов перед возвратом**. Это удобно для общего setup/teardown.

## Cleanup в подтесте

```go
t.Run("with file", func(t *testing.T) {
    f := mustCreateTemp(t)
    t.Cleanup(func() { os.Remove(f.Name()) })
    // ...
})
```

`Cleanup` срабатывает после конкретного подтеста, не после всего `Test*`.

## Имя подтеста

- Пробелы и спецсимволы заменяются на `_`.
- Для `-run` используй итоговое имя (с подчёркиванием).
- Хорошие имена короткие и дескриптивные: `valid input`, `nil context`, `bad token`.

## Когда НЕ subtests

- Если каждый случай требует совсем разной подготовки и проверки → отдельные `Test*`.
- Если subtests превращаются в каскад из 4 уровней `t.Run` — это переусложнено.

## Правило

`t.Run` — это базовая структура нормального Go-теста. По умолчанию используй table-driven + subtests, отдельные `Test*` — для крупных независимых сценариев.
