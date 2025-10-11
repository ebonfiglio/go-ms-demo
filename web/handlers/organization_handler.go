package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"go-db-demo/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	orgService domain.OrganizationService
}

func NewOrganizationHandler(orgService domain.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{orgService: orgService}
}

func (h *OrganizationHandler) parseOrganizationID(c *gin.Context) (int64, bool) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		renderError(c, "organizations/list.html", "Invalid Organization ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (h *OrganizationHandler) validationOrganizationForm(name string) (string, string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", "Name is required"
	}
	return name, ""
}

func (h *OrganizationHandler) Index(c *gin.Context) {
	id, ok := h.parseOrganizationID(c)
	if !ok {
		return
	}

	org, err := h.orgService.GetOrganization(id)
	if err != nil {
		renderError(c, "organizations/list.html", "Organization not found", http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "organizations/index.html", gin.H{
		"Title":        "Organization - " + org.Name,
		"Organization": org,
	})
}

func (h *OrganizationHandler) List(c *gin.Context) {
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		renderError(c, "organizations/list.html", "Error loading organizations", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "organizations/list.html", gin.H{
		"Title":         "Organizations",
		"Organizations": organizations,
	})
}

func (h *OrganizationHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "organizations/form.html", gin.H{
		"Title":  "New Organization",
		"IsEdit": false,
	})
}

func (h *OrganizationHandler) Create(c *gin.Context) {
	name, errMsg := h.validationOrganizationForm(c.PostForm("name"))
	if errMsg != "" {
		h.renderFormWithError(c, errMsg, nil, false)
		return
	}

	organization := &domain.Organization{Name: name}
	_, err := h.orgService.CreateOrganization(organization)
	if err != nil {
		h.renderFormWithError(c, "Failed to create job", nil, false)
		return
	}

	c.Redirect(http.StatusSeeOther, "/organizations")
}

func (h *OrganizationHandler) Edit(c *gin.Context) {
	id, ok := h.parseOrganizationID(c)
	if !ok {
		return
	}

	organization, err := h.orgService.GetOrganization(id)
	if err != nil {
		renderError(c, "organizations/list/html", "Organization not found", http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "organizations/form.html", gin.H{
		"Title":        "Edit Organization",
		"Organization": organization,
		"IsEdit":       true,
	})
}

func (h *OrganizationHandler) Update(c *gin.Context) {
	id, ok := h.parseOrganizationID(c)
	if !ok {
		return
	}

	currentOrganization, err := h.orgService.GetOrganization(id)
	if err != nil {
		renderError(c, "organizations/list.html", "Organization not found", http.StatusNotFound)
		return
	}

	name, errMsg := h.validationOrganizationForm(c.PostForm("name"))
	if errMsg != "" {
		h.renderFormWithError(c, errMsg, currentOrganization, true)
		return
	}

	org := &domain.Organization{ID: id, Name: name}
	_, err = h.orgService.UpdateOrganization(org)
	if err != nil {
		h.renderFormWithError(c, "Failed to update job", currentOrganization, true)
		return
	}

	c.Redirect(http.StatusSeeOther, "/organizations")
}

func (h *OrganizationHandler) Delete(c *gin.Context) {
	id, ok := h.parseOrganizationID(c)
	if !ok {
		return
	}

	rowsAffected, err := h.orgService.DeleteOrganization(id)

	if err != nil {
		h.renderDeleteError(c, "Failed to delete organization: "+err.Error())
		return
	}

	if rowsAffected < 1 {
		h.renderDeleteError(c, "Organization not found or already deleted")
		return
	}

	c.Redirect(http.StatusSeeOther, "/organizations")
}

func (h *OrganizationHandler) renderDeleteError(c *gin.Context, message string) {
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		organizations = []domain.Organization{}
	}

	renderError(c, "organizations/list.html", message, http.StatusInternalServerError, gin.H{
		"Organizations": organizations,
	})
}

func (h *OrganizationHandler) renderFormWithError(c *gin.Context, errorMessage string, organization *domain.Organization, isEdit bool) {
	title := "New Organization"
	if isEdit {
		title = "Edit Organization"
	}

	data := gin.H{
		"Title":  title,
		"Error":  errorMessage,
		"IsEdit": isEdit,
	}

	if organization != nil {
		data["Organization"] = organization
	} else {
		data["FormData"] = gin.H{
			"Name": c.PostForm("name"),
		}
	}

	c.HTML(http.StatusBadRequest, "organizations/form.html", data)
}
