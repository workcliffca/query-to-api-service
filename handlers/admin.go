package handlers

import (
	"encoding/json"
	"net/http"

	"query-to-api-service/storage"
	"query-to-api-service/utils"
)

type AdminHandler struct {
	repo   *storage.Repository
	router RouterManager
}

// RouterManager interface for registering routes dynamically
type RouterManager interface {
	HandleFunc(path string, handler http.HandlerFunc)
}

func NewAdminHandler(repo *storage.Repository, router RouterManager) *AdminHandler {
	return &AdminHandler{
		repo:   repo,
		router: router,
	}
}

// HandleCreateEndpoint processes admin requests to create new API endpoints
func (ah *AdminHandler) HandleCreateEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req storage.CreateAPIRequest

	// Parse JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
		return
	}

	// Validate path
	if err := utils.ValidateAPIPath(req.Path); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Validate query
	if err := utils.ValidateReadOnlyQuery(req.Query); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Persist to database
	apiDef, err := ah.repo.CreateAPIDefinition(req.Path, req.Query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create API definition"})
		return
	}

	// Register the dynamic handler
	db := ah.repo.GetDB()
	ah.router.HandleFunc(req.Path, HandlerFactory(db, req.Query))

	// Return response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiDef)
}
