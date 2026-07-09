# Use-cases функций

Где возможности Go-функций проявляются в реальной работе.

## 1. CLI: dispatch команд

Классическая таблица «команда → обработчик»:

```go
cmds := map[string]func(args []string) error{
    "build":  runBuild,
    "clean":  runClean,
    "deploy": runDeploy,
}

h, ok := cmds[os.Args[1]]
if !ok {
    return fmt.Errorf("unknown command: %s", os.Args[1])
}
return h(os.Args[2:])
```

Функция первого класса заменяет километровый `switch`.

## 2. Backend: middleware

HTTP middleware в Go — это функции высшего порядка `func(http.Handler) http.Handler`. Их удобно собирать в цепочку:

```go
h := logging(recovery(auth(coreHandler)))
```

Каждая обёртка добавляет ответственность, не трогая основной обработчик.

## 3. Infra: retry/timeout через декораторы

```go
func WithRetry(fn func() error, attempts int) func() error {
    return func() error {
        var err error
        for i := 0; i < attempts; i++ {
            if err = fn(); err == nil {
                return nil
            }
        }
        return err
    }
}
```

Одна обёртка — одна идея, их можно комбинировать.

## 4. Тесты: подмена поведения

Чтобы подменять внешние вызовы в тестах, делаем их переменными функционального типа:

```go
var now = time.Now // прод

func IsExpired(t time.Time) bool { return now().After(t) }
```

В тесте можно присвоить `now = func() time.Time { return fixed }`.

## 5. Data pipeline через map/filter/reduce

Для трансформаций коллекций функции высшего порядка удобнее, чем императивные циклы с накоплением.

## 6. Инкапсуляция состояния через замыкания

Счётчики, кэши, `once`-инициализация — часто лаконичнее замыканием, чем отдельной структурой.
