# Методы — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.

## Синтаксис

Подробно: [01_basics.md](./01_theory/01_basics.md).

| Задача | Код |
| :--- | :--- |
| Метод с value receiver | `func (r Rectangle) Area() float64 { return r.W * r.H }` |
| Метод с pointer receiver | `func (c *Counter) Inc() { c.N++ }` |
| Вызов | `rect.Area()` |
| Метод на не-struct типе | `type Celsius float64`<br>`func (c Celsius) ToF() float64 { return float64(c)*9/5 + 32 }` |
| Метод для чужого/встроенного типа | нельзя; заводим свой именованный тип: `type IntSet []int` |
| Экспорт | `Area` — публичный, `area` — приватный (как у полей) |
| Method value (receiver зафиксирован) | `f := rect.Area; f()` |
| Method expression (receiver — первый аргумент) | `g := Rectangle.Area; g(rect)` — `func(Rectangle) float64` |
| Method expression для `*T` | `(*Counter).Inc` — `func(*Counter)` |
| Receiver без имени (поля не нужны) | `func (stderrLogger) Log(msg string) {...}` |

Чего нет: перегрузки по типам аргументов, наследования, методов на интерфейсе, методов на типах из чужих пакетов.

## Value или pointer receiver

Подробно: [02_value_pointer.md](./01_theory/02_value_pointer.md).

| Ситуация | Receiver |
| :--- | :--- |
| Метод меняет состояние | `*T` |
| Большая структура (десятки полей) | `*T` |
| В типе `sync.Mutex`, `sync.WaitGroup` (копировать нельзя) | `*T` |
| Маленький неизменяемый value object (Money, Coord, Color) | `T` |
| Хотя бы один метод уже на `*T` | все методы на `*T` |
| Нужен «тот же самый» объект в нескольких местах | `*T` |

## Method set и addressability

Подробно: [03_method_set.md](./01_theory/03_method_set.md).

| Тип | Method set |
| :--- | :--- |
| `T` | только методы с receiver `T` |
| `*T` | методы с receiver `T` **и** `*T` |

| Выражение | Addressable (`v.PtrMethod()` работает) |
| :--- | :--- |
| Локальная переменная | да |
| Поле структуры, элемент слайса, элемент массива-переменной | да |
| Элемент map `m["a"]` | **нет** |
| Результат функции «на лету» `f().M()` | нет (сначала присвоить переменной) |

```go
package main

import "fmt"

type Counter struct{ N int }

func (c *Counter) Inc()      { c.N++ }
func (c Counter) Value() int { return c.N }

func main() {
	var c Counter
	c.Inc()             // авто &c: c addressable
	f := Counter.Value  // method expression: func(Counter) int
	g := (*Counter).Inc // func(*Counter)
	g(&c)
	fmt.Println(c.Value(), f(c)) // 2 2

	m := map[string]Counter{"a": {}}
	// m["a"].Inc()     // ошибка: элемент map не addressable
	v := m["a"]
	v.Inc()
	m["a"] = v
	fmt.Println(m["a"].N) // 1
}
```

## Стиль

Подробно: [04_best_practices.md](./01_theory/04_best_practices.md), [06_use_cases.md](./01_theory/06_use_cases.md).

| Правило | Пример |
| :--- | :--- |
| Короткое одинаковое имя receiver | `func (u *User) Rename(...)`, не `user`, не `this` |
| `error` — последний результат | `func (r *Repo) Get(id int) (User, error)` |
| `context.Context` — первый аргумент для I/O | `func (r *Repo) Get(ctx context.Context, id int) (User, error)` |
| Геттер без побочных эффектов | `func (u *User) Name() string { return u.name }` |
| Красивый вывод в логах | `func (s Status) String() string` (`fmt.Stringer`) |
| Doc-комментарий с имени метода | `// Get loads a user by id.` |
| Value object возвращает новое значение | `func (m Money) Add(o Money) (Money, error)` |
| Fluent builder | методы возвращают `*T`: `NewRequest().WithHeader(k, v).WithTimeout(d)` |
| Утилита без состояния | обычная функция, а не метод |

## Ловушки

Подробно: [05_pitfalls.md](./01_theory/05_pitfalls.md).

- ⚠️ Value receiver в мутаторе: `func (c Counter) Inc() { c.N++ }` меняет копию.
- ⚠️ Смесь `T` и `*T` receiver в одном типе → `T` не реализует интерфейс, путаница «меняется или нет».
- ⚠️ `m["a"].Inc()` с pointer receiver не компилируется → храните `map[string]*T` или «прочитать–изменить–записать».
- ⚠️ `func (s S) Do() { s.mu.Lock() }` копирует mutex → `go vet`: copies lock.
- ⚠️ `var u *User; u.Name()` → nil pointer dereference, если метод читает поля; защита `if u == nil`.
- ⚠️ Внутри `String()` нельзя `fmt.Sprintf("%v", e)` на себе → бесконечная рекурсия.
- ⚠️ Method value: value receiver фиксирует **копию** на момент `f := c.Value`; pointer receiver фиксирует `&c` — `f()` увидит последующие присваивания `c`.
- ⚠️ Поле `Name` и метод `Name()` в одном типе — конфликт имён, не компилируется.
- ⚠️ Два метода `Do(int)` и `Do(string)` — перегрузки нет, ошибка компиляции.
