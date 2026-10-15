# Pitfalls type assertions

## 1. Assertion без ok = panic

```go
s := i.(string) // если i не строка — runtime panic
```

В коде, работающем с пользовательскими/внешними данными — **всегда** через `s, ok := ...`.

## 2. Assertion к wrong receiver

Метод объявлен на `*T`. Через `interface{}` прилетает `T`:

```go
var v interface{} = T{}
_, ok := v.(*T) // false, внутри T, а не *T
```

Проверяй именно тот вид, который ожидаешь, или приводи заранее.

## 3. Assertion к интерфейсу с pointer receiver

```go
type I interface{ Foo() }

type S struct{}
func (*S) Foo() {}

var v interface{} = S{}
_, ok := v.(I) // false: value S не реализует I
```

Как и везде: метод с `*T` нужен указатель внутри интерфейса.

## 4. Ambiguity с одноимёнными интерфейсами

```go
type A interface{ X() }
type B interface{ X() }
```

Assertion к `A` и `B` — это два разных assertion'а. Даже если метод один и тот же, компилятор относится к ним как к разным.

## 5. Забыть default в type switch

```go
switch x := v.(type) {
case int: /* ... */
case string: /* ... */
}
```

Если прилетит `float64` — ничего не произойдёт, баг. Всегда имей `default`, особенно на входе из внешнего мира.

## 6. fallthrough

Не работает в type switch. Если нужно общее поведение для нескольких типов — перечисли в одном `case: A, B, C`.

## 7. Стоимость

Type switch по многим типам — хоть и быстрая операция, но в цикле на миллионы значений заметна. Для перфа лучше интерфейс с методом.

## 8. Nil-интерфейс vs typed nil

```go
var p *MyType
var v interface{} = p // v не nil, но внутри nil-указатель

_, ok := v.(*MyType) // ok = true, но p может быть nil
```

После assertion — проверяй nil явно.

## 9. Ассерт к своему пакету

Идиома проверки: `if _, ok := v.(*mytype); ok` внутри своего же пакета — это часто сигнал плохой архитектуры. Открой сначала интерфейс с нужными методами.

## 10. Смешение type assertion и conversion

```go
x := int(y) // conversion — работает с конкретными типами
x := y.(int) // assertion — только из интерфейсного значения
```

На этапе ревью легко перепутать. Следи за контекстом.
