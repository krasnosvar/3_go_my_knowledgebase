# Arrange – Act – Assert

Идиома для структуры теста: каждый тест разделён на три явные фазы.

## Структура

```go
func TestRegister(t *testing.T) {
    // Arrange — подготовка
    repo := newFakeRepo()
    svc := New(repo, fakeClock{})

    // Act — действие
    err := svc.Register(context.Background(), "alice@example.com")

    // Assert — проверка
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if !repo.saved {
        t.Errorf("expected repo.saved = true")
    }
}
```

Фазы можно отделять пустой строкой или комментариями — выбор стиля. Главное — каждая фаза узнаваема.

## Зачем

- Тест читается за 5 секунд. Глаза легко находят «что мы готовим», «что вызываем», «что ожидаем».
- Маленькие тесты с одной AAA-структурой проще поддерживать, чем 50 строк смешанной логики.
- Дисциплинирует мысль: сначала продумай контекст, потом действие, потом ожидание.

## Сигналы плохого теста

- Несколько `Act` подряд — лучше разбить на отдельные тесты.
- `Assert` спрятан внутри цикла без понятной цели.
- 30+ строк подготовки — вынеси в helper / setup.

## Helpers

Когда подготовка повторяется — выносится в helper:

```go
func newTestService(t *testing.T) (*Service, *fakeRepo) {
    t.Helper()
    repo := &fakeRepo{}
    return New(repo), repo
}
```

`t.Helper()` — пометка, что место ошибки нужно показывать в вызывающем тесте, а не внутри хелпера.

## Setup / teardown

В Go нет `BeforeAll/AfterAll` фреймворков. Вместо этого:

- `TestMain(m *testing.M)` — глобальная инициализация.
- `t.Cleanup(func() { ... })` — освобождение ресурсов в конце теста (Go 1.14+).

```go
func TestStuff(t *testing.T) {
    f, _ := os.CreateTemp("", "test-*")
    t.Cleanup(func() { os.Remove(f.Name()) })
    // ...
}
```

`Cleanup` гарантирует выполнение даже при `t.Fatal`.

## Параллельность

```go
func TestX(t *testing.T) {
    t.Parallel()
    // ...
}
```

Тесты с `Parallel` запускаются одновременно. Полезно для скорости. Условие: тест не должен зависеть от глобального state. С `t.Parallel()` особенно важно `-race`.

## Правило

- Чёткая AAA-структура — best default.
- Длинные тесты разбивай на маленькие.
- Используй `t.Helper`, `t.Cleanup`, `t.Parallel` — это инструменты языка, а не «продвинутые фишки».
