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

// CreateProposal godoc
// @Summary      Create Proposal
// @Description  Create a new proposal for a job
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        request  body     dto.CreateProposalRequest  true  "Proposal data"
// @Success      201      {object} dto.ProposalResponse "Proposal submitted successfully"
// @Failure      400      {object} utils.ErrorResponseSwagger "Invalid request body"
// @Failure      401      {object} utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object} utils.ErrorResponseSwagger "Only freelancers can apply for jobs"
// @Failure      500      {object} utils.ErrorResponseSwagger "Failed to submit proposal"
// @Router       /proposals [post]
// @Security     BearerAuth
func (c *ProposalController) CreateProposal(ctx *gin.Context) {
	var request dto.CreateProposalRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

// GetProposalsByJobID godoc
// @Summary      Get Proposals By Job ID
// @Description  Get all proposals for a job
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        job_id  path      int     true  "Job ID"
// @Success      200      {array}   dto.ProposalResponse "Proposals retrieved successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid job ID"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Only companies can view proposals"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to retrieve proposals"
// @Router       /proposals/{job_id} [get]
// @Security     BearerAuth
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

// GetProposalsByFreelancer godoc
// @Summary      Get Proposals By Freelancer
// @Description  Get all proposals submitted by a freelancer
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Success      200      {array}   dto.ProposalResponse "Proposals retrieved successfully"
// @Failure      400      {object}  utils.ErrorResponseSwagger "Invalid user ID"
// @Failure      401      {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403      {object}  utils.ErrorResponseSwagger "Only freelancers can view their proposals"
// @Failure      500      {object}  utils.ErrorResponseSwagger "Failed to retrieve proposals"
// @Router       /proposals/me [get]
// @Security     BearerAuth
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

// UpdateProposalStatus godoc
// @Summary      Update Proposal Status
// @Description  Allows a company to update the status of a proposal to either pending, accepted, or rejected.
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        proposal_id path int true "Proposal ID"
// @Param        request body dto.UpdateProposalStatusRequest true "Proposal status data"
// @Success      200  {object} dto.ProposalResponse "Proposal status updated successfully"
// @Failure      400  {object} utils.ErrorResponseSwagger "Invalid proposal ID or request body"
// @Failure      401  {object} utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403  {object} utils.ErrorResponseSwagger "Only companies can update proposal status"
// @Failure      500  {object} utils.ErrorResponseSwagger "Failed to update proposal status"
// @Router       /proposals/{proposal_id}/status [put]
// @Security     BearerAuth

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

	updatedProposal, err := c.proposalService.UpdateProposalStatus(uint(proposalID), request.Status, companyID.(uint))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Proposal status updated successfully", updatedProposal)
}

// DeleteProposal godoc
// @Summary      Delete Proposal
// @Description  Delete a proposal that you submitted
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        proposal_id  path      int     true  "Proposal ID"
// @Success      200          {object}  dto.ProposalResponse "Proposal deleted successfully"
// @Failure      400          {object}  utils.ErrorResponseSwagger "Invalid proposal ID"
// @Failure      401          {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403          {object}  utils.ErrorResponseSwagger "Only freelancers can delete their proposals"
// @Failure      500          {object}  utils.ErrorResponseSwagger "Failed to delete proposal"
// @Router       /proposals/{proposal_id} [delete]
// @Security     BearerAuth
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

// GetProposalsByCompany godoc
// @Summary      Get Proposals By Company
// @Description  Retrieve all proposals for jobs posted by the authenticated company
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.ProposalResponse "All proposals retrieved successfully"
// @Failure      401  {object}  utils.ErrorResponseSwagger "Unauthorized"
// @Failure      403  {object}  utils.ErrorResponseSwagger "Only companies can view proposals"
// @Failure      500  {object}  utils.ErrorResponseSwagger "Failed to retrieve proposals"
// @Router       /proposals/company [get]
// @Security     BearerAuth

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
