
Задача:
есть таблица users
с полями id, name, order_count
1 млн строк
имена юзеров вида "Den", "Max", "John", "Bob", "Alice", "Ann", "Vasya", "Petya", "Masha", "Olya"
order_count у всех 0
необходимо всем изменить order_count, увеличив на 1


```
docker exec postgres-primary psql -U postgres -c "SELECT * FROM users ORDER BY id LIMIT 10;"
 id | name  | order_count 
----+-------+-------------
  1 | Den   |           1
  2 | Alice |           1
  3 | Max   |           1
  4 | Max   |           1
  5 | Olya  |           1
  6 | Olya  |           1
  7 | Max   |           1
  8 | Olya  |           1
  9 | Max   |           1
 10 | Bob   |           1
(10 rows)

```



 Когда использовать что?
✅ Используй Batch Update когда:

Массовое обновление по условию (SET x = y WHERE ...)
Вся логика - это SQL (вычисления, проверки)
Нет внешних зависимостей (API, файлы)
Примеры:

sql  -- Пометить старые заказы
  UPDATE orders SET status = 'archived' WHERE created_at < NOW() - INTERVAL '1 year'
  
  -- Пересчитать счетчики
  UPDATE users SET order_count = (SELECT COUNT(*) FROM orders WHERE user_id = users.id)
  
  -- Обновить цены с учетом инфляции
  UPDATE products SET price = price * 1.05 WHERE category = 'electronics'
✅ Используй Worker Pool когда:

Для каждой записи нужно:

🌐 Вызвать внешний API
📁 Прочитать/записать файл
🧮 Сложные вычисления в Go (не SQL)
📧 Отправить email/уведомление


Примеры:

go  // Загрузить аватары пользователей из внешнего API
  for userID := range jobs {
      avatar := fetchAvatarFromAPI(userID)  // HTTP запрос
      db.Exec("UPDATE users SET avatar = $1 WHERE id = $2", avatar, userID)
  }
  
  // Обработать изображения
  for productID := range jobs {
      image := loadImage(productID)
      thumbnail := resizeImage(image)  // CPU-intensive операция
      saveThumbnail(thumbnail)
  }
```

## 🎓 Золотое правило:

> **"Отдавай работу туда, где она выполняется эффективнее"**
```
PostgreSQL ❤️:                    Go ❤️:
- Работа с множествами данных     - Внешние API
- JOIN'ы, GROUP BY, агрегации     - Сложная бизнес-логика
- Обновления по условию           - Работа с файлами
- Транзакционность                - Параллельные вычисления
💡 Твой случай:
sqlUPDATE users SET order_count = 1 WHERE id BETWEEN 1 AND 100000

✅ Простое условие (BETWEEN)
✅ Простое значение (константа 1)
✅ Никаких внешних зависимостей
✅ Никаких вычислений

→ PostgreSQL справится за 0.2 секунды, Go + Worker Pool за 5 секунд