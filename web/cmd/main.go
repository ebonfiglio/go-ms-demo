package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"go-db-demo/internal/config"
	"go-db-demo/internal/db"
	"go-db-demo/internal/service"
	"go-db-demo/web"
	"go-db-demo/web/handlers"
	"go-db-demo/web/routes"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := db.Connect()
	if dbConn != nil {
		defer dbConn.Close()
	}

	var orgService *service.OrganizationService
	var jobService *service.JobService
	var userService *service.UserService

	if dbConn != nil {
		orgRepo := db.NewOrganizationRepository(dbConn)
		orgService = service.NewOrganizationService(orgRepo)

		jobRepo := db.NewJobRepository(dbConn)
		jobService = service.NewJobService(jobRepo)

		userRepo := db.NewUserRepository(dbConn)
		userService = service.NewUserService(userRepo)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	htmlTemplate := web.Parse()
	if htmlTemplate != nil {
		router.SetHTMLTemplate(htmlTemplate)
		log.Println("Templates loaded successfully")
	} else {
		log.Println("Warning: No templates loaded")
	}

	homeHandler := handlers.NewHomeHandler()
	routes.SetupHomeRoutes(router, homeHandler)

	organizationHandler := handlers.NewOrganizationHandler(orgService)
	routes.SetupOrganizationRoutes(router, organizationHandler)

	jobHander := handlers.NewJobHandler(jobService, orgService)
	routes.SetupJobRoutes(router, jobHander)

	userHander := handlers.NewUserHandler(userService, jobService, orgService)
	routes.SetupUserRoutes(router, userHander)

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", serverAddr)
	_ = router.Run(serverAddr)
}
