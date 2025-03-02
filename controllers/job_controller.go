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

// ✅ CreateJob - Menambahkan pekerjaan
func (c *JobController) CreateJob(ctx *gin.Context) {
	var request dto.JobRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Ambil user ID dan role dari token JWT
	companyID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userRole, _ := ctx.Get("role")

	// ✅ Cek apakah user adalah perusahaan
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

// ✅ GetJobs - Mengambil semua pekerjaan
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

// ✅ GetJobByID - Mengambil detail pekerjaan berdasarkan ID
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

// ✅ UpdateJob - Mengupdate pekerjaan
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

// ✅ DeleteJob - Menghapus pekerjaan
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
