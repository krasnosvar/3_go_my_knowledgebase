# Паттерны работы с ошибками

## 1. Early return

Главный идиом: сразу выходим при ошибке, основной путь остаётся плоским.

```go
if err := step1(); err != nil {
    return fmt.Errorf("step1: %w", err)
}
if err := step2(); err != nil {
    return fmt.Errorf("step2: %w", err)
}
return nil
```

Никаких лестниц из `if err == nil { ... }`.

## 2. Контекстное оборачивание

Каждый слой добавляет «где мы и что делали», а не повторяет нижний текст:

```
repo:   save user 42: insert user: duplicate key
handler: create user: save user 42: insert user: duplicate key
```

Читается как трасса снизу вверх.

## 3. Sentinel + Is

Публичный sentinel для «типовых» состояний:

```go
var ErrNotFound = errors.New("not found")

if errors.Is(err, ErrNotFound) { return http.StatusNotFound }
```

## 4. Типизированная ошибка + As

Если нужны детали (поле, код, причина) — структура и `errors.As`:

```go
var ve *ValidationError
if errors.As(err, &ve) { /* использовать ve.Field */ }
```

## 5. «Не глотать» ошибку

```go
// плохо
_ = doIt()

// плохо
if err != nil {
    // ничего
}
```

Либо обрабатывай, либо возвращай, либо логируй — но не все три сразу и не ни одно.

## 6. Дифференциация по уровню

- Библиотеке: возвращать типизированные/sentinel-ошибки.
- Сервису: оборачивать и логировать.
- Транспорту: маппить в HTTP/gRPC-коды и user-facing сообщения.

## 7. errors.Join для множественных проблем

Если собираешь результат параллельных проверок:

```go
err := errors.Join(validateName(n), validateAge(a))
```

Получатель пройдётся через `errors.Is/As` по каждому.
