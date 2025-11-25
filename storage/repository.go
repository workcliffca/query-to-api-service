package storage

import (
	"database/sql"
	"time"
)

type Repository struct {
	db     *sql.DB
	dbType string
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:     db,
		dbType: detectDBType(db),
	}
}

func detectDBType(db *sql.DB) string {
	// Try to detect database type by driver name
	// This is a simple heuristic - you could make it more sophisticated
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err == nil {
		return "postgres"
	}
	return "mssql"
}

// CreateAPIDefinition saves a new API definition
func (r *Repository) CreateAPIDefinition(path, query string) (*APIDefinition, error) {
	var id int
	now := time.Now()

	if r.dbType == "postgres" {
		err := r.db.QueryRow(
			`INSERT INTO _api_definitions (path, query, created_at, updated_at, is_active)
             VALUES ($1, $2, $3, $4, true)
             RETURNING id`,
			path, query, now, now,
		).Scan(&id)

		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.QueryRow(
			`INSERT INTO _api_definitions (path, query, created_at, updated_at, is_active)
             OUTPUT INSERTED.id
             VALUES (@p1, @p2, @p3, @p4, 1)`,
			path, query, now, now,
		).Scan(&id)

		if err != nil {
			return nil, err
		}
	}

	return &APIDefinition{
		ID:        id,
		Path:      path,
		Query:     query,
		CreatedAt: now,
		IsActive:  true,
	}, nil
}

// GetAllActiveDefinitions retrieves all active API definitions
func (r *Repository) GetAllActiveDefinitions() ([]APIDefinition, error) {
	var query string
	if r.dbType == "postgres" {
		query = `SELECT id, path, query, created_at, updated_at, is_active
                 FROM _api_definitions
                 WHERE is_active = true
                 ORDER BY created_at DESC`
	} else {
		query = `SELECT id, path, query, created_at, updated_at, is_active
                 FROM _api_definitions
                 WHERE is_active = 1
                 ORDER BY created_at DESC`
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var definitions []APIDefinition

	for rows.Next() {
		var def APIDefinition
		err := rows.Scan(&def.ID, &def.Path, &def.Query, &def.CreatedAt, &def.UpdatedAt, &def.IsActive)
		if err != nil {
			return nil, err
		}
		definitions = append(definitions, def)
	}

	return definitions, rows.Err()
}

// GetByPath retrieves a specific API definition by path
func (r *Repository) GetByPath(path string) (*APIDefinition, error) {
	var def APIDefinition
	var query string

	if r.dbType == "postgres" {
		query = `SELECT id, path, query, created_at, updated_at, is_active
                 FROM _api_definitions
                 WHERE path = $1 AND is_active = true`
	} else {
		query = `SELECT id, path, query, created_at, updated_at, is_active
                 FROM _api_definitions
                 WHERE path = @p1 AND is_active = 1`
	}

	err := r.db.QueryRow(query, path).Scan(
		&def.ID, &def.Path, &def.Query, &def.CreatedAt, &def.UpdatedAt, &def.IsActive,
	)

	if err != nil {
		return nil, err
	}

	return &def, nil
}

// GetDB returns the underlying database connection
func (r *Repository) GetDB() *sql.DB {
	return r.db
}
