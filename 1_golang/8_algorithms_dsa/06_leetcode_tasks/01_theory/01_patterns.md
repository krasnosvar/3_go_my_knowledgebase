# LeetCode patterns: краткий гид

Большинство алгоритмических задач сводится к нескольким повторяющимся паттернам. Знание паттерна = минус 30 минут на задачу.

## Two pointers

Два указателя, движутся навстречу или в одном направлении.

- Reverse string / array.
- Two sum в отсортированном массиве.
- Max water container.
- Remove duplicates in-place.

## Sliding window

Окно фиксированного или переменного размера, скользит по массиву/строке.

- Max sum subarray of size K.
- Longest substring without repeating chars.
- Min window covering set of chars.

```go
left := 0
for right := 0; right < len(s); right++ {
    // расширяем окно [left, right]
    for windowInvalid {
        // сдвигаем left
        left++
    }
    // обновляем ответ
}
```

## Hash map / counting

Map для подсчёта частот, dedup, lookup'а.

- Two sum (unsorted) — O(n) с map.
- Group anagrams.
- Subarray sum equals K.
- First unique character.

## Stack

LIFO для парсинга / истории / undo-логики.

- Balanced parentheses.
- Min stack.
- Daily temperatures (monotonic stack).
- Largest rectangle in histogram.

## BFS / DFS

Графы, деревья, матрицы (грид как граф).

- Number of islands.
- Word ladder.
- Course schedule (cycle in DAG).

## Binary search

В отсортированном массиве или на пространстве ответа («найти минимальное X, удовлетворяющее условию»).

- Search in rotated sorted array.
- Find peak element.
- Min capacity to ship in K days.

## Dynamic programming

Когда подзадача переиспользуется. Top-down (memoization) или bottom-up (table).

- Climbing stairs / Fibonacci.
- House robber.
- Longest common subsequence.
- Coin change.

Всегда формулируй state и transition. Без чёткого state DP не пишется.

## Greedy

На каждом шаге берём «локально лучшее». Работает не всегда — нужна доказуемость.

- Interval scheduling.
- Jump game.
- Buy/sell stock.

## Backtracking

Рекурсивный перебор с откатом.

- Permutations / combinations.
- N-queens.
- Word search.

## Heap (priority queue)

- Top-K elements.
- Merge K sorted lists.
- Median of stream.

## Union-Find

- Number of provinces.
- Accounts merge.
- Redundant connection.

## Подход к задаче

1. Прочитать. Перефразировать своими словами.
2. Прогнать примеры **руками**. Часто видно паттерн.
3. Определить паттерн → применить шаблон.
4. Оценить сложность → принять, если совпадает с лимитами.
5. Кодить от частных случаев к общему. Сначала тривиальные edge cases.
6. Тестировать на собственных примерах.

## Правило

Сила не в знании 500 решений, а в умении быстро выделить паттерн и аккуратно реализовать. Solve, repeat, review.
