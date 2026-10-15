# Встраивание (embedding) — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## Синтаксис

Подробно: [01_basics.md](./01_theory/01_basics.md).

| Задача | Код |
| :--- | :--- |
| Встроить тип (анонимное поле) | `type Car struct { Engine; Brand string }` |
| Литерал | `Car{Engine: Engine{Power: 150}, Brand: "Toyota"}` — имя поля = имя типа |
| Promoted поле / метод | `c.Power`, `c.Start()` = `c.Engine.Power`, `c.Engine.Start()` |
| Встроить указатель | `type Car struct { *Engine; Brand string }` |
| Встроить интерфейс в struct | `type loggingRepo struct { UserRepo; log *slog.Logger }` |
| Встроить интерфейс в интерфейс | `type ReadWriter interface { Reader; Writer }` |
| Встроить generic-тип | `type LoggedCache[K comparable, V any] struct { *Cache[K, V]; log *slog.Logger }` |

Это композиция, не наследование: `Car` не является `Engine`, нет `super`, нет виртуальных вызовов.

## Method promotion

Подробно: [02_method_promotion.md](./01_theory/02_method_promotion.md), [05_pointer_embedding.md](./01_theory/05_pointer_embedding.md).

| Встроено в `Outer` | Методы с receiver `T` | Методы с receiver `*T` |
| :--- | :--- | :--- |
| `T` | у `Outer` и `*Outer` | только у `*Outer` (или у addressable `Outer`) |
| `*T` | у `Outer` и `*Outer` | у `Outer` и `*Outer` |

| | Встраивание `T` | Встраивание `*T` |
| :--- | :--- | :--- |
| Хранит | значение (копируется вместе с `Outer`) | указатель (несколько `Outer` делят один `T`) |
| Zero value `Outer{}` | готовый `T` | `nil` → вызов метода, читающего поля, паникует |
| Когда | value object | инфраструктура: клиент, логгер, `*sql.DB`, `*os.File` |

Если встроенный тип реализует интерфейс, внешний — тоже (через promoted методы).

## Перекрытие и конфликты

Подробно: [03_override.md](./01_theory/03_override.md), [04_name_conflicts.md](./01_theory/04_name_conflicts.md).

| Ситуация | Результат |
| :--- | :--- |
| У `Outer` свой метод с тем же именем | перекрывает встроенный; встроенный — `o.Inner.M()` |
| `M` у двух встроенных на одном уровне | `o.M()` — ошибка `ambiguous selector`; явно `o.A.M()` / `o.B.M()` |
| Разрешить конфликт | объявить `func (o Outer) M()` на внешнем типе |
| Одноимённые поля на разной глубине | побеждает ближайший уровень (shadowing): `o.Name` = `o.B.Name`, глубже — `o.B.A.Name` |

```go
package main

import "fmt"

type Animal struct{ Name string }

func (a Animal) Speak() string { return "..." }
func (a Animal) Hello() string { return "я " + a.Name + ": " + a.Speak() }

type Dog struct{ Animal }

func (d Dog) Speak() string { return "гав" } // перекрывает Animal.Speak

func main() {
	d := Dog{Animal{Name: "Бобик"}}
	fmt.Println(d.Speak())        // гав
	fmt.Println(d.Animal.Speak()) // ...
	fmt.Println(d.Hello())        // я Бобик: ...  (нет виртуального вызова)
	fmt.Println(d.Name)           // Бобик (promoted field)
}
```

## Применение

Подробно: [06_use_cases.md](./01_theory/06_use_cases.md).

| Задача | Приём |
| :--- | :--- |
| Декоратор: заменить один метод, остальные делегировать | встроить интерфейс, перекрыть метод, внутри вызвать `r.UserRepo.Save(ctx, u)` |
| Добавить поля/счётчики к типу stdlib | `type FancyFile struct { *os.File; bytesRead int64 }` + свой `Read` |
| Миксины | `type Service struct { Logger; Metrics; Clock }` |
| Набор опций | `type Client struct { HTTPOptions; baseURL string }` |
| Нарастить контракт | композиция интерфейсов |
| «Base»-класс с общими методами | не стоит: лучше интерфейс (`Identifiable`) или явные функции |

## Ловушки

- ⚠️ Ждать полиморфизма: метод встроенного типа вызывает **свою** версию, а не перекрытую во внешнем.
- ⚠️ Встроенный `*T` = `nil` в zero value → nil-panic при первом вызове.
- ⚠️ Встроили `T` с методами на `*T`, а интерфейс присваивают из `Outer` по значению → не реализует.
- ⚠️ Два встроенных типа с одинаковым методом/полем → `ambiguous selector`, линтеры не предупреждают заранее.
- ⚠️ Общий встроенный `*T` у нескольких структур + запись из горутин → нужна синхронизация.
- ⚠️ Встраивание выставляет **все** экспортированные методы встроенного типа в API внешнего; нужна пара — сделайте обычное поле.
