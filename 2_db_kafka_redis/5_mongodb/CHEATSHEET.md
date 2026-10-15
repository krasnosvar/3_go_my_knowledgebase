# MongoDB — шпаргалка

Подробности и объяснения — в [interview_questions/theory.md](./interview_questions/theory.md)
и [basic_mongodb_questions.md](./interview_questions/basic_mongodb_questions.md). ⚠️ — частая ловушка.
Команды `mongosh` не запускались — проверяйте на стенде из [docker_compose/](./docker_compose/README.md).
Go-код — драйвер `go.mongodb.org/mongo-driver` v1 (в упражнениях v1.17).

## Модель данных

Подробно: [theory.md → Документы и Коллекции](./interview_questions/theory.md#документы-и-коллекции),
[theory.md → BSON vs JSON](./interview_questions/theory.md#bson-vs-json).

| MongoDB | SQL |
| :--- | :--- |
| Database | Database |
| Collection | Table |
| Document (BSON) | Row |
| Field | Column |
| `_id` (по умолчанию `ObjectId`) | Primary key |
| embedding / `$lookup` | JOIN |

| Факт | Детали |
| :--- | :--- |
| BSON | бинарный JSON: быстрее парсинг, доп. типы `ObjectId`, `ISODate`, `BinData`, `NumberInt`, `NumberLong`, `NumberDecimal`, regex, `Timestamp` |
| Лимит документа | 16 MB |
| `ObjectId` | 12 байт: 4 — timestamp (сек), 5 — случайное значение, 3 — счётчик; уникален без координатора, сортируется по времени создания |
| Схема | гибкая: документы одной коллекции могут отличаться |

| Брать MongoDB | Не брать |
| :--- | :--- |
| часто меняющаяся схема, иерархические JSON-данные | много сложных JOIN |
| каталоги с разными атрибутами, CMS | строгая консистентность и сложные транзакции |
| нужен горизонтальный масштаб | финансовые системы, данные с чёткими связями → SQL |

## CRUD

Подробно: [basic_mongodb_questions.md → вопрос 4](./interview_questions/basic_mongodb_questions.md#4-какие-основные-операции-crud-вы-знаете-в-mongodb).

| Задача | mongosh | Go |
| :--- | :--- | :--- |
| Подключиться | `docker exec -it mongodb mongosh -u admin -p admin --authenticationDatabase admin` | `mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))` |
| Коллекция | `use test_db` → `db.users` | `client.Database("test_db").Collection("users")` |
| Вставить | `db.users.insertOne({name: "John", age: 30})` / `insertMany([...])` | `coll.InsertOne(ctx, bson.D{{"name", "John"}, {"age", 30}})` / `InsertMany` |
| Найти один | `db.users.findOne({name: "John"})` | `coll.FindOne(ctx, bson.M{"name": "John"}).Decode(&u)` |
| Найти много | `db.users.find({age: {$gt: 18}})` | `cur, err := coll.Find(ctx, filter)` + `defer cur.Close(ctx)` |
| Проекция | `db.users.find({}, {name: 1, email: 1, _id: 0})` | `options.Find().SetProjection(bson.M{"name": 1, "_id": 0})` |
| Обновить | `db.users.updateOne({name: "John"}, {$set: {age: 31}})` / `updateMany` / `replaceOne` | `coll.UpdateOne(ctx, filter, bson.D{{"$set", bson.D{{"age", 31}}}})` |
| Атомарный инкремент | `{$inc: {balance: -100}}` | `bson.M{"$inc": bson.M{"balance": -100}}` |
| Удалить | `db.users.deleteOne({name: "John"})` / `deleteMany` | `coll.DeleteOne(ctx, filter)` / `DeleteMany` |
| Документа нет | `findOne` → `null` | `errors.Is(err, mongo.ErrNoDocuments)` |

| Go-тип | Что это | Когда |
| :--- | :--- | :--- |
| `bson.M` | `map` — порядок полей не сохраняется | фильтры, где порядок не важен |
| `bson.D` | срез пар — упорядоченный | команды, `$sort`, составные ключи индекса, pipeline |
| `bson.A` | массив | `$in`, списки значений |
| `primitive.ObjectID` | тип `_id` | поля `bson:"_id"` в структурах |

Курсор: `for cur.Next(ctx) { cur.Decode(&doc) }`, затем обязательно `cur.Err()`; `defer cur.Close(ctx)`.

## Индексы

Подробно: [theory.md → Индексы в MongoDB](./interview_questions/theory.md#индексы-в-mongodb).

| Тип | Создание | Для чего |
| :--- | :--- | :--- |
| Single field | `db.users.createIndex({email: 1})` (`-1` — убывание) | равенство, диапазон, сортировка по полю |
| Compound | `db.orders.createIndex({user_id: 1, created_at: -1})` | фильтр по префиксу + сортировка |
| Multikey | `db.articles.createIndex({tags: 1})` на массиве | запись в индексе на каждый элемент; `{tags: "x"}`, `$in` |
| Text | `db.articles.createIndex({title: "text", content: "text"})` | `{$text: {$search: "mongodb"}}`, релевантность `{$meta: "textScore"}` |
| Unique | `createIndex({email: 1}, {unique: true})` | дубликат → `duplicate key` (Go: `mongo.IsDuplicateKeyError(err)`) |
| TTL | `createIndex({created_at: 1}, {expireAfterSeconds: 3600})` | автоудаление старых документов |
| Geospatial | `createIndex({loc: "2dsphere"})` | геозапросы |
| Hashed | `{user_id: "hashed"}` | hashed shard key |
| Go | `coll.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{"email", 1}}, Options: options.Index().SetUnique(true)})` | — |

| Правило | Суть |
| :--- | :--- |
| ESR | порядок полей в compound: Equality → Sort → Range |
| Префикс | индекс `{a, b}` обслуживает `{a}` и `{a, b}`, но не `{b}` |
| Covered query | фильтр и проекция только по полям индекса, `_id: 0` → документы не читаются |
| Анализ | `db.c.find(q).explain("executionStats")`: `totalKeysExamined`, `totalDocsExamined`, `executionTimeMillis`; `COLLSCAN` = нет индекса |
| Цена | каждый индекс замедляет запись — удаляйте неиспользуемые |

## Aggregation pipeline

Подробно: [theory.md → Aggregation Pipeline](./interview_questions/theory.md#aggregation-pipeline).

| Стадия | SQL-аналог | Пример |
| :--- | :--- | :--- |
| `$match` | `WHERE` | `{$match: {status: "completed"}}` — ставить как можно раньше |
| `$group` | `GROUP BY` | `{$group: {_id: "$category", total: {$sum: "$amount"}, count: {$sum: 1}, avg: {$avg: "$amount"}}}` |
| `$sort` | `ORDER BY` | `{$sort: {total: -1}}` |
| `$limit` | `LIMIT` | `{$limit: 5}` |
| `$project` | `SELECT` полей | `{$project: {name: 1, _id: 0, full: {$concat: ["$first", " ", "$last"]}}}` |
| `$lookup` | `JOIN` | `{$lookup: {from: "users", localField: "user_id", foreignField: "_id", as: "user"}}` → массив |
| Go | — | `coll.Aggregate(ctx, mongo.Pipeline{{{"$match", bson.D{{"status", "active"}}}}, ...})` |

## Embedding vs referencing

Подробно: [theory.md → Embedding vs Referencing](./interview_questions/theory.md#embedding-vs-referencing).

| | Embedding | Referencing |
| :--- | :--- | :--- |
| Связь | one-to-few | one-to-many, many-to-many |
| Рост данных | ограничен | может расти неограниченно |
| Чтение | один запрос | несколько запросов или `$lookup` |
| Атомарность | изменение одного документа атомарно | между документами — только транзакция |
| Примеры | адрес, атрибуты товара, немного комментариев | заказы пользователя, много комментариев, авторы книг |

Гибрид: ссылка + снимок нужных полей (`user_id` + `user_snapshot{name,email}`); паттерны Subset, Bucket, Outlier, Extended Reference.

## Replica set

Подробно: [theory.md → Replica Set](./interview_questions/theory.md#replica-set).

| Понятие | Суть |
| :--- | :--- |
| Состав | PRIMARY (запись) + SECONDARY (копии через oplog); минимум 3 ноды (PSS) |
| Failover | выборы нового PRIMARY за ~10–30 c, драйвер переподключается сам |
| URI | `mongodb://h1:27017,h2:27017,h3:27017/?replicaSet=rs0` |
| Change Streams | `coll.Watch(ctx, pipeline)` — поток изменений на основе oplog (нужен replica set) |

| Read preference | Откуда читать | Цена |
| :--- | :--- | :--- |
| `primary` (по умолчанию) | только PRIMARY | вся нагрузка на PRIMARY |
| `primaryPreferred` | PRIMARY, при недоступности — SECONDARY | — |
| `secondary` | только SECONDARY | возможны устаревшие данные |
| `secondaryPreferred` | SECONDARY, иначе PRIMARY | устаревшие данные, разгрузка PRIMARY |
| `nearest` | минимальная сетевая задержка | устаревшие данные |

| Write concern | Подтверждение | Риск |
| :--- | :--- | :--- |
| `w: 1` (`writeconcern.W1()`) | только PRIMARY | потеря при падении PRIMARY до репликации |
| `w: "majority"` (`writeconcern.Majority()`) | большинство нод | медленнее; рекомендуется |

## Sharding

Подробно: [theory.md → Sharding](./interview_questions/theory.md#sharding).

| Компонент | Роль |
| :--- | :--- |
| `mongos` | роутер запросов (можно несколько) |
| Shard | данные; каждый shard — replica set |
| Config servers | метаданные: какие chunks где (replica set) |
| Команды | `sh.enableSharding("mydb")`, `sh.shardCollection("mydb.users", {user_id: 1})` или `{user_id: "hashed"}` |

| Shard key | Хорошо | Плохо |
| :--- | :--- | :--- |
| Кардинальность | высокая (`user_id`, `email`) | низкая (`country` ~200 значений) |
| Распределение | равномерное (`hashed`, `{region, user_id}`) | монотонный рост (`timestamp`, `_id`) → все вставки в один shard |
| Запросы | ключ есть в частых фильтрах → targeted query | ключа нет в фильтре → broadcast на все шарды |

| Стратегия | Плюс | Минус |
| :--- | :--- | :--- |
| Range | быстрые range-запросы по ключу | hotspot при монотонном ключе |
| Hashed | равномерно, без hotspot | range-запросы — broadcast |

## Транзакции

Подробно: [theory.md → Транзакции](./interview_questions/theory.md#транзакции).

| Факт | Детали |
| :--- | :--- |
| Без транзакции | атомарна операция над одним документом |
| Multi-document | 4.0+ в replica set, 4.2+ в sharded cluster; на одиночном `mongod` без replica set не работают |
| API Go | `session, _ := client.StartSession()`; `defer session.EndSession(ctx)`; `session.WithTransaction(ctx, fn)` или `StartTransaction` / `CommitTransaction` / `AbortTransaction` внутри `mongo.WithSession` |
| Лимит | 60 c на транзакцию по умолчанию |
| Совет | проектировать схему так, чтобы хватало атомарности одного документа |

## Go: клиент и эксплуатация

Подробно: [basic_mongodb_questions.md → вопрос 13](./interview_questions/basic_mongodb_questions.md#13-как-реализовать-connection-pooling-в-go-приложении-с-mongodb).

| Задача | Код |
| :--- | :--- |
| Один клиент на приложение (пул внутри) | `SetMaxPoolSize(100)`, `SetMinPoolSize(10)`, `SetMaxConnIdleTime(30*time.Second)` |
| Таймаут выбора сервера | `SetServerSelectionTimeout(5*time.Second)` |
| Health check | `client.Ping(ctx, readpref.Primary())` |
| Завершение | `defer client.Disconnect(ctx)` |
| Классификация ошибок | `mongo.IsDuplicateKeyError`, `mongo.IsTimeout`, `mongo.IsNetworkError` |
| Мониторинг команд | `options.Client().SetMonitor(&event.CommandMonitor{Started: ..., Succeeded: ...})` |
| Безопасность | SCRAM-SHA-256, RBAC, TLS (`SetTLSConfig`), шифрование at-rest, закрытая сеть |

## Ловушки

- ⚠️ Перенос SQL-схемы 1:1 → море `$lookup`; моделируйте под запросы.
- ⚠️ Неограниченно растущий массив внутри документа → упрётесь в 16 MB и медленные обновления.
- ⚠️ `bson.M` там, где важен порядок (`$sort` по нескольким полям, ключи индекса) → используйте `bson.D`.
- ⚠️ Не проверили `cur.Err()` после цикла `Next` → ошибка курсора теряется; не закрыли курсор → утечка.
- ⚠️ Новый `mongo.Client` на каждый запрос → нет пула, лишние соединения.
- ⚠️ Составной индекс не в порядке ESR или запрос без префикса индекса → индекс не используется полностью.
- ⚠️ Монотонный shard key → один «горячий» shard; сменить ключ дорого (refine — 4.4+, resharding — 5.0+).
- ⚠️ Чтение с SECONDARY → возможны устаревшие данные.
- ⚠️ `w: 1` → подтверждённая запись может пропасть при failover.
- ⚠️ Транзакции на одиночном инстансе без replica set (как на учебном стенде) → ошибка.
- ⚠️ Без индекса запрос делает `COLLSCAN` (O(n)) — проверяйте `explain()`.
