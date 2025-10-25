package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Services ServicesConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type ServicesConfig struct {
	OrganizationServiceURL string
	UserServiceURL         string
	JobServiceURL          string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables and defaults")
	}

	return &Config{
		Server: ServerConfig{
			Host: getEnvString("SERVER_HOST", "localhost"),
			Port: getEnvString("SERVER_PORT", "3000"),
		},
		Services: ServicesConfig{
			OrganizationServiceURL: getEnvString("ORGANIZATION_SERVICE_URL", "http://localhost:8080"),
			UserServiceURL:         getEnvString("USER_SERVICE_URL", "http://localhost:8081"),
			JobServiceURL:          getEnvString("JOB_SERVICE_URL", "http://localhost:8082"),
		},
	}
}

func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
