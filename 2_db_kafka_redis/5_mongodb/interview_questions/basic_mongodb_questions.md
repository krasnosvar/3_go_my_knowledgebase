# ТОП-20 вопросов по MongoDB для Go-разработчика бэкенда

## Простые вопросы (Junior уровень)

### 1. Что такое MongoDB и для каких задач она используется?
**Что ожидают услышать:**
- NoSQL документо-ориентированная БД
- Хранение данных в формате BSON (бинарный JSON)
- Гибкая схема данных
- Горизонтальное масштабирование
- Примеры использования: каталоги продуктов, контент-менеджмент, real-time аналитика

### 2. Как подключиться к MongoDB из Go приложения?
**Что ожидают услышать:**
- Использование официального драйвера `go.mongodb.org/mongo-driver`
- Пример кода:
```go
import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
if err != nil {
    log.Fatal(err)
}
defer client.Disconnect(ctx)

database := client.Database("mydb")
collection := database.Collection("users")
```

### 3. Что такое коллекция и документ в MongoDB?
**Что ожидают услышать:**
- Документ - запись в формате BSON (аналог строки в SQL)
- Коллекция - группа документов (аналог таблицы в SQL)
- Документы в коллекции могут иметь разную структуру
- _id - уникальный идентификатор документа (ObjectID)

### 4. Какие основные операции CRUD вы знаете в MongoDB?
**Что ожидают услышать:**
- Create: InsertOne(), InsertMany()
- Read: FindOne(), Find()
- Update: UpdateOne(), UpdateMany(), ReplaceOne()
- Delete: DeleteOne(), DeleteMany()
- Пример:
```go
// Insert
result, err := collection.InsertOne(ctx, bson.D{{"name", "John"}, {"age", 30}})

// Find
var user User
err = collection.FindOne(ctx, bson.M{"name": "John"}).Decode(&user)

// Update
update := bson.D{{"$set", bson.D{{"age", 31}}}}
collection.UpdateOne(ctx, bson.M{"name": "John"}, update)

// Delete
collection.DeleteOne(ctx, bson.M{"name": "John"})
```

### 5. В чем разница между bson.M, bson.D и bson.A в Go драйвере?
**Что ожидают услышать:**
- bson.M - map (неупорядоченный) `bson.M{"name": "John"}`
- bson.D - slice пар (упорядоченный) `bson.D{{"name", "John"}, {"age", 30}}`
- bson.A - slice/array `bson.A{"value1", "value2"}`
- bson.D используется когда важен порядок полей (например, в $set операциях)

## Средние вопросы (Middle уровень)

### 6. Объясните концепцию индексов в MongoDB. Какие типы индексов вы знаете?
**Что ожидают услышать:**
- Single field index (по одному полю)
- Compound index (по нескольким полям)
- Multikey index (для массивов)
- Text index (полнотекстовый поиск)
- Geospatial index (геопространственный)
- TTL index (автоматическое удаление)
- Пример создания:
```go
indexModel := mongo.IndexModel{
    Keys: bson.D{{"email", 1}},
    Options: options.Index().SetUnique(true),
}
collection.Indexes().CreateOne(ctx, indexModel)
```

### 7. Как работать с транзакциями в MongoDB из Go?
**Что ожидают услышать:**
- Поддержка ACID транзакций с версии 4.0
- Требуется replica set
- Использование сессий:
```go
session, err := client.StartSession()
defer session.EndSession(ctx)

err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
    if err := session.StartTransaction(); err != nil {
        return err
    }
    
    // Операции с БД
    _, err := collection.InsertOne(sc, document1)
    if err != nil {
        session.AbortTransaction(sc)
        return err
    }
    
    _, err = collection.UpdateOne(sc, filter, update)
    if err != nil {
        session.AbortTransaction(sc)
        return err
    }
    
    return session.CommitTransaction(sc)
})
```

### 8. Что такое aggregation pipeline и для чего он используется?
**Что ожидают услышать:**
- Последовательность стадий обработки данных
- Основные стадии: $match, $group, $project, $sort, $limit, $lookup
- Используется для сложных аналитических запросов
- Пример:
```go
pipeline := mongo.Pipeline{
    {{"$match", bson.D{{"status", "active"}}}},
    {{"$group", bson.D{
        {"_id", "$category"},
        {"total", bson.D{{"$sum", "$amount"}}},
        {"count", bson.D{{"$sum", 1}}},
    }}},
    {{"$sort", bson.D{{"total", -1}}}},
}

cursor, err := collection.Aggregate(ctx, pipeline)
```

### 9. Как правильно обрабатывать большие результаты запросов (курсоры)?
**Что ожидают услышать:**
- Использование курсоров вместо загрузки всех данных
- Итерация с помощью Next()
- Обязательное закрытие курсора (defer)
- Пример:
```go
cursor, err := collection.Find(ctx, bson.M{})
if err != nil {
    return err
}
defer cursor.Close(ctx)

for cursor.Next(ctx) {
    var doc bson.M
    if err := cursor.Decode(&doc); err != nil {
        return err
    }
    // Обработка документа
}

if err := cursor.Err(); err != nil {
    return err
}
```

