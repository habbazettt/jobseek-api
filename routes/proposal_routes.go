package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/habbazettt/jobseek-go/controllers"
	"github.com/habbazettt/jobseek-go/middleware"
)

func ProposalRoutes(r *gin.Engine, proposalController *controllers.ProposalController) {
	proposals := r.Group("/api/v1/proposals")
	proposals.Use(middleware.AuthMiddleware())
	{
		proposals.POST("/", proposalController.CreateProposal)                         // Freelancer mengajukan proposal
		proposals.GET("/job/:job_id", proposalController.GetProposalsByJobID)          // Perusahaan melihat proposal berdasarkan Job ID
		proposals.GET("/freelancer", proposalController.GetProposalsByFreelancer)      // Freelancer melihat proposal mereka
		proposals.GET("/company", proposalController.GetProposalsByCompany)            // Perusahaan melihat semua proposal yang masuk
		proposals.PUT("/:proposal_id/status", proposalController.UpdateProposalStatus) // Perusahaan update status
		proposals.DELETE("/:proposal_id", proposalController.DeleteProposal)           // Freelancer menghapus proposal mereka
	}
}
