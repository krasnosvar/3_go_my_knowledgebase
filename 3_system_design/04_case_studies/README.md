# Тема 04. Сквозные кейсы и интервью в Go (Case Studies)

## О чём эта тема (краткий контекст)

Финальный модуль объединяет все строительные блоки курса в целостные архитектуры
реальных систем, которые чаще всего разбирают на System Design собеседованиях и
строят в продакшене. Здесь мы связываем методологию интервью и расчёты ёмкости
(QPS, память, трафик, хранилище) со спецификой рантайма Go и разбираем
реализацию ключевых узлов четырёх классических систем: денежных переводов,
медиа-галереи, сокращателя ссылок (URL Shortener) и ленты новостей (News Feed).

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [01. Фреймворк System Design интервью и расчёты «на салфетке»](../../../5_devops_sre_knowledgebase/14_system_design/07_case_studies_and_interview/01_theory/01_interview_framework_and_estimations.md)
> - [02. Сквозной разбор 4 кейсов: Денежные переводы, Медиа-галерея, URL Shortener и News Feed](../../../5_devops_sre_knowledgebase/14_system_design/07_case_studies_and_interview/01_theory/02_case_studies_transfers_and_media.md)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

**Уже знакомы с темой?** → [CHEATSHEET.md](./CHEATSHEET.md) — опорные числа
рантайма Go для расчётов, формулы QPS/Storage, шаблоны кода для переводов,
медиа, URL Shortener и News Feed, ловушки на интервью.

## Пререквизиты

- Пройдены модули [01_architecture_styles](../01_architecture_styles/),
  [02_performance_and_scaling](../02_performance_and_scaling/) и
  [03_distributed_components](../03_distributed_components/).

## 🗂 Структура раздела

*   **[01_theory/](./01_theory/)** — теория Go-специфики:
    *   [01_estimations_and_go_in_interviews.md](./01_theory/01_estimations_and_go_in_interviews.md)
        — опорные числа стоимости горутин, сокетов, мьютексов и каналов в Go;
        пример расчёта кластера на 1 млн WebSocket-соединений; когда на интервью
        выигрывает Go.
    *   [02_case_studies_in_go.md](./01_theory/02_case_studies_in_go.md) —
        реализация ключевых узлов 4 кейсов на Go: идемпотентные денежные
        переводы без deadlock (`SELECT ... FOR UPDATE` + Outbox), потоковая
        загрузка медиа без OOM, URL Shortener (Base62 + `singleflight`) и
        гибридный fan-out ленты новостей.
*   **[02_examples/](./02_examples/)** — runnable-примеры (только stdlib):
    *   [01_base62_shortener](./02_examples/01_base62_shortener/main.go) —
        рабочий прототип URL Shortener в памяти: счётчик ID + кодирование
        Base62 + статистика переходов.
    *   [02_capacity_estimation](./02_examples/02_capacity_estimation/main.go)
        — калькулятор Back-of-the-Envelope: расчёт среднего и пикового QPS,
        пропускной способности сети, объёма хранилища на N лет и памяти под
        горячий кэш (правило 80/20).
*   **[03_exercises/](./03_exercises/)** — задачи с автопроверкой `PASS`/`FAIL`:
    *   `01_base62` — двустороннее преобразование `uint64 ↔ Base62` (`encode`
        и `decode`).
    *   `02_storage_estimation` — расчёт объёма хранилища в ГБ на заданный
        горизонт лет по суточному притоку записей.
    *   `03_qps_estimation` — расчёт среднего и пикового QPS по DAU, числу
        действий на пользователя и коэффициенту пика.
*   **[04_exercises_answers/](./04_exercises_answers/)** — эталонные решения.

## 🚀 Как проходить тему

1. Сначала прочитайте разбор методологии и 4 сквозных кейсов в каноническом
   разделе DevOps, затем изучите уроки в [01_theory/](./01_theory/).
2. Запустите примеры:
   ```bash
   go run ./02_examples/01_base62_shortener/main.go
   go run ./02_examples/02_capacity_estimation/main.go
   ```
3. Выполните упражнения из [03_exercises/](./03_exercises/):
   ```bash
   go run ./03_exercises/01_base62/main.go
   ```
4. Сверьтесь с эталонными ответами в
   [04_exercises_answers/](./04_exercises_answers/).

## Когда двигаться дальше

Поздравляем! Вы прошли все четыре модуля прикладного System Design для Go.
Используйте шпаргалки разделов для быстрого повторения перед собеседованиями и
при проектировании новых сервисов.