### 10. Что такое replica set в MongoDB и как с ним работать?
**Что ожидают услышать:**
- Группа MongoDB инстансов с репликацией данных
- Primary (запись) + Secondary (чтение/резерв)
- Автоматический failover при падении Primary
- Read Preference для управления чтением
- Пример конфигурации:
```go
clientOpts := options.Client().
    ApplyURI("mongodb://host1:27017,host2:27017,host3:27017/?replicaSet=myReplicaSet").
    SetReadPreference(readpref.SecondaryPreferred())
```

## Продвинутые вопросы (Senior уровень)

### 11. Объясните концепцию sharding в MongoDB. Как выбрать shard key?
**Что ожидают услышать:**
- Горизонтальное распределение данных между серверами
- Shard key - поле для распределения документов
- Критерии выбора shard key:
  - Высокая кардинальность
  - Равномерное распределение
  - Поддержка частых запросов
- Range-based vs Hash-based sharding
- Shard key можно уточнить (`refineCollectionShardKey`, MongoDB 4.4+: добавить поля-суффиксы) и сменить целиком (`reshardCollection`, MongoDB 5.0+)
- Пример плохих ключей: timestamp (монотонный рост), country (низкая кардинальность)
- Пример хороших ключей: user_id, hashed email

### 12. Как оптимизировать запросы в MongoDB? Какие инструменты используете?
**Что ожидают услышать:**
- explain() для анализа плана выполнения
- Использование индексов (ESR правило: Equality, Sort, Range)
- Покрывающие индексы (covered queries)
- Проекция полей для уменьшения трафика
- Инструменты: MongoDB Compass, explain, profiler
- Метрики: executionTimeMillis, totalDocsExamined, totalKeysExamined
```go
cursor, err := collection.Find(ctx, filter, options.Find().SetProjection(bson.M{
    "name": 1,
    "email": 1,
    "_id": 0,
}))
```

### 13. Как реализовать connection pooling в Go приложении с MongoDB?
**Что ожидают услышать:**
- Встроенный пул соединений в mongo-driver
- Настройка параметров:
```go
clientOpts := options.Client().
    ApplyURI(uri).
    SetMaxPoolSize(100).           // Максимум соединений
    SetMinPoolSize(10).            // Минимум соединений
    SetMaxConnIdleTime(30 * time.Second).  // Timeout для idle соединений
    SetServerSelectionTimeout(5 * time.Second)

client, err := mongo.Connect(ctx, clientOpts)
```
- Использование одного клиента для всего приложения
- Health checks с помощью Ping()
- Graceful shutdown

### 14. Расскажите про Change Streams. Как их использовать в Go?
**Что ожидают услышать:**
- Реактивное отслеживание изменений в реальном времени
- Работает на основе oplog (replica set)
- Фильтрация событий
- Пример:
```go
pipeline := mongo.Pipeline{
    {{"$match", bson.D{
        {"operationType", bson.D{
            {"$in", bson.A{"insert", "update"}},
        }},
    }}},
}

stream, err := collection.Watch(ctx, pipeline)
defer stream.Close(ctx)

for stream.Next(ctx) {
    var event bson.M
    if err := stream.Decode(&event); err != nil {
        log.Fatal(err)
    }
    // Обработка изменения
    fmt.Printf("Change: %v\n", event)
}
```
- Use cases: кеш инвалидация, синхронизация, notifications

### 15. Как обрабатывать ошибки при работе с MongoDB в Go?
**Что ожидают услышать:**
- Проверка конкретных ошибок:
```go
import "go.mongodb.org/mongo-driver/mongo"

if err != nil {
    if mongo.IsDuplicateKeyError(err) {
        // Обработка duplicate key
    } else if mongo.IsTimeout(err) {
        // Обработка timeout
    } else if mongo.IsNetworkError(err) {
        // Обработка сетевой ошибки
    }
}
```
- Retry логика для транзакций
- Context timeout для предотвращения зависаний
- Graceful degradation
- Circuit breaker pattern для защиты от перегрузки

### 16. Объясните модели данных в MongoDB: embedding vs referencing. Когда что использовать?
**Что ожидают услышать:**
- **Embedding** (вложенные документы):
  - Связанные данные в одном документе
  - Быстрое чтение (один запрос)
  - Дублирование данных
  - Подходит для one-to-few отношений
```go
type User struct {
    ID      primitive.ObjectID `bson:"_id"`
    Name    string             `bson:"name"`
    Address Address            `bson:"address"` // Встроенный документ
}
```

- **Referencing** (ссылки):
  - Нормализация данных
  - Требует $lookup или несколько запросов
  - Подходит для many-to-many, one-to-many
```go
type Order struct {
    ID     primitive.ObjectID `bson:"_id"`
    UserID primitive.ObjectID `bson:"user_id"` // Ссылка
}
```

### 17. Как проектировать схему данных для высоконагруженного приложения?
**Что ожидают услышать:**
- Денормализация для производительности чтения
- Ограничение размера документа (16MB)
- Избегание unbounded arrays
- Паттерны:
  - **Subset Pattern**: хранение только части данных
  - **Bucket Pattern**: группировка данных по временным интервалам
  - **Outlier Pattern**: обработка аномалий отдельно
  - **Extended Reference Pattern**: дублирование часто используемых полей
