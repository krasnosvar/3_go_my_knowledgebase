# Method promotion

Когда тип встраивает другой тип, методы встроенного «продвигаются» в method set внешнего. Это и называется method promotion.

```go
type Logger struct{}

func (Logger) Info(msg string) { fmt.Println("INFO:", msg) }

type Service struct {
    Logger
    Name string
}

s := Service{Name: "users"}
s.Info("started") // доступно через Service
```

## Правила продвижения

1. Продвигаются методы, у которых **method set встроенного типа** содержит метод.
2. При встраивании `T` внутрь `Outer`:
   - методы `T` доступны у `Outer` и `*Outer`;
3. При встраивании `*T`:
   - методы `T` и `*T` доступны у `Outer` и `*Outer`.
4. Методы не дублируются, просто «пробрасываются» через синтаксис.

## Интерфейсы и promotion

Если встроенный тип реализует интерфейс, внешний тоже автоматически:

```go
type Reader interface{ Read(p []byte) (int, error) }

type FancyReader struct {
    *bufio.Reader
    stats Stats
}

var r Reader = FancyReader{} // OK, Read продвинут
```

Это частый приём, когда нужно «обернуть» существующий тип, добавив поля, но сохранить совместимость с интерфейсом.

## Ограничения

- Нельзя продвигать методы «в глубину по иерархии», если один и тот же метод определён на двух разных встроенных типах — это конфликт имён (см. соответствующий файл).
- Продвижение — это синтаксический сахар, оно не создает новый метод; рефлексия видит его на встроенном типе.

## Специфический случай: метод с receiver другого типа

```go
type Inner struct{}
func (*Inner) Do() {}

type Outer struct{ Inner }
```

`Outer.Do` доступен через `*Outer`, но НЕ через value `Outer` — как обычно, метод `*T` не попадает в method set `T`.

```go
var o Outer
// o.Do() // ошибка, если Outer не addressable
(&o).Do() // OK
```

## Проверка

В сомнительных случаях смотри сигнатуру:

```go
method, ok := reflect.TypeOf(Outer{}).MethodByName("Do")
```

Если продвинулся — будет найден у внешнего типа.
