# Дизайн и type assertions / switches

Когда стоит использовать ассерты и switch — а когда это сигнал проблемы архитектуры.

## 1. Assertion как способ **опционального** поведения

```go
type Closer interface{ Close() error }

func process(r io.Reader) {
    // ...
    if c, ok := r.(Closer); ok {
        _ = c.Close()
    }
}
```

Сценарий: «если реализация поддерживает это расширение — используем». Это идиома stdlib (`io.WriterTo`, `sort.Interface`).

## 2. Type switch в маршалерах/декодерах

```go
switch v := rawVal.(type) {
case bool: /* ... */
case float64: /* ... */
case string: /* ... */
}
```

Уместен, когда мы **правда** не знаем тип заранее (декодированный JSON).

## 3. Assertion к ошибке

```go
var ne *net.OpError
if errors.As(err, &ne) {
    // доступны поля ne.Op, ne.Addr
}
```

`errors.As` делает то же самое, что assertion, но плюс идёт по цепочке `Unwrap`. Всегда предпочти `errors.As`, если работаешь с ошибками.

## 4. Type switch вместо полиморфизма — anti-pattern

Если у тебя 5 кейсов, каждый вызывает одноимённый метод — сделай интерфейс:

```go
// плохо
switch x := v.(type) {
case *Cat: x.Speak()
case *Dog: x.Speak()
}

// хорошо
type Speaker interface{ Speak() }
v.(Speaker).Speak()
```

## 5. Generics вместо any-assert

Если функция формально принимает `any`, а потом делает `.(int)` — скорее всего, это кандидат на generic:

```go
func Sum(xs []any) any { /* ... assertion ... */ }

// лучше
func Sum[T Number](xs []T) T { /* ... */ }
```

## 6. Единое правило для входа

При работе с внешними данными:

1. На границе (handler/decoder) ассертим/switch'им один раз.
2. Дальше по коду ходят уже типизированные значения.

Избегай ассертов «внутри внутри внутри» — это как `interface{}` в Java, сразу видно плохую архитектуру.

## 7. Документация

Если метод принимает `interface{}` и делает ассерт к ограниченному набору типов — **документируй**. Пользователь иначе не узнает.

## Итог

Type assertion и switch — мощные, но локальные инструменты. Если ты ими часто пользуешься — пересмотри, нет ли возможности выразить логику через интерфейс или generic.
