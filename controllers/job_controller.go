package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/utils"
)

type JobController struct {
	jobService services.JobService
}

func NewJobController(jobService services.JobService) *JobController {
	return &JobController{jobService}
}

// @Summary Create a new job
// @Description Create a new job
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   body  body      dto.JobRequest  true  "Job request"
// @Success 201 {object} dto.JobResponse "Job created successfully"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid request body"
// @Failure 401 {object} utils.ErrorResponseSwagger "Unauthorized"
// @Failure 403 {object} utils.ErrorResponseSwagger "Only companies can create jobs"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to create job"
// @Router /jobs [post]
func (c *JobController) CreateJob(ctx *gin.Context) {
	var request dto.JobRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	companyID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userRole, _ := ctx.Get("role")

	if userRole != "perusahaan" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only companies can create jobs")
		return
	}

	job, err := c.jobService.CreateJob(request, companyID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Job created successfully", job)
}

// @Summary Get list of jobs
// @Description Get list of jobs with pagination and filtering
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   page     query     int     false "Page number"
// @Param   limit     query     int     false "Limit per page"
// @Param   search_query     query     string     false "Search query"
// @Param   category     query     string     false "Job category"
// @Param   location     query     string     false "Job location"
// @Param   experience_level     query     string     false "Job experience level"
// @Param   min_salary     query     int     false "Minimum salary"
// @Param   max_salary     query     int     false "Maximum salary"
// @Security BearerAuth
// @Success 200 {object} dto.JobResponse "Jobs retrieved successfully"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid query parameters"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to retrieve jobs"
// @Router /jobs [get]
func (c *JobController) GetJobs(ctx *gin.Context) {
	var filters dto.JobFilterRequest
	if err := ctx.ShouldBindQuery(&filters); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid query parameters")
		return
	}

	jobs, err := c.jobService.GetJobs(filters)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Jobs retrieved successfully", jobs)
}

// @Summary Get Job By ID
// @Description Get Job By ID
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   id   path      int  true  "Job ID"
// @Security BearerAuth
// @Success 200 {object} dto.JobResponse "Job retrieved successfully"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid job ID"
// @Failure 404 {object} utils.ErrorResponseSwagger "Job not found"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to retrieve job"
// @Router /jobs/{id} [get]
func (c *JobController) GetJobByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid job ID")
		return
	}

	job, err := c.jobService.GetJobByID(uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Job not found")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Job retrieved successfully", job)
}

// @Summary Update Job
// @Description Update job details based on the provided job ID. Only the job owner or an admin
// can perform this update. Accepts a JSON payload for job details.
//
// @Tags jobs
// @Accept  json
// @Produce  json
// @Param   id   path      int  true  "Job ID"
// @Param   body  body      dto.UpdateJobRequest  true  "Job request"
// @Security BearerAuth
// @Success 200 {object} dto.JobResponse "Job updated successfully"
// @Failure 400 {object} utils.ErrorResponseSwagger "Invalid request body"
// @Failure 401 {object} utils.ErrorResponseSwagger "Unauthorized"
// @Failure 403 {object} utils.ErrorResponseSwagger "Only companies can update jobs"
// @Failure 500 {object} utils.ErrorResponseSwagger "Failed to update job"
// @Router /jobs/{id} [put]
func (c *JobController) UpdateJob(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid job ID")
		return
	}

	var request dto.UpdateJobRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	companyID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	job, err := c.jobService.UpdateJob(uint(id), request, companyID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Job updated successfully", job)
}

// @Summary      Delete Job
// @Description  Delete a job based on the provided job ID. Only the job owner or an admin
// can perform this deletion. The function verifies the user identity and role
// from the token to ensure authorization.
//
// @Tags         jobs
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Job ID"
// @Security     BearerAuth
// @Success      200  {object}  dto.JobResponse "Job deleted successfully"
// @Failure      400  {object}  utils.ErrorResponseSwagger "Invalid job ID"
// @Failure      403  {object}  utils.ErrorResponseSwagger "Forbidden: Only companies can delete jobs"
// @Failure      500  {object}  utils.ErrorResponseSwagger "Failed to delete job"
// @Router       /jobs/{id} [delete]
func (c *JobController) DeleteJob(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid job ID")
		return
	}

	companyID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userRole, _ := ctx.Get("role")

	err = c.jobService.DeleteJob(uint(id), companyID.(uint), userRole.(string))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Job deleted successfully", nil)
}
