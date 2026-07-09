# Redis: учебный стенд + материалы

Этот каталог — учебный набор по Redis:

- **теория/конспект** (примитивы, персистентность, репликация/cluster, eviction)
- **практика на Docker Compose**
- **исходные вопросы/ответы** (в `interview_questions/`)

## Структура

- `docker_compose/`
  - `docker-compose.yaml` — поднимает Redis + RedisInsight
  - `README.md` — как запускать, подключаться, проверять статусы
- `exercises/`
  - `ex1/` — простое упражнение (SET/GET) на Go
- `interview_questions/`
  - исходные материалы (вопросы/ответы/задачи)

## Быстрые ссылки

- Compose: `docker_compose/docker-compose.yaml`
- Q&A: `interview_questions/`
