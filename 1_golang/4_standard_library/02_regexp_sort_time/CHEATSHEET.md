# regexp, sort, time — шпаргалка

Подробности и объяснения — в [01_theory/](./01_theory/). ⚠️ — частая ловушка.
Версии указаны там, где API появилось недавно.

## regexp

Подробно: [01_regexp.md](./01_theory/01_regexp.md).

| Задача | Код |
| :--- | :--- |
| Скомпилировать константный паттерн (один раз, на уровне пакета) | ``var re = regexp.MustCompile(`^\d{3}-\d{4}$`)`` — panic при ошибке |
| Паттерн из пользовательского ввода | `re, err := regexp.Compile(p)` |
| Совпадает? | `re.MatchString(s)` |
| Первое / все вхождения | `re.FindString(s)` (`""` — нет) / `re.FindAllString(s, -1)` |
| Позиция | `re.FindStringIndex(s)` → `[start, end]` |
| Группы | `m := re.FindStringSubmatch(s)` → `m[0]` целиком, `m[1]`… группы |
| Все совпадения с группами | `re.FindAllStringSubmatch(s, -1)` |
| Именованная группа | `(?P<user>\w+)` → `m[re.SubexpIndex("user")]` |
| Замена | `re.ReplaceAllString(s, "X")`; ссылка на группу — `"${domain}"` |
| Замена функцией | `re.ReplaceAllStringFunc(s, strings.ToUpper)` |
| Split | `re.Split(s, -1)` |

| Синтаксис RE2 | Значение |
| :--- | :--- |
| `.` | любой символ, кроме `\n` (с `(?s)` — и `\n`) |
| `\d` `\w` `\s` | цифра, word, пробельный |
| `^` `$` | начало/конец текста (с `(?m)` — строки) |
| `*?` `+?` `??` | ленивые квантификаторы |
| `(?i)` | без учёта регистра |
| look-behind, backreferences | ⚠️ нет — `Compile` вернёт ошибку; зато время всегда O(n) |

| Задача | Вместо regexp |
| :--- | :--- |
| Точное совпадение / префикс | `strings.Contains` / `HasPrefix` |
| Разбор по полям | `strings.Fields`, `bufio.Scanner` |
| JSON / XML / HTML | `encoding/json`, `encoding/xml`, парсер — не regexp |

## sort и slices

Подробно: [02_sort.md](./01_theory/02_sort.md).

| Задача | Go 1.21+ (`slices`, `cmp`) | Классический `sort` |
| :--- | :--- | :--- |
| Сортировать `[]int`/`[]string` | `slices.Sort(s)` | `sort.Ints(s)` / `sort.Strings(s)` / `sort.Float64s` |
| По ключу | `slices.SortFunc(us, func(a, b User) int { return cmp.Compare(a.Age, b.Age) })` | `sort.Slice(us, func(i, j int) bool { return us[i].Age < us[j].Age })` |
| Стабильно | `slices.SortStableFunc(us, cmp)` | `sort.SliceStable` / `sort.Stable(x)` |
| По двум ключам | `cmp.Or(cmp.Compare(a.Age, b.Age), cmp.Compare(a.Name, b.Name))` (`cmp.Or` — 1.22+) | составной `Less` |
| По убыванию | `cmp.Compare(b.X, a.X)` (аргументы наоборот) | `sort.Sort(sort.Reverse(sort.IntSlice(s)))` |
| Развернуть | `slices.Reverse(s)` | — |
| Проверить сортировку | `slices.IsSorted(s)` | `sort.IntsAreSorted(s)` |
| Бинарный поиск | `i, found := slices.BinarySearch(s, 42)`; свой — `BinarySearchFunc` | `sort.SearchInts(s, 42)`, `sort.Search(n, pred)` |
| Свой тип с порядком | — | `sort.Interface`: `Len`, `Less(i, j)`, `Swap(i, j)` → `sort.Sort(ByAge(us))` |

| Функция | Стабильна? |
| :--- | :--- |
| `slices.Sort`, `slices.SortFunc`, `sort.Sort`, `sort.Slice` | нет |
| `slices.SortStableFunc`, `sort.Stable`, `sort.SliceStable` | да — равные сохраняют исходный порядок |

## time

Подробно: [03_time.md](./01_theory/03_time.md).

| Задача | Код |
| :--- | :--- |
| Сейчас / сдвиг / разница | `now := time.Now()`; `now.Add(10*time.Second)`; `later.Sub(now)` |
| Сравнить моменты | `t1.Before(t2)`, `After`, `t1.Equal(t2)` (не `==`) |
| Unix-время | `now.Unix()`, `now.UnixMilli()` |
| Сколько прошло / осталось | `time.Since(t)` / `time.Until(t)` (монотонные часы) |
| Длительность | `2*time.Second + 500*time.Millisecond` → `2.5s`; `d.Milliseconds()` → `2500` |
| Конкретная дата | `time.Date(2026, 4, 17, 15, 4, 5, 0, time.UTC)` |
| Форматировать | `t.Format(time.RFC3339)` → `2026-04-17T15:04:05Z`; `t.Format("02.01.2006 15:04")` |
| Парсить | `time.Parse("2006-01-02", "2026-04-17")` (UTC); в поясе — `time.ParseInLocation(layout, s, loc)` |
| Готовые layout | `time.RFC3339`, `RFC1123`, `DateTime`, `DateOnly`, `TimeOnly` (последние три — 1.20+) |
| «Не задано»? | `t.IsZero()` (`time.Time{}` — 1 января года 1, не Unix-эпоха) |
| Пауза | `time.Sleep(d)` |
| Одноразовый таймер | `t := time.NewTimer(d); defer t.Stop(); <-t.C` |
| Периодически | `tk := time.NewTicker(d); defer tk.Stop()` + `select { case <-tk.C: case <-ctx.Done(): }` |
| Таймаут в `select` | `case <-time.After(2*time.Second):` (в hot-цикле — `NewTimer` + `Reset`) |

| Компонент layout | Значение |
| :--- | :--- |
| `2006` / `06` | год |
| `01` / `Jan` | месяц |
| `02` | день |
| `15` / `03` + `PM` | час 24ч / 12ч |
| `04` / `05` | минуты / секунды |
| `MST` / `-07:00` | пояс |

## Ловушки

- ⚠️ `regexp.MustCompile` на пользовательском вводе — panic; `Compile` + проверка `err`.
- ⚠️ Компиляция regexp в цикле/обработчике — дорого; храните скомпилированный.
- ⚠️ Не экранированы `.` `*` `(` — используйте raw-строки `` `...` `` и `regexp.QuoteMeta`.
- ⚠️ `sort.Slice` медленнее `slices.SortFunc` (reflection-swap); на миллионах элементов заметно.
- ⚠️ Сортировка по двум ключам двумя проходами без stable — порядок первого ключа теряется.
- ⚠️ Сортировать слайс, который читает другая горутина, — гонка.
- ⚠️ Layout `2002-01-02` вместо `2006-01-02` — неверный разбор без ошибки компиляции.
- ⚠️ `t1 == t2` сравнивает и монотонную часть, и пояс; для моментов — `t1.Equal(t2)`.
- ⚠️ Длительность через разницу `Unix()`/`UnixMilli()` — это настенные часы, ломается при NTP/смене времени; `time.Since`/`Sub` берут монотонную часть.
- ⚠️ Не остановленный `Ticker` в долгоживущем коде продолжает тикать; до Go 1.23 не остановленные `Timer`/`Ticker` ещё и не собирались GC — всегда `defer Stop()`.
