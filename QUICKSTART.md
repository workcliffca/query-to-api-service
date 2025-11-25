# Quick Start Guide - PostgreSQL Local Testing

This guide will help you get started with local testing using PostgreSQL in just a few minutes.

## Prerequisites

- Docker and Docker Compose installed
- Go 1.24.2 or higher

## Step 1: Start PostgreSQL

```bash
docker-compose up -d
```

This will:
- Start a PostgreSQL 16 container
- Create the `queryapi_db` database
- Run the migration script to create the `_api_definitions` table
- Expose PostgreSQL on port 5432

## Step 2: Verify Database is Running

```bash
docker-compose ps
```

You should see the `queryapi-postgres` container running.

## Step 3: Run the Application

The application is pre-configured to use PostgreSQL by default:

```bash
go run main.go
```

You should see output like:
```
✓ Database connected successfully (type: postgres)
✓ Loaded 0 persisted endpoints
🚀 Starting server on http://localhost:8080
```

## Step 4: Test the Application

### Health Check
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{"status":"healthy"}
```

### Create a Test Table

Connect to PostgreSQL:
```bash
docker-compose exec postgres psql -U queryapi_user -d queryapi_db
```

Create a test table:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

INSERT INTO users (name, email) VALUES
    ('Alice', 'alice@example.com'),
    ('Bob', 'bob@example.com'),
    ('Charlie', 'charlie@example.com');

\q
```

### Create a Dynamic API Endpoint

```bash
curl -X POST http://localhost:8080/admin/api/create \
  -H "X-Admin-Key: default-secret-key" \
  -H "Content-Type: application/json" \
  -d '{
    "path": "/api/users",
    "query": "SELECT * FROM users"
  }'
```

### Query Your New Endpoint

```bash
curl http://localhost:8080/api/users
```

Expected response:
```json
{
  "data": [
    {"id": 1, "name": "Alice", "email": "alice@example.com"},
    {"id": 2, "name": "Bob", "email": "bob@example.com"},
    {"id": 3, "name": "Charlie", "email": "charlie@example.com"}
  ],
  "count": 3
}
```

## Switching to Production (MS SQL Server)

When you're ready to deploy to production with Azure SQL:

1. Create a `.env` file:
```env
DB_TYPE=mssql
DB_HOST=your-server.database.windows.net
DB_PORT=1433
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=queryapi_db
ADMIN_API_KEY=your-production-secret-key
```

2. Source the environment:
```bash
source .env
go run main.go
```

## Troubleshooting

### PostgreSQL won't start
```bash
docker-compose down -v
docker-compose up -d
```

### Can't connect to database
Check if PostgreSQL is running:
```bash
docker-compose logs postgres
```

### Reset everything
```bash
docker-compose down -v
docker-compose up -d
go run main.go
```

## Useful Commands

| Command | Description |
|---------|-------------|
| `docker-compose up -d` | Start PostgreSQL |
| `docker-compose down` | Stop PostgreSQL |
| `docker-compose down -v` | Stop and delete all data |
| `docker-compose logs -f postgres` | View PostgreSQL logs |
| `docker-compose exec postgres psql -U queryapi_user -d queryapi_db` | Connect to database |

## Next Steps

- See [README.md](README.md) for full documentation
- Modify `.env` file for custom configuration
- Add your own tables and queries
- Deploy to production with MS SQL Server