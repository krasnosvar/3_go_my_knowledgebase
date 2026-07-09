# break, continue, labels

Управление потоком внутри циклов и `switch`.

## break

Немедленно выходит из ближайшего `for` или `switch`.

```go
for _, v := range nums {
    if v == target {
        break // нашли — дальше не идем
    }
}
```

## continue

Пропускает оставшуюся часть итерации и переходит к следующей.

```go
for _, v := range nums {
    if v%2 != 0 {
        continue // пропускаем нечетные
    }
    process(v)
}
```

## Метки (labels)

Нужны, чтобы `break` или `continue` управляли внешним циклом, а не внутренним:

```go
Outer:
for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
        if grid[i][j] == target {
            found = true
            break Outer // выходим сразу из обоих циклов
        }
    }
}
```

## break внутри select

В `select` (например, внутри цикла) `break` выходит именно из `select`, не из внешнего `for`. Чтобы выйти из цикла, используй label:

```go
Loop:
for {
    select {
    case <-done:
        break Loop
    case v := <-ch:
        handle(v)
    }
}
```

## Что нельзя

- `break` и `continue` не работают в `if` без цикла — это не C.
- `goto` в Go существует, но применяется редко — предпочитай `return`, `break`, функции.
