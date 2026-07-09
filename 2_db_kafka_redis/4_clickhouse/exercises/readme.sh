# Выполнить простой запрос напрямую
curl "http://localhost:8123/?query=SELECT+1"

# Вставка данных через HTTP
curl -X POST "http://localhost:8123/?query=INSERT+INTO+users+(id,name,age)+VALUES+(2,'Alex',25)"

# Чтение данных через HTTP
curl "http://localhost:8123/?query=SELECT+*+FROM+users"
