package main

import (
	"context"
	"go-ms-demo/organization-service/internal/db"
	"go-ms-demo/organization-service/internal/handlers"
	"go-ms-demo/organization-service/internal/routers"
	"go-ms-demo/organization-service/internal/services"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Organization Service...")

	database := db.Connect()
	defer database.Close()

	router := gin.Default()

	orgRepo := db.NewOrganizationRepository(database)
	orgService := services.NewOrganizationService(orgRepo)
	orgHandler := handlers.NewHandler(orgService)

	routers.SetupRoutes(router, orgHandler)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Close database connection cleanly
	if err := database.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Server exited")

}
