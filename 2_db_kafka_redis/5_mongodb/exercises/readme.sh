# 3. Проверка через curl (HTTP интерфейс MongoDB отсутствует)
# MongoDB не имеет встроенного HTTP‑API, поэтому проверка идёт через Mongo Shell или mongosh:

# Подключение к контейнеру
docker exec -it mongodb mongosh -u admin -p admin --authenticationDatabase admin

# Вставка документа
use test_db
db.users.insertOne({name: "Denis", order_count: 5})

# Чтение
db.users.find({name: "Denis"})

# Обновление
db.users.updateOne({name: "Denis"}, {$set: {order_count: 10}})

# Удаление
db.users.deleteOne({name: "Denis"})
