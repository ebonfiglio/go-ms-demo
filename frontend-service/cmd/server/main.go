package main

import (
	"go-ms-demo/frontend-service/internal/config"
	"go-ms-demo/frontend-service/internal/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Frontend Service...")

	cfg := config.LoadConfig()

	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Create handlers
	handler := handlers.NewHandler(cfg)

	// Routes
	router.GET("/", handler.HomePage)
	router.GET("/organizations", handler.OrganizationsPage)

	// Start server
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Frontend service starting on %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
