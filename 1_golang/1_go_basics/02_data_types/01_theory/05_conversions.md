# Приведение типов (type conversion)

В Go нет неявных конверсий. Даже совместимые по смыслу типы нужно приводить явно.

## Базовый синтаксис

```go
var i int = 42
var f float64 = float64(i) // int -> float64
var u uint = uint(i)       // int -> uint
```

Без `float64(i)` компилятор даст ошибку.

## string <-> []byte / []rune

```go
s := "hello"
b := []byte(s)  // копия данных
s2 := string(b) // копия обратно

runes := []rune("Привет") // каждая буква — один элемент
```

Каждая такая конверсия копирует память.

## Числовые конверсии

```go
var x int32 = 1000
var y int = int(x) // OK

var big int64 = 1 << 40
var small int32 = int32(big) // усечение, если не помещается
```

При сужающем приведении старшие биты теряются — Go не предупреждает об этом.

## Что НЕЛЬЗЯ конвертировать напрямую

- `int` <-> `string` **не** конвертируется как число -> цифры:

```go
// Неправильно:
s := string(65) // "A" (rune), а не "65"

// Правильно:
s := strconv.Itoa(65) // "65"
n, _ := strconv.Atoi("42")
```

- Несовместимые типы (`struct{ A int }` и `struct{ B int }`) — отдельные типы, несмотря на одинаковый layout.

## Custom-типы и alias-типы

```go
type UserID int

var id UserID = 1
var raw int = int(id) // нужно явно
```

Custom-тип — новая сущность даже при одинаковом underlying type.

## unsafe.Pointer (кратко)

Для редких кейсов, когда нужно обойти систему типов (zero-copy `[]byte` <-> `string` и т. п.), есть `unsafe.Pointer`. Он опасен и ломает переносимость, используй осознанно.

```go
import "unsafe"
s := "hello"
_ = unsafe.Pointer(&s)
```
