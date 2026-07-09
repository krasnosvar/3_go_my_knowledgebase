# 1. create db in posgres
docker compose up -d
cd 1_postgres/docker_compose
# via command
docker compose exec -T db-primary psql -U postgres -d postgres -c "CREATE DATABASE test_1;"
# via script
docker compose exec -T db-primary psql -U postgres -d postgres -v ON_ERROR_STOP=1 < ../exercises/1_task_db/simple_crud_app/scripts/schema.sql


# 2 create files, init module
go mod init crud-service

# 3 db.go - create database connection InitDB() function
# 4 create structs in models.go to describe tables in Postgres (to work with rows in golang code)
# 5 db_functions.go - create functions executing SQL-queries in postgres
# 6 handlers.go create handlers- functions, that are on API-call executing SQL-functions in db_functions.go
# 7 handlers_tx.go - create handlers with transactions ( both handlers and db-crud functions )

# CRUD
# ───────────────────────────────────────────────
# 7. CRUD API TESTS (users)
# Create
curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"order_count": 5}'

# Read
curl -X GET "http://localhost:8080/users?id=1"

# Update
curl -X PUT http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"ID":1,"OrderCount":10}'

# Delete
curl -X DELETE "http://localhost:8080/users?id=1"

# ───────────────────────────────────────────────
# 8. CRUD API TESTS (users_working_with_transactions)
# Create (Tx)
curl -X POST http://localhost:8080/users_working_with_transactions \
     -H "Content-Type: application/json" \
     -d '{"order_count": 5}'

# Update (Tx)
curl -X PUT http://localhost:8080/users_working_with_transactions \
     -H "Content-Type: application/json" \
     -d '{"ID":1,"OrderCount":15}'

# Delete (Tx)
curl -X DELETE "http://localhost:8080/users_working_with_transactions?id=1"




#######################################################################################

# GORM
# Create
curl -X POST http://localhost:8888/users \
     -H "Content-Type: application/json" \
     -d '{"order_count": 5}'

# Read
curl -X GET "http://localhost:8888/users?id=1"

# Update
curl -X PUT http://localhost:8888/users \
     -H "Content-Type: application/json" \
     -d '{"ID":1,"OrderCount":10}'

# Delete
curl -X DELETE "http://localhost:8888/users?id=1"
