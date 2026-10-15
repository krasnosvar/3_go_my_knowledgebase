# Функции — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии указаны там, где поведение зависит от них.

## Объявление и передача аргументов

Подробно: [01_basics.md](./01_theory/01_basics.md).

| Задача | Код |
| :--- | :--- |
| Полная сигнатура | `func add(x int, y int) int { return x + y }` |
| Параметры одного типа | `func add(x, y int) int` |
| Без результата / без параметров | `func log(msg string) {}` / `func hello() {}` |
| Изменить внешнюю переменную | передать указатель: `func set(x *int) { *x = 100 }`, вызов `set(&a)` |
| Экспорт | `Add` — публичная, `add` — видна только в пакете |
| Порядок по соглашению | `ctx context.Context` — первым, `error` — последним результатом |

| Что передали | Что копируется | Видны ли правки снаружи |
| :--- | :--- | :--- |
| `int`, `string`, `struct`, массив | всё значение | нет |
| `*T` | адрес | да, через `*p` |
| `[]T` | заголовок (ptr, len, cap) | элементы — да; `append` — нет (новый заголовок локален) |
| `map`, `chan` | указатель на runtime-структуру | да |

## Несколько результатов и variadic

Подробно: [02_multiple_return_and_variadic.md](./01_theory/02_multiple_return_and_variadic.md).

| Задача | Код |
| :--- | :--- |
| Результат + ошибка | `func divide(a, b float64) (float64, error)` |
| Принять результат | `res, err := divide(10, 2)`; лишний — `_, err := ...` |
| Именованные результаты | `func split(sum int) (x, y int) { ...; return }` — стартуют с нулевых значений |
| Variadic | `func sum(nums ...int) int` — внутри `nums` это `[]int` |
| Передать слайс в variadic | `sum(nums...)` |
| Вызов без аргументов | `sum()` → `nums == nil`, `len(nums) == 0` |

## Функции как значения и замыкания

Подробно: [03_first_class_functions.md](./01_theory/03_first_class_functions.md), [04_closures.md](./01_theory/04_closures.md).

| Задача | Код |
| :--- | :--- |
| Присвоить переменной | `greet := func(name string) { fmt.Println("Hello,", name) }` |
| Функциональный тип | `type Op func(int, int) int` |
| Передать как аргумент | `func compute(x, y int, op func(int, int) int) int { return op(x, y) }` |
| Вернуть функцию | `func multiplier(f int) func(int) int { return func(n int) int { return n * f } }` |
| Анонимная, вызвана сразу | `func() { ... }()` |
| Нулевое значение | `var f func()` → `nil`; вызов `f()` → panic |
| Сравнение | только с `nil`; `f1 == f2` не компилируется |

```go
func counter() func() int {
	c := 0                      // захвачена по ссылке → уезжает в кучу
	return func() int { c++; return c }
}

next := counter()
next()
fmt.Println(next()) // 2; новый counter() — своё независимое состояние
```

| Правило замыкания | Следствие |
| :--- | :--- |
| Захват по ссылке, а не копия | изменения видны и внутри, и снаружи |
| Захваченная переменная живёт, пока жива функция | escape analysis переносит её в кучу |
| Несколько замыканий из одного вызова | делят одну переменную (`add` и `reset` над одной `sum`) |
| Параллельное присваивание | `f1, f2 = f2, f1+f2` — сначала вычисляется правая часть |

## Дженерики (кратко)

Подробно: [05_generics.md](./01_theory/05_generics.md); полная тема — [2_oop/05_generics](../../2_oop/05_generics/).

| Задача | Код |
| :--- | :--- |
| Параметр типа | `func Map[T, U any](s []T, f func(T) U) []U` |
| Упорядоченные типы (Go 1.21+) | `func Max[T cmp.Ordered](a, b T) T` (`import "cmp"`) |
| Ключ мапы / `==` | `[K comparable, V any]` |
| Свой constraint | `type Number interface { ~int \| ~int64 \| ~float64 }` |
| Вывод типа / явно | `Max(3, 7)` / `Max[float64](1, 2.5)` |
| Встроенные `min`/`max` (Go 1.21+) | `min(3, 1, 2)` → `1`, `max(2.5, 1)` → `2.5` |
| Generic-тип | `type Stack[T any] struct { items []T }` → `Stack[int]{}` |
| ❌ запрещено | метод со своим параметром типа: `func (r R) M[T any]()` |

## defer, panic, recover

Подробно: [06_defer_panic_recover.md](./01_theory/06_defer_panic_recover.md).

| Правило | Пример |
| :--- | :--- |
| Аргументы вычисляются в строке `defer` | `i := 0; defer fmt.Println(i); i++` → печатает `0` |
| Порядок LIFO | `defer` 1, 2, 3 → выполняются 3, 2, 1 |
| Может менять именованный результат | `defer func() { result++ }(); return 1` → вернёт `2` |
| Выполняется и при panic | `defer`-функции каждого кадра отрабатывают при разворачивании стека |
| `recover()` | работает только в функции, вызванной через `defer`; иначе `nil` |
| Только своя горутина | panic в другой горутине так не поймать |

```go
func safeCall(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r) // panic → error на границе
		}
	}()
	f()
	return nil
}
```

| `defer` НЕ выполнится, если… | Почему |
| :--- | :--- |
| `os.Exit(n)` | процесс завершается немедленно |
| `log.Fatal(...)` | внутри вызывает `os.Exit(1)` |
| `SIGKILL` (`kill -9`) | процесс убит ОС без cleanup |
| бесконечный цикл / deadlock | функция никогда не возвращается |

## Use-cases

Подробно: [07_use_cases.md](./01_theory/07_use_cases.md).

| Задача | Приём |
| :--- | :--- |
| Dispatch команд CLI | `map[string]func(args []string) error` вместо `switch` |
| HTTP middleware | `func(http.Handler) http.Handler`, цепочка `logging(recovery(auth(h)))` |
| Retry-декоратор | `func WithRetry(fn func() error, n int) func() error` |
| Подмена в тестах | `var now = time.Now`; в тесте `now = func() time.Time { return fixed }` |
| Состояние без структуры | замыкание: счётчик, кэш, ленивая инициализация |

## Ловушки

- ⚠️ `append` к слайсу-аргументу не виден вызывающему — верните слайс или передайте `*[]T`.
- ⚠️ `:=` внутри замыкания создаёт новую локальную переменную вместо изменения захваченной.
- ⚠️ До Go 1.22 замыкания в `for` делили одну переменную цикла (все печатали последнее значение); с 1.22 (при `go 1.22+` в `go.mod`) — своя переменная на итерацию.
- ⚠️ `defer` в цикле копит вызовы до выхода из функции (файлы не закрываются) — выносите тело в отдельную функцию.
- ⚠️ `defer f.Close()` до проверки `err` → panic на nil-ресурсе.
- ⚠️ Голый `return` в длинной функции с именованными результатами плохо читается.
- ⚠️ Panic в `main` без `recover`: отложенные вызовы `main` всё равно выполнятся, потом программа упадёт с кодом 2.
- ⚠️ Вызов nil-функции → `panic: runtime error: invalid memory address or nil pointer dereference`.
