# Methods: основы

Метод в Go — это обычная функция с дополнительным параметром — **receiver**, который записывается перед именем метода.

## Синтаксис

```go
type Rectangle struct {
    W, H float64
}

func (r Rectangle) Area() float64 {
    return r.W * r.H
}
```

- `(r Rectangle)` — receiver: тип и имя.
- Через `r` обращаемся к полям как к переменной.

## Вызов

```go
rect := Rectangle{W: 10, H: 5}
rect.Area() // 50
```

Метод вызывается через точку, как поле.

## Методы на любом именованном типе

Receiver может быть не только структурой:

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}
```

Ограничение: нельзя определять методы на типах из чужих пакетов. Если нужно — заведи свой именованный тип (не alias: `type IntSet = []int` — тот же `[]int`, методы на нём объявить нельзя):

```go
type IntSet []int

func (s IntSet) Contains(v int) bool {
    for _, x := range s {
        if x == v { return true }
    }
    return false
}
```

## Публичность

Имя метода с заглавной буквы — экспортируется из пакета, со строчной — приватное. Те же правила, что и у полей.

## Что метод НЕ делает

- Не перегружается по типам аргументов (нельзя два `Add` с разными сигнатурами).
- Не наследуется (классов нет).
- Не может добавляться к типу из чужого пакета.

## Отличие от функции

Семантически метод — это функция, закреплённая за типом. Go позволяет и обратное превращение:

```go
f := rect.Area        // method value
fmt.Println(f())      // 50
```

Или method expression:

```go
g := Rectangle.Area   // func(r Rectangle) float64
fmt.Println(g(rect))  // 50
```

Это удобно для callback'ов.
