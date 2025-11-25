package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"query-to-api-service/storage"
	"query-to-api-service/utils"
)

// HandlerFactory creates a dynamic GET handler using closure
func HandlerFactory(db *sql.DB, query string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set response header
		w.Header().Set("Content-Type", "application/json")

		// Execute the query
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, `{"error":"Query execution failed"}`, http.StatusInternalServerError)
			return
		}

		// Dynamically scan rows
		data, err := utils.ScanRowsToMaps(rows)
		if err != nil {
			http.Error(w, `{"error":"Row scanning failed"}`, http.StatusInternalServerError)
			return
		}

		// Build response
		response := storage.APIResponse{
			Data:  data,
			Count: len(data),
		}

		// Return JSON
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
