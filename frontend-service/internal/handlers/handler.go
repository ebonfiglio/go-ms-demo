package handlers

import (
	"go-ms-demo/frontend-service/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	config *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{config: cfg}
}

func (h *Handler) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Microservices Dashboard",
	})
}

func (h *Handler) OrganizationsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "organizations.html", gin.H{
		"Title":                  "Organizations",
		"OrganizationServiceURL": h.config.Services.OrganizationServiceURL,
	})
}
