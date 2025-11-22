package storage

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// CreateAPIDefinition saves a new API definition
func (r *Repository) CreateAPIDefinition(path, query string) (*APIDefinition, error) {
	var id int

	err := r.db.QueryRow(
		`INSERT INTO _api_definitions (path, query, created_at, updated_at, is_active)
         OUTPUT INSERTED.id
         VALUES (@p1, @p2, @p3, @p4, 1)`,
		path, query, time.Now(), time.Now(),
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &APIDefinition{
		ID:        id,
		Path:      path,
		Query:     query,
		CreatedAt: time.Now(),
		IsActive:  true,
	}, nil
}

// GetAllActiveDefinitions retrieves all active API definitions
func (r *Repository) GetAllActiveDefinitions() ([]APIDefinition, error) {
	rows, err := r.db.Query(
		`SELECT id, path, query, created_at, updated_at, is_active
         FROM _api_definitions
         WHERE is_active = 1
         ORDER BY created_at DESC`,
	)
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

	err := r.db.QueryRow(
		`SELECT id, path, query, created_at, updated_at, is_active
         FROM _api_definitions
         WHERE path = @p1 AND is_active = 1`,
		path,
	).Scan(&def.ID, &def.Path, &def.Query, &def.CreatedAt, &def.UpdatedAt, &def.IsActive)

	if err != nil {
		return nil, err
	}

	return &def, nil
}

// GetDB returns the underlying database connection
func (r *Repository) GetDB() *sql.DB {
	return r.db
}
