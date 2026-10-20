# Реализация ключевых узлов System Design кейсов на Go

## О чём эта тема (краткий контекст)

На завершающем этапе System Design интервью (Deep Dive) интервьюер просит
детализировать самый критичный узел системы: как именно сервис не допустит
двойного списания денег при сетевом ретрае, как не упадёт по памяти при загрузке
гигабайтного видео, как сгенерирует короткую ссылку и защитит базу от шторма
запросов при вирусном трафике. Ниже разобраны готовые идиоматичные решения этих
узлов на Go.

> **Базовая теория и инфраструктурный фундамент (DevOps / SRE):**
> Полный разбор концепции, базовые определения, схемы, алгоритмы и
> платформенные компромиссы вынесены в канонический раздел:
> - [Сквозной разбор 4 кейсов: Денежные переводы, Медиа-галерея, URL Shortener и News Feed](../../../../5_devops_sre_knowledgebase/14_system_design/07_case_studies_and_interview/01_theory/02_case_studies_transfers_and_media.md)
>
> **Фокус этого урока:** здесь мы не дублируем общую теорию, а разбираем
> исключительно особенности реализации и эксплуатации в экосистеме
> **Go** (планировщик G-M-P, сборщик мусора, netpoll, стандартная
> библиотека и идиомы языка).

## 1. Кейс «Денежные переводы»: `int64`, защита от Deadlock и Outbox в Go

При реализации биллинга и переводов средств на Go соблюдаются четыре правила:

1. **Никаких `float64` для денег!** Числа с плавающей точкой (`0.1 + 0.2 !=
   0.3`) теряют точность. В Go суммы хранятся либо в минимальных неделимых
   единицах (**копейки / центы в `int64`**), либо через библиотеку десятичной
   арифметики фиксированной точности (`shopspring/decimal`).
2. **Упорядоченная блокировка строк (`SELECT ... FOR UPDATE`) против
   Deadlock:** если горутина 1 переводит со счёта `A` на `B`, а горутина 2
   одновременно переводит с `B` на `A`, захват блокировок в порядке аргументов
   приведёт к взаимной блокировке (deadlock) в PostgreSQL. Всегда блокируйте
   счета строго **по возрастанию ID** (`firstID = min(from, to)`).
3. **Атомарная идемпотентность + Outbox в одном `*sql.Tx`:** проверка
   `Idempotency-Key`, списание/зачисление и запись события в таблицу `outbox`
   выполняются внутри одной транзакции.

