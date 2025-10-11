package handlers

import (
	"go-db-demo/internal/domain"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService domain.UserService
	jobService  domain.JobService
	orgService  domain.OrganizationService
	validator   validator.Validate
}

func NewUserHandler(userService domain.UserService, jobService domain.JobService, orgService domain.OrganizationService) *UserHandler {
	return &UserHandler{
		userService: userService,
		jobService:  jobService,
		orgService:  orgService,
		validator:   *validator.New(),
	}
}

func (h *UserHandler) formatValidationErrors(err error) string {
	validationErrors := err.(validator.ValidationErrors)

	for _, e := range validationErrors {
		switch e.Tag() {
		case "required":
			return e.Field() + " is required"
		case "min":
			return e.Field() + " must be at least " + e.Param() + " characters"
		case "max":
			return e.Field() + " must be at most " + e.Param() + " characters"
		case "numeric":
			return e.Field() + " must be a valid number"
		}
	}

	return "Invalid input"
}

func (h *UserHandler) validateUserForm(name, orgIDStr, jobIDStr string) (string, int64, int64, string) {
	req := struct {
		Name  string `validate:"required,min=2,max=100"`
		OrgID string `validate:"required,numeric"`
		JobID string `validate:"required,numeric"`
	}{
		Name:  strings.TrimSpace(name),
		OrgID: orgIDStr,
		JobID: jobIDStr,
	}

	if err := h.validator.Struct(&req); err != nil {
		return "", 0, 0, h.formatValidationErrors(err)
	}

	orgID, _ := strconv.ParseInt(orgIDStr, 10, 64)
	jobID, _ := strconv.ParseInt(jobIDStr, 10, 64)

	return req.Name, orgID, jobID, ""
}

func (h *UserHandler) parseUserID(c *gin.Context) (int64, bool) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		renderError(c, "users/list.html", "Invalid User ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		renderError(c, "users/list.html", "Failed to load users", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "users/list.html", gin.H{
		"Title": "Users",
		"Users": users,
	})
}

func (h *UserHandler) New(c *gin.Context) {
	// TODO: User concurrency here
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		renderError(c, "users/list.html", "Failed to load organizations", http.StatusInternalServerError)
		return
	}

	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		renderError(c, "users/list.html", "Failed to load jobs", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "users/form.html", gin.H{
		"Title":         "New User",
		"Organizations": organizations,
		"Jobs":          jobs,
	})
}

func (h *UserHandler) Create(c *gin.Context) {
	name, orgID, jobID, errMsg := h.validateUserForm(c.PostForm("name"), c.PostForm("organizationID"), c.PostForm("jobID"))
	if errMsg != "" {
		h.renderFormWithError(c, errMsg, nil, false)
		return
	}

	user := &domain.User{
		Name:           name,
		OrganizationID: orgID,
		JobID:          jobID,
	}

	_, err := h.userService.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		h.renderFormWithError(c, "Failed to create user", user, false)
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func (h *UserHandler) Edit(c *gin.Context) {
	id, ok := h.parseUserID(c)
	if !ok {
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		renderError(c, "users/list/html", "User not found", http.StatusNotFound)
		return
	}
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		organizations = []domain.Organization{}
	}

	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		jobs = []domain.Job{}
	}

	c.HTML(http.StatusOK, "users/form.html", gin.H{
		"Title":         "Edit User",
		"User":          user,
		"Organizations": organizations,
		"Jobs":          jobs,
		"IsEdit":        true,
	})

}

func (h *UserHandler) Update(c *gin.Context) {
	id, ok := h.parseUserID(c)
	if !ok {
		return
	}

	currentUser, err := h.userService.GetUser(id)
	if err != nil {
		renderError(c, "users/list.html", "User not found", http.StatusNotFound)
		return
	}

	name, orgID, jobID, errMsg := h.validateUserForm(c.PostForm("name"), c.PostForm("organizationID"), c.PostForm("jobID"))
	if errMsg != "" {
		h.renderFormWithError(c, errMsg, currentUser, true)
		return
	}

	user := &domain.User{
		ID:             id,
		Name:           name,
		OrganizationID: orgID,
		JobID:          jobID,
	}
	_, err = h.userService.UpdateUser(user)
	if err != nil {
		h.renderFormWithError(c, "Failed to update user", currentUser, true)
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")

}

func (h *UserHandler) Delete(c *gin.Context) {
	id, ok := h.parseUserID(c)
	if !ok {
		return
	}

	rowsAffected, err := h.userService.DeleteUser(id)

	if err != nil {
		h.renderDeleteError(c, "Failed to delete user: "+err.Error())
		return
	}

	if rowsAffected < 1 {
		h.renderDeleteError(c, "User not found or already deleted")
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func (h *UserHandler) renderDeleteError(c *gin.Context, message string) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		users = []domain.User{}
	}

	renderError(c, "users/list.html", message, http.StatusInternalServerError, gin.H{
		"Users": users,
	})
}
func (h *UserHandler) renderFormWithError(c *gin.Context, errorMessage string, user *domain.User, isEdit bool) {
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		organizations = []domain.Organization{}
	}

	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		jobs = []domain.Job{}
	}

	title := "New User"
	if isEdit {
		title = "Edit User"
	}

	data := gin.H{
		"Title":         title,
		"Error":         errorMessage,
		"Organizations": organizations,
		"Jobs":          jobs,
		"IsEdit":        isEdit,
	}

	if user != nil {
		data["User"] = user
	} else {
		data["FormData"] = gin.H{
			"Name":           c.PostForm("name"),
			"OrganizationID": c.PostForm("organizationID"),
			"JobID":          c.PostForm("jobID"),
		}
	}

	c.HTML(http.StatusBadRequest, "users/form.html", data)
}

func (h UserHandler) Index(c *gin.Context) {
	id, ok := h.parseUserID(c)
	if !ok {
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		renderError(c, "users/list.html", "User not found", http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"Title": "Jobs - " + user.Name,
		"User":  user,
	})
}
