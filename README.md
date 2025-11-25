# Query-to-API Service

A dynamic API service that allows you to create REST endpoints backed by database queries on the fly.

## Features

- Create dynamic API endpoints via admin interface
- Support for both PostgreSQL and MS SQL Server
- Docker-based local development with PostgreSQL
- Production-ready with Azure SQL support

## Database Support

This service supports two database types:

- **PostgreSQL** - Recommended for local testing
- **MS SQL Server** - For Azure/production deployments

## Local Development Setup

### Prerequisites

- Go 1.24.2 or higher
- Docker and Docker Compose

### Quick Start with PostgreSQL

1. Start the PostgreSQL database:
```bash
docker-compose up -d
```

2. Verify the database is running:
```bash
docker-compose ps
```

3. Run the application:
```bash
go run main.go
```

The service will start on `http://localhost:8080` and automatically connect to the PostgreSQL database.

### Configuration

Copy the example environment file:
```bash
cp .env.example .env
```

Edit `.env` to configure your database connection:

**For PostgreSQL (local testing):**
```env
DB_TYPE=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=queryapi_user
DB_PASSWORD=queryapi_password
DB_NAME=queryapi_db
```

**For MS SQL Server (production):**
```env
DB_TYPE=mssql
DB_HOST=your-server.database.windows.net
DB_PORT=1433
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=queryapi_db
```

### Database Management

**Connect to PostgreSQL database:**
```bash
docker-compose exec postgres psql -U queryapi_user -d queryapi_db
```

**View logs:**
```bash
docker-compose logs -f postgres
```

**Stop the database:**
```bash
docker-compose down
```

**Reset the database (removes all data):**
```bash
docker-compose down -v
docker-compose up -d
```

## API Usage

### Health Check
```bash
curl http://localhost:8080/health
```

### Create a Dynamic Endpoint
```bash
curl -X POST http://localhost:8080/admin/api/create \
  -H "X-Admin-Key: your-secret-admin-key" \
  -H "Content-Type: application/json" \
  -d '{
    "path": "/api/users",
    "query": "SELECT * FROM users"
  }'
```

### Access Your Dynamic Endpoint
```bash
curl http://localhost:8080/api/users
```

## Project Structure

```
.
├── config/          # Configuration management
├── docker/          # Docker files
├── handlers/        # HTTP request handlers
├── k8s/            # Kubernetes manifests
├── middleware/     # HTTP middleware
├── migrations/     # Database schemas
│   ├── schema.sql          # MS SQL schema
│   └── postgres_schema.sql # PostgreSQL schema
├── storage/        # Database repositories
├── utils/          # Utility functions
├── docker-compose.yml  # Local PostgreSQL setup
├── main.go         # Application entry point
└── README.md       # This file
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_TYPE | Database type: "postgres" or "mssql" | postgres |
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 (postgres), 1433 (mssql) |
| DB_USER | Database username | queryapi_user |
| DB_PASSWORD | Database password | queryapi_password |
| DB_NAME | Database name | queryapi_db |
| SERVER_PORT | HTTP server port | 8080 |
| ADMIN_API_KEY | Admin API authentication key | default-secret-key |

## Security Notes

- Always change the `ADMIN_API_KEY` in production
- Use strong passwords for database connections
- Enable SSL/TLS for production database connections
- The default PostgreSQL setup disables SSL for local development convenience

## Switching Between Databases

Simply change the `DB_TYPE` environment variable and restart the application:

```bash
# Use PostgreSQL
export DB_TYPE=postgres

# Use MS SQL Server
export DB_TYPE=mssql
```

The application will automatically use the correct SQL syntax for your chosen database.