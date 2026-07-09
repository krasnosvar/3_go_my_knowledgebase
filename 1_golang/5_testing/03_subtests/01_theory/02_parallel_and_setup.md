# Параллельность и setup/teardown

## t.Parallel

`t.Parallel()` помечает тест как «можно запускать одновременно с другими параллельными».

Правила:

- Тест не должен зависеть от глобального state.
- Тест не должен изменять общие ресурсы без синхронизации.
- Подтесты внутри одного `Test*` тоже могут быть параллельными.

```go
func TestX(t *testing.T) {
    t.Parallel()
    // ...
}
```

С `-race` параллельные тесты — лучший способ ловить data race.

## Параллельные подтесты с loop variable

Классическая ловушка:

```go
for _, tt := range tests {
    tt := tt // важно для Go < 1.22
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()
        // tt уже свой
    })
}
```

Без локальной копии все подтесты увидят последний `tt`.

## TestMain — глобальный setup

Когда нужна инициализация на всю тестовую программу (например, поднять БД-контейнер):

```go
func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}
```

`m.Run()` запускает все тесты пакета. После него — общий cleanup.

`os.Exit` не выполняет defer, так что teardown — до него.

## t.Cleanup — локальный teardown

```go
func TestX(t *testing.T) {
    f := mustOpen(t)
    t.Cleanup(func() { f.Close() })
    // ...
}
```

`Cleanup`:

- работает даже после `t.Fatal`;
- срабатывает в обратном порядке (LIFO);
- предпочтительнее `defer` внутри теста, потому что виден в helper-функциях.

## t.Setenv

```go
t.Setenv("APP_ENV", "test")
```

Устанавливает env только на время теста. По завершении возвращает старое значение. Альтернатива ручному `os.Setenv` + cleanup.

## t.TempDir

```go
dir := t.TempDir()
```

Создаёт временную директорию, **автоматически удаляемую** после теста. Без `Cleanup` руками.

## Helper

```go
func mustOpen(t *testing.T) *os.File {
    t.Helper()
    f, err := os.Open("data.txt")
    if err != nil { t.Fatal(err) }
    return f
}
```

`t.Helper()` — место ошибки в логе будет в вызывающем тесте, не в helper'е.

## Правило

- Маленькие тесты — `t.Parallel()`, чтобы запускались быстрее.
- Долгий setup — `TestMain`.
- Локальный setup — `t.Cleanup` / `t.TempDir` / `t.Setenv`.
- Helper'ы — всегда с `t.Helper()`.
