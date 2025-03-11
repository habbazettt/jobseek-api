package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/utils"
)

type SavedController struct {
	savedService services.SavedService
}

func NewSavedController(savedService services.SavedService) *SavedController {
	return &SavedController{savedService}
}

// @Summary      Save Job
// @Description  Save a job to the freelancer's saved jobs. Only freelancers can perform this action.
// @Tags         saved
// @Accept       json
// @Produce      json
// @Param        job_id  path      int  true  "Job ID"
// @Security     BearerAuth
// @Success      201      {object}  dto.SavedJobResponse "Job saved successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid request body"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Forbidden: Only freelancers can save jobs"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to save job"
// @Router       /saved/jobs/{job_id} [post]
func (c *SavedController) SaveJob(ctx *gin.Context) {
	jobID, err := strconv.Atoi(ctx.Param("job_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid job ID")
		return
	}

	freelancerID, _ := ctx.Get("user_id")
	userRole, _ := ctx.Get("role")

	// Hanya freelancer yang bisa menyimpan pekerjaan
	if userRole != "freelancer" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only freelancers can save jobs")
		return
	}

	err = c.savedService.SaveJob(freelancerID.(uint), uint(jobID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Job saved successfully", nil)
}


// @Summary      Get Saved Jobs
// @Description  Get list of jobs saved by the freelancer. Only freelancers can perform this action.
// @Tags         saved
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  []dto.SavedJobResponse "Jobs retrieved successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid request body"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Forbidden: Only freelancers can view saved jobs"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to retrieve jobs"
// @Router       /saved/jobs [get]
func (c *SavedController) GetSavedJobs(ctx *gin.Context) {
	freelancerID, _ := ctx.Get("user_id")
	userRole, _ := ctx.Get("role")

	// Hanya freelancer yang bisa melihat daftar pekerjaan yang disimpan
	if userRole != "freelancer" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only freelancers can view saved jobs")
		return
	}

	savedJobs, err := c.savedService.GetSavedJobs(freelancerID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Saved jobs retrieved successfully", savedJobs)
}



// @Summary      Remove Saved Job
// @Description  Remove a job from the freelancer's saved jobs. Only freelancers can perform this action.
// @Tags         saved
// @Accept       json
// @Produce      json
// @Param        job_id  path      int  true  "Job ID"
// @Security     BearerAuth
// @Success      200      {object}  dto.SavedJobResponse "Job removed from saved list successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid request body"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Forbidden: Only freelancers can remove saved jobs"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to remove saved job"
// @Router       /saved/jobs/{job_id} [delete]
func (c *SavedController) RemoveSavedJob(ctx *gin.Context) {
	jobID, err := strconv.Atoi(ctx.Param("job_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid job ID")
		return
	}

	freelancerID, _ := ctx.Get("user_id")
	userRole, _ := ctx.Get("role")

	// Hanya freelancer yang bisa menghapus pekerjaan dari daftar favorit
	if userRole != "freelancer" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only freelancers can remove saved jobs")
		return
	}

	err = c.savedService.RemoveSavedJob(freelancerID.(uint), uint(jobID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Saved job removed successfully", nil)
}


// SaveFreelancer godoc
// @Summary      Save Freelancer
// @Description  Save a freelancer to the company's saved list. Only companies can perform this action.
// @Tags         saved
// @Accept       json
// @Produce      json
// @Param        freelancer_id  path      int  true  "Freelancer ID"
// @Security     BearerAuth
// @Success      201      {object}  dto.SavedFreelancerResponse "Freelancer saved successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid freelancer ID"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Forbidden: Only companies can save freelancers"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to save freelancer"
// @Router       /saved/freelancers/{freelancer_id} [post]

func (c *SavedController) SaveFreelancer(ctx *gin.Context) {
	freelancerID, err := strconv.Atoi(ctx.Param("freelancer_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid freelancer ID")
		return
	}

	companyID, _ := ctx.Get("user_id")
	userRole, _ := ctx.Get("role")

	// Hanya perusahaan yang bisa menyimpan freelancer ke daftar favorit
	if userRole != "perusahaan" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only companies can save freelancers")
		return
	}

	err = c.savedService.SaveFreelancer(companyID.(uint), uint(freelancerID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Freelancer saved successfully", nil)
}


// GetSavedFreelancers godoc
// @Summary      Get Saved Freelancers
// @Description  Get list of freelancers saved by the company. Only companies can perform this action.
// @Tags         saved
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  []dto.SavedFreelancerResponse "Freelancers retrieved successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid request body"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Forbidden: Only companies can view saved freelancers"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to retrieve freelancers"
// @Router       /saved/freelancers [get]
func (c *SavedController) GetSavedFreelancers(ctx *gin.Context) {
	companyID, _ := ctx.Get("user_id")
	userRole, _ := ctx.Get("role")

	// Hanya perusahaan yang bisa melihat daftar freelancer yang disimpan
	if userRole != "perusahaan" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only companies can view saved freelancers")
		return
	}

	savedFreelancers, err := c.savedService.GetSavedFreelancers(companyID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Saved freelancers retrieved successfully", savedFreelancers)
}


// RemoveSavedFreelancer godoc
// @Summary      Remove Saved Freelancer
// @Description  Remove a freelancer from the company's saved freelancers. Only companies can perform this action.
// @Tags         saved
// @Accept       json
// @Produce      json
// @Param        freelancer_id  path      int  true  "Freelancer ID"
// @Security     BearerAuth
// @Success      200      {object}  dto.SavedFreelancerResponse "Freelancer removed from saved list successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid request body"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Forbidden: Only companies can remove saved freelancers"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to remove saved freelancer"
// @Router       /saved/freelancers/{freelancer_id} [delete]
func (c *SavedController) RemoveSavedFreelancer(ctx *gin.Context) {
	freelancerID, err := strconv.Atoi(ctx.Param("freelancer_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid freelancer ID")
		return
	}

	companyID, _ := ctx.Get("user_id")
	userRole, _ := ctx.Get("role")

	// Hanya perusahaan yang bisa menghapus freelancer dari daftar favorit
	if userRole != "perusahaan" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only companies can remove saved freelancers")
		return
	}

	err = c.savedService.RemoveSavedFreelancer(companyID.(uint), uint(freelancerID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Saved freelancer removed successfully", nil)
}
