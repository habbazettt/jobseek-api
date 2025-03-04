package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/services"
	"github.com/habbazettt/jobseek-go/utils"
)

type ProposalController struct {
	proposalService services.ProposalService
}

func NewProposalController(proposalService services.ProposalService) *ProposalController {
	return &ProposalController{proposalService}
}

// ✅ 1. Freelancer mengajukan proposal
func (c *ProposalController) CreateProposal(ctx *gin.Context) {
	var request dto.CreateProposalRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Ambil freelancer ID dari token JWT
	freelancerID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userRole, _ := ctx.Get("role")
	if userRole != "freelancer" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only freelancers can apply for jobs")
		return
	}

	proposal, err := c.proposalService.CreateProposal(request, freelancerID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Proposal submitted successfully", proposal)
}

// ✅ 2. Perusahaan melihat proposal berdasarkan Job ID
func (c *ProposalController) GetProposalsByJobID(ctx *gin.Context) {
	jobID, err := strconv.Atoi(ctx.Param("job_id"))
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
	if userRole != "perusahaan" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only companies can view proposals")
		return
	}

	proposals, err := c.proposalService.GetProposalsByJobID(uint(jobID), companyID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Proposals retrieved successfully", proposals)
}

// ✅ 3. Freelancer melihat proposal yang mereka ajukan
func (c *ProposalController) GetProposalsByFreelancer(ctx *gin.Context) {
	freelancerID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userRole, _ := ctx.Get("role")
	if userRole != "freelancer" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only freelancers can view their proposals")
		return
	}

	proposals, err := c.proposalService.GetProposalsByFreelancerID(freelancerID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Freelancer proposals retrieved successfully", proposals)
}

// ✅ 4. Perusahaan mengubah status proposal (accept/reject)
func (c *ProposalController) UpdateProposalStatus(ctx *gin.Context) {
	proposalID, err := strconv.Atoi(ctx.Param("proposal_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid proposal ID")
		return
	}

	var request struct {
		Status string `json:"status" binding:"required,oneof=pending accepted rejected"`
	}
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
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only companies can update proposal status")
		return
	}

	// ✅ Perbarui proposal dan dapatkan response yang diperbarui
	updatedProposal, err := c.proposalService.UpdateProposalStatus(uint(proposalID), request.Status, companyID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// ✅ Kirim response dengan data proposal yang telah diperbarui
	utils.SuccessResponse(ctx, http.StatusOK, "Proposal status updated successfully", updatedProposal)
}

// ✅ 5. Freelancer menghapus proposal mereka
func (c *ProposalController) DeleteProposal(ctx *gin.Context) {
	proposalID, err := strconv.Atoi(ctx.Param("proposal_id"))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid proposal ID")
		return
	}

	freelancerID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userRole, _ := ctx.Get("role")
	if userRole != "freelancer" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only freelancers can delete their proposals")
		return
	}

	err = c.proposalService.DeleteProposal(uint(proposalID), freelancerID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Proposal deleted successfully", nil)
}

// ✅ 6. Perusahaan melihat semua proposal yang diterima
func (c *ProposalController) GetProposalsByCompany(ctx *gin.Context) {
	companyID, exists := ctx.Get("user_id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userRole, _ := ctx.Get("role")
	if userRole != "perusahaan" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Only companies can view proposals")
		return
	}

	proposals, err := c.proposalService.GetProposalsByCompanyID(companyID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "All proposals retrieved successfully", proposals)
}
