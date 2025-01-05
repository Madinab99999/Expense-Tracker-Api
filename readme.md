# Expense Tracker Api

Cсылка на задание:
https://roadmap.sh/projects/expense-tracker-api

Файлы конфигурации для запуска проекта в докере, разработаны для Docker Toolbox v19.03.1

Перед запуском необходимо создать .env файл:

```shell
API_HOST=localhost
API_PORT=4012
DB_NAME=expense-tracker
DB_USER=postgres
DB_PASSWORD=postgresql
DB_PORT=5432
DB_HOST=localhost
TOKEN_SECRET=ExpenseTrackerSecret
TOKEN_PEPPER=ExpenseTrackerPepper
```

Database Seeds:

```shell
go run ./internal/cli/... seed
```
