package routes

import (
	"go-db-demo/web/handlers"

	"github.com/gin-gonic/gin"
)

func SetupHomeRoutes(router *gin.Engine, homeHandler *handlers.HomeHandler) {
	router.GET("/", homeHandler.Index)
}

func SetupOrganizationRoutes(router *gin.Engine, organizationHandler *handlers.OrganizationHandler) {
	orgGroup := router.Group("/organizations")
	{
		orgGroup.GET("", organizationHandler.List)
		orgGroup.GET("/:id", organizationHandler.Index)
		orgGroup.GET("/new", organizationHandler.New)
		orgGroup.POST("", organizationHandler.Create)
		orgGroup.GET("/:id/edit", organizationHandler.Edit)
		orgGroup.PUT("/:id", organizationHandler.Update)
		orgGroup.POST("/:id/delete", organizationHandler.Delete)
	}
}

func SetupJobRoutes(router *gin.Engine, jobHandler *handlers.JobHandler) {
	jobGroup := router.Group("/jobs")
	{
		jobGroup.GET("", jobHandler.List)
		jobGroup.GET("/:id", jobHandler.Index)
		jobGroup.GET("/new", jobHandler.New)
		jobGroup.POST("", jobHandler.Create)
		jobGroup.GET("/:id/edit", jobHandler.Edit)
		jobGroup.PUT("/:id", jobHandler.Update)
		jobGroup.POST("/:id/delete", jobHandler.Delete)
	}
}

func SetupUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("", userHandler.List)
		userGroup.GET("/:id", userHandler.Index)
		userGroup.GET("/new", userHandler.New)
		userGroup.POST("", userHandler.Create)
		userGroup.GET("/:id/edit", userHandler.Edit)
		userGroup.PUT("/:id", userHandler.Update)
		userGroup.POST("/:id/delete", userHandler.Delete)
	}
}
