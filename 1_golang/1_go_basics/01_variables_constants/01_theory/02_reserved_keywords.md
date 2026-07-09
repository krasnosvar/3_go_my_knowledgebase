# Зарезервированные слова Go

В Go 25 зарезервированных ключевых слов. Их нельзя использовать в качестве идентификаторов (имён переменных, функций, типов и т.д.).

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

Краткое назначение:

- **Объявления:** `var`, `const`, `func`, `type`, `struct`, `interface`, `map`, `chan`, `import`, `package`.
- **Управление потоком:** `if`, `else`, `switch`, `case`, `default`, `fallthrough`, `for`, `range`, `break`, `continue`, `goto`, `return`.
- **Конкурентность и отложенное выполнение:** `go`, `defer`, `select`.

Помимо ключевых слов, в языке есть предопределённые идентификаторы (`true`, `false`, `nil`, `int`, `len`, `make` и др.) — их формально можно переопределить, но делать этого не стоит.

Полный список — в спецификации: [go.dev/ref/spec#Keywords](https://go.dev/ref/spec#Keywords).
