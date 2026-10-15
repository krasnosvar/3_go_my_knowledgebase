# Интерфейсы — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## Объявление и реализация

Подробно: [01_basics.md](./01_theory/01_basics.md).

| Задача | Код |
| :--- | :--- |
| Объявить | `type Speaker interface { Speak() string }` |
| Реализовать (неявно, без `implements`) | `func (Dog) Speak() string { return "woof" }` |
| Принять интерфейс | `func Shout(s Speaker) { fmt.Println(strings.ToUpper(s.Speak())) }` |
| Пустой интерфейс | `any` (Go 1.18+) = `interface{}` — любое значение |
| Compile-time проверка реализации | `var _ Speaker = (*Dog)(nil)` (под объявлением типа) |
| Композиция | `type ReadWriter interface { Reader; Writer }` |
| Устройство значения | пара (dynamic type, dynamic value), 16 байт на 64-бит |

## Pointer vs value receiver

Подробно: [02_pointer_vs_value.md](./01_theory/02_pointer_vs_value.md).

| Методы типа | `T` реализует? | `*T` реализует? |
| :--- | :--- | :--- |
| Все на `T` | да | да |
| Хоть один нужный метод на `*T` | **нет** | да |

```go
type Incr interface{ Inc() }
func (c *Counter) Inc() { c.v++ }

var x Incr = &Counter{} // OK
// var y Incr = Counter{} // ошибка: Counter does not implement Incr (method Inc has pointer receiver)
```

## Type assertion и type switch

Подробно: [03_assertions_and_switch.md](./01_theory/03_assertions_and_switch.md); вся тема — [04_type_assertions_switches](../04_type_assertions_switches/CHEATSHEET.md).

| Задача | Код |
| :--- | :--- |
| Достать тип (паника при промахе) | `s := i.(string)` |
| Безопасно | `s, ok := i.(string)` |
| Опциональное расширение | `if rw, ok := w.(io.ReadWriter); ok { ... }` |
| Разобрать несколько типов | `switch x := v.(type) { case int: ...; case nil: ...; default: ... }` |
| Ошибки по цепочке `Unwrap` | `errors.Is` / `errors.As` вместо assertion |

## Nil interface

Подробно: [04_nil_interface.md](./01_theory/04_nil_interface.md).

| Значение | `== nil` |
| :--- | :--- |
| `var err error` | true (type = nil, value = nil) |
| `var p *MyError; var err error = p` | **false** (type = `*MyError`, value = nil) |

```go
package main

import "fmt"

type MyError struct{}

func (*MyError) Error() string { return "boom" }

var _ error = (*MyError)(nil) // compile-time проверка реализации

func bad() error {
	var e *MyError // nil-указатель
	return e       // интерфейс (type=*MyError, value=nil) != nil
}

func good() error { return nil }

func main() {
	fmt.Println(bad() == nil, good() == nil) // false true
	fmt.Printf("%T\n", bad())                // *main.MyError
}
```

Правило: нет ошибки → `return nil` голым литералом; различить typed nil — `errors.As(err, &t) && t != nil`.

## Дизайн

Подробно: [05_design_principles.md](./01_theory/05_design_principles.md), [06_use_cases.md](./01_theory/06_use_cases.md).

| Принцип | Как |
| :--- | :--- |
| Маленькие интерфейсы | 1–3 метода: `io.Reader`, `io.Writer` |
| Интерфейс — у потребителя | `service` объявляет `UserRepo`, `postgresRepo` о нём не знает |
| Accept interfaces, return structs | `func NewService(repo UserRepo) *Service` |
| Не «на будущее» | вводить при ≥ 2 реализациях, для моков или контракта между модулями |
| Имя | `-er` для одного метода (`Reader`, `Notifier`); собирательное для нескольких (`http.Handler`) |
| `any` | только когда тип действительно неизвестен (JSON в `map[string]any`); иначе generics |
| Типичные точки расширения | репозиторий, платёжный адаптер, notifier, логгер, кэш, CLI-команда `Run(args []string) error` |

## Ловушки

- ⚠️ Возврат typed nil (`var e *MyError; return e`) → `err != nil` всегда true, хендлер отвечает 500.
- ⚠️ Метод на `*T`, а в интерфейс кладут `T` → «does not implement (pointer receiver)».
- ⚠️ `i.(T)` без `ok` на внешних данных → panic.
- ⚠️ Длинный type switch с одинаковым вызовом в каждой ветке → поведение должно быть методом интерфейса.
- ⚠️ Большой «толстый» интерфейс → трудно реализовать и мокать; дробите и комбинируйте.
- ⚠️ Публичные поля/сеттеры «для интерфейса» бесполезны: через интерфейс видны только методы.
