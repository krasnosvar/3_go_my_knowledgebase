# Docker Compose: Kafka + Kafka UI

## Что поднимается

`docker-compose.yaml` запускает учебный стенд из одного брокера:

- **zookeeper** (`confluentinc/cp-zookeeper:7.4.4`) — хранилище метаданных кластера, порт `2181` внутри сети compose;
- **kafka** (`confluentinc/cp-kafka:7.4.4`, это Kafka 3.4) — брокер:
  - `localhost:9092` — listener для клиентов с хоста (Go-программы из [../exercises/](../exercises/));
  - `kafka:29092` — listener внутри сети compose (для Kafka UI и init-контейнера);
- **kafka-ui** (`provectuslabs/kafka-ui`) — веб-интерфейс на `http://localhost:8080`;
- **init-kafka** — одноразовый контейнер: ждёт готовности брокера и создаёт топик `test-topic` (1 партиция, replication factor 1).

> Стенд работает в режиме ZooKeeper — так устроены образы Confluent 7.4. В Kafka 4.0 ZooKeeper удалён, поддерживается только режим KRaft (метаданные хранят сами контроллеры Kafka). Для изучения API продюсера/консьюмера разницы нет; про KRaft см. [../interview_questions/theory.md](../interview_questions/theory.md).

## Запуск и остановка

```bash
docker compose up -d
docker compose ps
docker compose logs init-kafka   # должно быть "Successfully created the following topics"
```

Остановка (с удалением данных):

```bash
docker compose down -v
```

## Проверка из командной строки

Список топиков:

```bash
docker compose exec kafka kafka-topics --bootstrap-server kafka:29092 --list
```

Описание топика (партиции, лидер, ISR):

```bash
docker compose exec kafka kafka-topics --bootstrap-server kafka:29092 --describe --topic test-topic
```

Записать сообщения (каждая строка — сообщение, выход — Ctrl+D):

```bash
docker compose exec kafka kafka-console-producer --bootstrap-server kafka:29092 --topic test-topic
```

Прочитать сообщения с начала:

```bash
docker compose exec kafka kafka-console-consumer --bootstrap-server kafka:29092 --topic test-topic --from-beginning --timeout-ms 5000
```

Состояние consumer group из упражнения:

```bash
docker compose exec kafka kafka-consumer-groups --bootstrap-server kafka:29092 --describe --group group-1
```

## UI

Kafka UI: `http://localhost:8080` — кластер `local` уже настроен, там видны топики, сообщения и consumer groups.
