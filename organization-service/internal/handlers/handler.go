package handlers

import (
	"go-ms-demo/organization-service/internal/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	orgService domain.OrganizationService
}

func NewHandler(orgService domain.OrganizationService) *Handler {
	return &Handler{orgService: orgService}
}

// Request DTOs with validation
type CreateOrganizationRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
}

type UpdateOrganizationRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
}

// GetAllOrganizations handles GET /organizations
func (h *Handler) GetAllOrganizations(c *gin.Context) {
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve organizations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  organizations,
		"count": len(organizations),
	})
}

// GetOrganization handles GET /organizations/:id
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
		// Check if it's a "not found" error
		if strings.Contains(err.Error(), "no rows") || strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Organization not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve organization",
		})
		return
	}

	c.JSON(http.StatusOK, org)
}

// CreateOrganization handles POST /organizations
func (h *Handler) CreateOrganization(c *gin.Context) {
	var req CreateOrganizationRequest

	// Validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	// Create organization
	org := &domain.Organization{
		Name: strings.TrimSpace(req.Name),
	}

	createdOrg, err := h.orgService.CreateOrganization(org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create organization",
		})
		return
	}

	c.JSON(http.StatusCreated, createdOrg)
}

// UpdateOrganization handles PUT /organizations/:id
func (h *Handler) UpdateOrganization(c *gin.Context) {
	// Validate ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid organization ID",
		})
		return
	}

	// Validate request body
	var req UpdateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	// Check if organization exists
	_, err = h.orgService.GetOrganization(id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") || strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Organization not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check organization",
		})
		return
	}

	// Update organization
	org := &domain.Organization{
		ID:   id,
		Name: strings.TrimSpace(req.Name),
	}

	updatedOrg, err := h.orgService.UpdateOrganization(org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update organization",
		})
		return
	}

	c.JSON(http.StatusOK, updatedOrg)
}

// DeleteOrganization handles DELETE /organizations/:id
func (h *Handler) DeleteOrganization(c *gin.Context) {
	// Validate ID parameter
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid organization ID",
		})
		return
	}

	// Delete organization
	rowsAffected, err := h.orgService.DeleteOrganization(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete organization",
		})
		return
	}

	// Check if organization existed
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Organization not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Organization deleted successfully",
		"id":      id,
	})
}
