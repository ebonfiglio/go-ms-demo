package routers

import (
	"go-ms-demo/organization-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, organizationHandler *handlers.Handler) {
	orgGroup := router.Group("/organizations")
	{
		orgGroup.GET("", organizationHandler.GetAllOrganizations)
		orgGroup.GET("/:id", organizationHandler.GetOrganization)
		orgGroup.POST("", organizationHandler.CreateOrganization)
		orgGroup.PUT("/:id", organizationHandler.UpdateOrganization)
		orgGroup.DELETE("/:id", organizationHandler.DeleteOrganization)
	}
}
