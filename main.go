package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/microsoft/go-mssqldb"

	"query-to-api-service/config"
	"query-to-api-service/handlers"
	"query-to-api-service/middleware"
	"query-to-api-service/storage"
)

type DynamicRouter struct {
	mux *mux.Router
	mu  sync.RWMutex
}

func (dr *DynamicRouter) HandleFunc(path string, handler http.HandlerFunc) {
	dr.mu.Lock()
	defer dr.mu.Unlock()
	dr.mux.HandleFunc(path, handler).Methods("GET")
}

func connectDatabase(cfg *config.Config) (*sql.DB, error) {
	var dsn string
	var driver string

	switch cfg.DBType {
	case "postgres":
		driver = "postgres"
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	case "mssql":
		driver = "sqlserver"
		query := url.Values{}
		query.Add("database", cfg.DBName)
		query.Add("encrypt", "true")
		query.Add("trustServerCertificate", "false")

		dsnURL := &url.URL{
			Scheme:   "sqlserver",
			User:     url.UserPassword(cfg.DBUser, cfg.DBPassword),
			Host:     fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
			RawQuery: query.Encode(),
		}
		dsn = dsnURL.String()

	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.DBType)
	}

	return sql.Open(driver, dsn)
}

func main() {
	cfg := config.LoadConfig()

	// Connect to database based on type
	db, err := connectDatabase(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	log.Printf("✓ Database connected successfully (type: %s)", cfg.DBType)

	// Initialize repository
	repo := storage.NewRepository(db)

	// Setup router
	dynamicRouter := &DynamicRouter{mux: mux.NewRouter()}

	// Load and register persisted endpoints
	loadPersistedEndpoints(repo, dynamicRouter)

	// Admin handler
	adminHandler := handlers.NewAdminHandler(repo, dynamicRouter)

	// Main router
	mainRouter := mux.NewRouter()

	// Health check endpoint
	mainRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	}).Methods("GET")

	// Admin endpoints with auth middleware
	adminRouter := mux.NewRouter()
	adminRouter.HandleFunc("/admin/api/create", adminHandler.HandleCreateEndpoint).Methods("POST")
	mainRouter.PathPrefix("/admin").Handler(
		middleware.AdminAuthMiddleware(cfg.AdminKey)(adminRouter),
	)

	// Dynamic API endpoints
	mainRouter.PathPrefix("/api").Handler(dynamicRouter.mux)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Starting server on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mainRouter); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}

func loadPersistedEndpoints(repo *storage.Repository, router *DynamicRouter) {
	definitions, err := repo.GetAllActiveDefinitions()
	if err != nil {
		log.Printf("⚠️  Warning: Failed to load persisted endpoints: %v", err)
		return
	}

	db := repo.GetDB()
	for _, def := range definitions {
		router.HandleFunc(def.Path, handlers.HandlerFactory(db, def.Query))
		log.Printf("✓ Registered endpoint: %s", def.Path)
	}

	log.Printf("✓ Loaded %d persisted endpoints", len(definitions))
}
