package main

import (
	"context"
	"go-ms-demo/organization-service/internal/config"
	"go-ms-demo/organization-service/internal/db"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg = config.LoadConfig()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Organization Service...")

	database := db.Connect()
	defer database.Close()

	router := gin.Default()

	orgRepo := db.NewOrganizationRepository(database)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Close database connection cleanly
	if err := database.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Server exited")

}
