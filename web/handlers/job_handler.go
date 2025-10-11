package handlers

import (
	"go-db-demo/internal/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobService domain.JobService
	orgService domain.OrganizationService
}

func NewJobHandler(jobService domain.JobService, orgService domain.OrganizationService) *JobHandler {
	return &JobHandler{
		jobService: jobService,
		orgService: orgService,
	}
}

func (h *JobHandler) parseJobID(c *gin.Context) (int64, bool) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		renderError(c, "jobs/list.html", "Invalid job ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (h *JobHandler) validateJobForm(name, orgIDStr string) (string, int64, string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", 0, "Name is required"
	}
	if orgIDStr == "" {
		return "", 0, "Organization is required"
	}

	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		return "", 0, "Invalid organization ID"
	}
	return name, orgID, ""
}

func (h *JobHandler) List(c *gin.Context) {
	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		renderError(c, "jobs/list.html", "Failed to load jobs", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "jobs/list.html", gin.H{
		"Title": "Jobs",
		"Jobs":  jobs,
	})
}

func (h *JobHandler) Index(c *gin.Context) {
	id, ok := h.parseJobID(c)
	if !ok {
		return
	}

	job, err := h.jobService.GetJob(id)
	if err != nil {
		renderError(c, "jobs/list.html", "Job not found", http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "jobs/index.html", gin.H{
		"Title": "Jobs - " + job.Name,
		"Job":   job,
	})
}

func (h *JobHandler) New(c *gin.Context) {
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		renderError(c, "jobs/list.html", "Failed to load organizations", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "jobs/form.html", gin.H{
		"Title":         "New Job",
		"Organizations": organizations,
		"IsEdit":        false,
	})
}

func (h *JobHandler) Create(c *gin.Context) {
	name, orgID, errMsg := h.validateJobForm(c.PostForm("name"), c.PostForm("organizationID"))
	if errMsg != "" {
		h.renderFormWithError(c, errMsg, nil, false)
		return
	}

	job := &domain.Job{Name: name, OrganizationID: orgID}
	_, err := h.jobService.CreateJob(job)
	if err != nil {
		h.renderFormWithError(c, "Failed to create job", nil, false)
		return
	}

	c.Redirect(http.StatusSeeOther, "/jobs")
}

func (h *JobHandler) Edit(c *gin.Context) {
	id, ok := h.parseJobID(c)
	if !ok {
		return
	}

	job, err := h.jobService.GetJob(id)
	if err != nil {
		renderError(c, "jobs/list.html", "Job not found", http.StatusNotFound)
		return
	}

	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		renderError(c, "jobs/list.html", "Failed to load organizations", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "jobs/form.html", gin.H{
		"Title":         "Edit Job",
		"Job":           job,
		"Organizations": organizations,
		"IsEdit":        true,
	})
}

func (h *JobHandler) Update(c *gin.Context) {
	id, ok := h.parseJobID(c)
	if !ok {
		return
	}

	currentJob, err := h.jobService.GetJob(id)
	if err != nil {
		renderError(c, "jobs/list.html", "Job not found", http.StatusNotFound)
		return
	}

	name, orgID, errMsg := h.validateJobForm(c.PostForm("name"), c.PostForm("organizationID"))
	if errMsg != "" {
		h.renderFormWithError(c, errMsg, currentJob, true)
		return
	}

	job := &domain.Job{ID: id, Name: name, OrganizationID: orgID}
	_, err = h.jobService.UpdateJob(job)
	if err != nil {
		h.renderFormWithError(c, "Failed to update job", currentJob, true)
		return
	}

	c.Redirect(http.StatusSeeOther, "/jobs")
}

func (h *JobHandler) Delete(c *gin.Context) {
	id, ok := h.parseJobID(c)
	if !ok {
		return
	}

	rowsAffected, err := h.jobService.DeleteJob(id)

	if err != nil {
		h.renderDeleteError(c, "Failed to delete job: "+err.Error())
		return
	}

	if rowsAffected < 1 {
		h.renderDeleteError(c, "Job not found or already deleted")
		return
	}

	c.Redirect(http.StatusSeeOther, "/jobs")
}

func (h *JobHandler) renderDeleteError(c *gin.Context, message string) {
	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		jobs = []domain.Job{}
	}

	renderError(c, "jobs/list.html", message, http.StatusInternalServerError, gin.H{
		"Jobs": jobs,
	})
}

func (h *JobHandler) renderFormWithError(c *gin.Context, errorMessage string, job *domain.Job, isEdit bool) {
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		organizations = []domain.Organization{}
	}

	title := "New Job"
	if isEdit {
		title = "Edit Job"
	}

	data := gin.H{
		"Title":         title,
		"Error":         errorMessage,
		"Organizations": organizations,
		"IsEdit":        isEdit,
	}

	if job != nil {
		data["Job"] = job
	} else {
		data["FormData"] = gin.H{
			"Name":           c.PostForm("name"),
			"OrganizationID": c.PostForm("organizationID"),
		}
	}

	c.HTML(http.StatusBadRequest, "jobs/form.html", data)
}
