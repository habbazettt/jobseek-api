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

// ✅ Simpan pekerjaan ke daftar favorit freelancer
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

// ✅ Ambil daftar pekerjaan yang disimpan oleh freelancer
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

// ✅ Hapus pekerjaan dari daftar favorit freelancer
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

// ✅ Simpan freelancer ke daftar favorit perusahaan
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

// ✅ Ambil daftar freelancer yang disimpan oleh perusahaan
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

// ✅ Hapus freelancer dari daftar favorit perusahaan
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
