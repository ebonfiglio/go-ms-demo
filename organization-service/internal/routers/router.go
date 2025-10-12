package router

import (
	"go-ms-demo/organization-service/internal/domain"
	"go-ms-demo/organization-service/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, organizationHandler *handlers.Handler) {
	orgGroup := router.Group("/organizations")
	{
		orgGroup.GET("/test", func(c *gin.Context) {
			testOrg := domain.Organization{
				ID:   1,
				Name: "Test",
			}
			c.JSON(http.StatusOK, gin.H{
				"message":      "Test endpoint working",
				"organization": testOrg,
			})
		})
	}
}
