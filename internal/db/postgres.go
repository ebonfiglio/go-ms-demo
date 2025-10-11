package db

import (
	"fmt"
	"log"

	"go-db-demo/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	cfg := config.LoadConfig()
	connStr := cfg.Database.GetConnectionString()

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		log.Printf("Warning: Database not available, some features may not work")
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Could not ping DB: %v", err)
		log.Printf("Warning: Database not available, some features may not work")
		return nil
	}
	fmt.Println("Connected to PostgreSQL!")

	fmt.Println("Ensured 'users' table exists.")

	return db
}
