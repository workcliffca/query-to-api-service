package storage

import "time"

type APIDefinition struct {
	ID        int       `json:"id"`
	Path      string    `json:"path"`
	Query     string    `json:"query"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

type CreateAPIRequest struct {
	Path  string `json:"path" binding:"required"`
	Query string `json:"query" binding:"required"`
}

type APIResponse struct {
	Data  []map[string]interface{} `json:"data"`
	Count int                      `json:"count"`
}
