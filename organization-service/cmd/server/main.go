package main

import (
	"context"
	"fmt"
	"go-ms-demo/organization-service/internal/config"
	"go-ms-demo/organization-service/internal/db"
	"go-ms-demo/organization-service/internal/handlers"
	"go-ms-demo/organization-service/internal/routers"
	"go-ms-demo/organization-service/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Organization Service...")

	cfg := config.LoadConfig()

	// Run database migrations
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	log.Println("Running database migrations...")
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to run")
		} else {
			log.Fatalf("Failed to run migrations: %v", err)
		}
	} else {
		log.Println("Migrations completed successfully")
	}

	// Close migration instance
	sourceErr, dbErr := m.Close()
	if sourceErr != nil {
		log.Printf("Error closing migration source: %v", sourceErr)
	}
	if dbErr != nil {
		log.Printf("Error closing migration database: %v", dbErr)
	}

	// Connect to database for application
	database := db.Connect()
	defer database.Close()

	router := gin.Default()

	orgRepo := db.NewOrganizationRepository(database)
	orgService := services.NewOrganizationService(orgRepo)
	orgHandler := handlers.NewHandler(orgService)

	routers.SetupRoutes(router, orgHandler)

	// Create HTTP server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
