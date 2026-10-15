# Тема 02. Типы данных (Data Types)

Типы в Go определяют точный формат представления данных в памяти и набор допустимых операций над ними. Раздел разбирает систему типов Go: числовые типы и их размеры, строки/`byte`/`rune` и UTF-8, нулевые значения, явное приведение типов и практические критерии выбора типа под задачу.

## 📋 Пререквизиты

*   Пройдена предыдущая тема: [01_variables_constants](../01_variables_constants/).
*   Справочно (по желанию): [кодировки ASCII/Unicode/UTF-8](../../../other/02_encodings_ascii_unicode_utf8.md) и [как компьютер хранит числа](../../../other/03_number_representation.md) — пригодятся для `byte`/`rune` и переполнения целых.

> ⏩ **Забегание вперёд.** Некоторые примеры и упражнения используют конструкции из следующих тем:
> *   Проверки в упражнениях используют `if` ([05_control_flow](../05_control_flow/)), `fmt.Printf` ([03_fmt_and_output](../03_fmt_and_output/)) и простые функции ([07_functions](../07_functions/)). Пока читайте эту обвязку как готовый код — менять нужно только места с `TODO`.
> *   Примеры используют структуры ([08_structs](../08_structs/)), методы ([2_oop/01_methods](../../2_oop/01_methods/)) и `strconv` ([4_standard_library/01_fmt_strconv_strings](../../4_standard_library/01_fmt_strconv_strings/)).
> *   Упражнение `08_composite_preview` — краткое превью слайсов ([09_slices](../09_slices/)) и мап ([11_maps](../11_maps/)).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория:
    *   [01_data_types_basics.md](./01_theory/01_data_types_basics.md) — вводный обзор системы типов: справочная таблица (размеры, нулевые значения, сравнимость, передача), категории типов (числовые, bool, строки/byte/rune, составные и ссылочные), ASCII-схемы представления в памяти (string, слайс, map, interface), именованные типы и псевдонимы (Named Types vs Type Aliases).
    *   [02_numeric_types.md](./01_theory/02_numeric_types.md) — числовые типы: целые со знаком и без, специальные (`byte`, `rune`, `uintptr`), числа с плавающей точкой и комплексные, переполнение (wraparound), сравнение float через epsilon, размеры типов.
    *   [03_string_byte_rune.md](./01_theory/03_string_byte_rune.md) — строки: внутреннее устройство, неизменяемость, UTF-8, разница `byte` и `rune`, итерация через `range`, конверсии `string` ↔ `[]byte`/`[]rune` и их стоимость.
    *   [04_zero_values.md](./01_theory/04_zero_values.md) — нулевые значения всех типов, нюансы `nil` map/slice/chan, принцип "useful zero value".
    *   [05_conversions.md](./01_theory/05_conversions.md) — явное приведение типов: числовые конверсии и усечение, `string` ↔ `[]byte`/`[]rune`, ловушка `string(number)`, custom-типы, краткое упоминание `unsafe.Pointer`.
    *   [06_use_cases.md](./01_theory/06_use_cases.md) — практический выбор типа: целые под задачу, работа со строками, контейнеры, `nil`-семантика, границы арифметики, comparable-ключи map, конверсии на горячем пути.
*   **[02_examples/](./02_examples/)** — практические примеры:
    *   [01_cli_utilities](./02_examples/01_cli_utilities/main.go) — использование типов `uint16`, `bool` и `string` в CLI-конфигурациях.
    *   [02_system_infrastructure](./02_examples/02_system_infrastructure/main.go) — расчёт процентов ресурсов через явное приведение к `float64` и безопасное сравнение float с epsilon.
    *   [03_product_backend](./02_examples/03_product_backend/main.go) — финансовые расчёты с кастомным типом `Money` на базе `int64` для защиты от погрешностей округления float.
*   **[03_exercises/](./03_exercises/)** — интерактивные задачи для самопроверки с автопроверками `PASS`/`FAIL`:
    *   `01_numeric` — сложение и деление с защитой от деления на ноль.
    *   `02_bool_logic` — проверка прав доступа пользователя (логическое И).
    *   `03_string_bytes` — безопасное получение первого байта строки (защита от panic).
    *   `04_rune_utf8` — извлечение первой руны UTF-8 строки с использованием `range`.
    *   `05_type_conversion` — вычисление среднего значения через явное приведение типов.
    *   `06_zero_values` — проверка параметров конфигурации (`bool`, `int`, `string`) на нулевые значения.
    *   `07_alias_types` — использование Named Type и Type Alias.
    *   `08_composite_preview` — инициализация простых слайсов и мап.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения с объяснением "почему сделано именно так".

## 🚀 Как проходить тему

1. Прочитайте вводный обзор: [01_theory/01_data_types_basics.md](./01_theory/01_data_types_basics.md), затем детальные файлы `02`–`06`.
2. Запустите и изучите примеры из папки `02_examples/` через `go run`:
   ```bash
   go run ./02_examples/02_system_infrastructure/main.go
   ```
3. Выполните задачи в `03_exercises/`. Запустите встроенный тест для самопроверки:
   ```bash
   go run ./03_exercises/01_numeric/main.go
   ```
   Добейтесь вывода `PASS` для всех тестов.
4. Сверьтесь с эталонными ответами в `04_exercises_answers/`.

## ➡️ Что дальше

Когда все упражнения темы печатают `PASS`, переходите к теме [03_fmt_and_output](../03_fmt_and_output/).