- Пример Bucket Pattern для time-series:
```go
type Measurement struct {
    SensorID   string    `bson:"sensor_id"`
    Timestamp  time.Time `bson:"timestamp"`
    Values     []Reading `bson:"values"` // 1 час данных
}
```

### 18. Как мониторить производительность MongoDB из Go приложения?
**Что ожидают услышать:**
- Database profiler для медленных запросов
- CommandMonitor для отслеживания всех команд:
```go
cmdMonitor := &event.CommandMonitor{
    Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
        log.Printf("Command: %s, DB: %s", evt.CommandName, evt.DatabaseName)
    },
    Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
        log.Printf("Duration: %v", evt.Duration)
    },
}

clientOpts := options.Client().SetMonitor(cmdMonitor)
```
- Метрики: currentOp(), serverStatus(), dbStats()
- Интеграция с Prometheus/Grafana
- Алерты на медленные запросы, connection pool exhaustion
- APM инструменты (DataDog, New Relic)

### 19. Какие security best practices вы знаете для MongoDB?
**Что ожидают услышать:**
- Включение аутентификации (SCRAM-SHA-256)
- Role-based access control (RBAC)
- Шифрование данных:
  - In-transit (TLS/SSL)
  - At-rest (encrypted storage engine)
- Network isolation (binding to localhost, firewall)
- Аудит логи
- Пример безопасного подключения:
```go
credential := options.Credential{
    AuthMechanism: "SCRAM-SHA-256",
    Username:      "user",
    Password:      "pass",
}

clientOpts := options.Client().
    ApplyURI("mongodb://localhost:27017").
    SetAuth(credential).
    SetTLSConfig(&tls.Config{
        MinVersion: tls.VersionTLS12,
    })
```

### 20. Опишите архитектуру высоконагруженной системы с MongoDB. Как бы вы спроектировали систему для e-commerce с 1M+ пользователей?
**Что ожидают услышать:**

**Архитектура:**
- Sharded cluster для горизонтального масштабирования
  - 3+ shards с репликацией (PSS конфигурация)
  - Config servers (3 ноды)
  - Mongos routers (2+ инстанса)
- Shard key: user_id с hashed sharding

**Коллекции:**
```go
// Users - sharded по user_id
type User struct {
    ID       primitive.ObjectID `bson:"_id"`
    Email    string             `bson:"email"`
    Profile  UserProfile        `bson:"profile"` // Embedded
}

// Orders - sharded по user_id (co-located с пользователем)
type Order struct {
    ID         primitive.ObjectID `bson:"_id"`
    UserID     primitive.ObjectID `bson:"user_id"`
    Items      []OrderItem        `bson:"items"` // Embedded
    TotalPrice float64            `bson:"total_price"`
}

// Products - отдельная коллекция с кешированием
type Product struct {
    ID          primitive.ObjectID `bson:"_id"`
    Name        string             `bson:"name"`
    Price       float64            `bson:"price"`
    Stock       int                `bson:"stock"`
    CategoryID  primitive.ObjectID `bson:"category_id"`
}
```

**Оптимизации:**
- Read preference: secondaryPreferred для чтения каталога
- Write concern: majority для критичных операций (заказы, платежи)
- Индексы: compound на user_id + created_at для истории заказов
- Кеширование популярных товаров (Redis)
- CDN для статики
- Change streams для инвалидации кеша

**Go сервис:**
```go
type MongoService struct {
    client     *mongo.Client
    usersCol   *mongo.Collection
    ordersCol  *mongo.Collection
    productsCol *mongo.Collection
}

func NewMongoService(uri string) (*MongoService, error) {
    opts := options.Client().
        ApplyURI(uri).
        SetMaxPoolSize(100).
        SetReadPreference(readpref.SecondaryPreferred()).
        SetWriteConcern(writeconcern.Majority())
    
    client, err := mongo.Connect(context.Background(), opts)
    // ...
}
```

**Мониторинг:**
- Метрики: latency, throughput, connection pool
- Alerts: replication lag, disk usage, slow queries
- Distributed tracing (OpenTelemetry)

**Backup/Recovery:**
- Continuous backups (Ops Manager/Atlas)
- Point-in-time recovery
- Disaster recovery план

**Масштабирование:**
- Вертикальное: увеличение RAM для working set
- Горизонтальное: добавление shards при росте данных
- Разделение чтения/записи (replica set)

---

## Дополнительные темы для обсуждения

- **Migrations**: инструменты для миграций схемы (migrate, atlas)
- **Testing**: integration тесты с testcontainers-go
- **Geospatial queries**: работа с геоданными
- **Full-text search**: MongoDB Atlas Search vs Elasticsearch
- **Time-series collections**: оптимизация для временных рядов (MongoDB 5.0+)
- **GridFS**: хранение больших файлов
- **Capped collections**: коллекции фиксированного размера для логов
- **Multi-tenancy**: стратегии изоляции данных разных клиентов
