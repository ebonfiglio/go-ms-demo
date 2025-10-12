package handlers

import (
	"go-ms-demo/organization-service/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	orgService domain.OrganizationService
}

func NewHandler(orgService domain.OrganizationService) *Handler {
	return &Handler{orgService: orgService}
}

func (h *Handler) GetOrganization(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid organization ID",
		})
		return
	}
	org, err := h.orgService.GetOrganization(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Organization not found",
		})
		return
	}

	c.JSON(http.StatusOK, org)
}