```go
func ExecuteTransfer(ctx context.Context, db *sql.DB, idempKey string, fromID, toID, amountCents int64) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	// 1. Защита от дублей по Idempotency-Key.
	res, err := tx.ExecContext(ctx,
		`INSERT INTO idempotency_keys (key, created_at) VALUES ($1, NOW()) ON CONFLICT DO NOTHING`,
		idempKey,
	)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return nil // Перевод с этим ключом уже был успешно выполнен!
	}

	// 2. Блокируем оба счёта строго в порядке возрастания ID (защита от deadlock).
	firstID, secondID := fromID, toID
	if firstID > secondID {
		firstID, secondID = secondID, firstID
	}
	_, err = tx.ExecContext(ctx,
		`SELECT id FROM accounts WHERE id IN ($1, $2) ORDER BY id FOR UPDATE`,
		firstID, secondID,
	)
	if err != nil {
		return err
	}

	// 3. Проверяем баланс и обновляем счета.
	var balance int64
	if err := tx.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id = $1`, fromID).Scan(&balance); err != nil {
		return err
	}
	if balance < amountCents {
		return errors.New("insufficient funds")
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amountCents, fromID); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amountCents, toID); err != nil {
		return err
	}

	// 4. Записываем событие в Transactional Outbox в той же транзакции.
	if _, err := tx.ExecContext(ctx,
		`INSERT INTO outbox_events (aggregate_id, payload) VALUES ($1, $2)`,
		fromID, fmt.Sprintf(`{"key":%q,"from":%d,"to":%d,"cents":%d}`, idempKey, fromID, toID, amountCents),
	); err != nil {
		return err
	}

	return tx.Commit()
}
```

## 2. Кейс «Галерея изображений и медиа-пайплайн» в Go

Главная ловушка в Go при работе с файлами — вызов `io.ReadAll(r.Body)` при
загрузке файла пользователем. Если 50 пользователей одновременно отправят видео
по 100 МБ, `io.ReadAll` аллоцирует в куче `> 5 ГБ` слайсов (с учётом удвоения
при `append`), и под мгновенно упадёт по `OOMKilled`.

### Как правильно строить медиа-пайплайн на Go

1. **Прямая загрузка в S3 по Presigned URL:** Go API-сервис лишь проверяет права
   пользователя и за 1 мс генерирует подписанный URL через AWS SDK v2 /
   `minio-go`. Сам бинарный поток идёт от клиента напрямую в S3/MinIO, минуя
   память Go-подов.
2. **Если проксирование через Go неизбежно — только потоковый `io.Copy` с
   лимитом:** оборачивайте `r.Body` в `http.MaxBytesReader(w, r.Body,
   10<<20)` (жёсткий предел 10 МБ) и передавайте `io.Reader` потоком прямо в
   S3-клиент с фиксированным буфером 32 КБ (`io.CopyBuffer`).
3. **Вынос CPU-bound ресайза картинок из HTTP-хендлера:** декодирование и
   сжатие JPEG/WebP активно нагружает CPU и аллоцирует большие буферы.
   HTTP-хендлер только кладёт задачу в Kafka, а отдельный пул воркеров на Go с
   ограничением параллелизма по числу ядер
   (`errgroup.SetLimit(runtime.NumCPU())`)
   генерирует превью и складывает их в S3/CDN.

## 3. Кейс «URL Shortener» на Go: Base62 + L1/Redis + `singleflight`

Для сокращателя ссылок (см. рабочий пример
`02_examples/01_base62_shortener/main.go`
и упражнение `03_exercises/01_base62/main.go`) архитектура на Go складывается
из трёх быстрых узлов:

- **Генерация уникального ID без координации на каждый запрос:** каждый
  Go-инстанс
  либо генерирует 64-битный **Snowflake ID** (`timestamp + worker_id + sequence`
  под `sync.Mutex` / `atomic`), либо берёт у координатора (БД/etcd) диапазон из
  `10 000` ID сразу в память (`atomic.Uint64.Add(1)`), после чего кодирует число
  в строку **Base62** (`[0-9a-zA-Z]`, 7 символов = $62^7 \approx 3.5$ трлн
  ссылок).
- **Чтение редиректа (`GET /{code}` → `302 Found` или `301`):**
  1. Сначала проверяется локальный шардированный in-memory L1-кэш в памяти
     Go-процесса (~50 нс) с коротким TTL (например, 10–30 секунд для самых
     горячих вирусных ссылок).
  2. При промахе L1 вызов оборачивается в **`singleflight.Group`**, чтобы при
     вирусном спайке (100 000 кликов в секунду по одной ссылке) в Redis (L2) и
     PostgreSQL ушёл ровно **один** запрос с каждого Go-пода.
- **Асинхронный подсчёт кликов:** чтобы `UPDATE clicks = clicks + 1` не тормозил
  редирект, хендлер неблокирующе кидает событие в буферизованный канал
  `clicksChan`, а фоновая горутина раз в секунду сбрасывает агрегированные
  счётчики батчем в БД или Kafka.

## 4. Кейс «Лента новостей (News Feed)» на Go

В гибридной модели ленты (Push / Fan-out on Write для обычных пользователей и
Pull / Fan-out on Read для знаменитостей с миллионами подписчиков):

- **Fan-out воркеры на Go:** консьюмер читает событие нового поста из Kafka и
  раскладывает `post_id` в Redis Sorted Set (`ZADD feed:<user_id> <timestamp>
  <post_id>`) подписчиков пачками через **Redis Pipeline** в пуле горутин
  (`errgroup.WithContext` с лимитом `SetLimit(50)`).
- **Сборка ленты при чтении (`GET /feed`):** Go-сервис за один round-trip читает
  подготовленный кэш пользователя из Redis и параллельно (через `errgroup`)
  дотягивает свежие посты 3–5 «знаменитостей», на которых подписан пользователь,
  сливая отсортированные списки в памяти за микросекунды.
