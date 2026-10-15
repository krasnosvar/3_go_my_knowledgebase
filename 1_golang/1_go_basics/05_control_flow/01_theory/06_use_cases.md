# Use-cases control flow

Практические случаи, где важен правильно выбранный инструмент ветвления или цикла.

## 1. Роутинг команд (switch как table dispatch)

```go
switch cmd {
case "build":
    return runBuild()
case "clean":
    return runClean()
default:
    return fmt.Errorf("unknown command: %s", cmd)
}
```

Плоско, читается сверху вниз, легко расширять.

## 2. State machine

`switch` по текущему состоянию — типовой способ описать конечный автомат:

```go
for {
    switch state {
    case "idle":
        state = "running"
    case "running":
        if done {
            state = "stopped"
        }
    case "stopped":
        return
    }
}
```

## 3. Ранний выход вместо вложенности

`if cond { return }` делает функцию плоской и убирает "лестницу" из `if/else`.

## 4. Поиск в данных с ранним выходом

```go
for _, u := range users {
    if u.ID == id {
        return u, true
    }
}
return User{}, false
```

## 5. Обработка событий через for + select

Комбинация `for { select {} }` — основной паттерн event-loop в Go-сервисах.

## 6. Чистка ресурсов

Ресурс (файл, соединение) закрывают через `defer f.Close()` сразу после успешного открытия — тогда он закроется и при `break`, и при `return`. Но `defer` срабатывает при выходе из **функции**, а не из итерации: в цикле по многим файлам выносите тело итерации в отдельную функцию, иначе все дескрипторы останутся открытыми до конца цикла.
