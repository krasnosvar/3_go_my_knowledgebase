# switch

`switch` в Go — это удобная замена длинным цепочкам `if/else if`.

## Классический switch по значению

По умолчанию нет fallthrough — каждый `case` сам завершает блок.

```go
switch day {
case "sat", "sun":
    fmt.Println("weekend")
case "fri":
    fmt.Println("almost weekend")
default:
    fmt.Println("weekday")
}
```

## switch без выражения

Работает как цепочка `if/else if`, удобно для сложных условий:

```go
switch {
case age < 18:
    fmt.Println("minor")
case age < 65:
    fmt.Println("adult")
default:
    fmt.Println("senior")
}
```

## switch с инициализацией

Похож на `if` с короткой декларацией:

```go
switch x := compute(); {
case x < 0:
    // ...
case x == 0:
    // ...
}
```

## fallthrough

Явно передает управление в следующий `case`. Используется редко, только когда реально нужно:

```go
switch n {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("<= 2")
}
```

## type switch

Отдельная конструкция для работы с интерфейсами:

```go
switch v := x.(type) {
case int:
    fmt.Println("int", v)
case string:
    fmt.Println("string", v)
}
```
