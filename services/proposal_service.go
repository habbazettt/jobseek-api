package services

import (
	"errors"

	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"github.com/habbazettt/jobseek-go/repositories"
)

type ProposalService interface {
	GetProposalsByCompanyID(companyID uint) ([]dto.ProposalResponse, error)
	CreateProposal(request dto.CreateProposalRequest, freelancerID uint) (*dto.ProposalResponse, error)
	GetProposalsByJobID(jobID uint, companyID uint) ([]dto.ProposalResponse, error)
	GetProposalsByFreelancerID(freelancerID uint) ([]dto.ProposalResponse, error)
	UpdateProposalStatus(proposalID uint, status string, companyID uint) (*dto.ProposalResponse, error)
	DeleteProposal(proposalID uint, freelancerID uint) error
}

type proposalService struct {
	proposalRepo repositories.ProposalRepository
	jobRepo      repositories.JobRepository
	userRepo     repositories.UserRepository
}

func NewProposalService(proposalRepo repositories.ProposalRepository, jobRepo repositories.JobRepository, userRepo repositories.UserRepository) ProposalService {
	return &proposalService{proposalRepo, jobRepo, userRepo}
}

func (s *proposalService) GetProposalsByCompanyID(companyID uint) ([]dto.ProposalResponse, error) {
	return s.proposalRepo.GetProposalsByCompanyID(companyID)
}

// ✅ 1. Freelancer mengajukan proposal
func (s *proposalService) CreateProposal(request dto.CreateProposalRequest, freelancerID uint) (*dto.ProposalResponse, error) {
	// Cek apakah job tersedia
	job, err := s.jobRepo.GetJobByID(request.JobID)
	if err != nil {
		return nil, errors.New("job not found")
	}

	// Pastikan freelancer tidak mengajukan proposal untuk job sendiri
	if job.CompanyID == freelancerID {
		return nil, errors.New("you cannot apply for your own job")
	}

	proposal := models.Proposal{
		JobID:        request.JobID,
		FreelancerID: freelancerID,
		CoverLetter:  request.CoverLetter,
		BidAmount:    request.BidAmount,
		Currency:     request.Currency, // ✅ Tambahkan currency
		Status:       "pending",
	}

	err = s.proposalRepo.CreateProposal(&proposal)
	if err != nil {
		return nil, err
	}

	// Ambil data freelancer untuk response
	freelancer, err := s.userRepo.GetUserByID(freelancerID)
	if err != nil {
		return nil, err
	}

	response := dto.ProposalResponse{
		ID:           proposal.ID,
		JobID:        proposal.JobID,
		JobTitle:     job.Title,
		FreelancerID: proposal.FreelancerID,
		Freelancer:   freelancer.FullName,
		CoverLetter:  proposal.CoverLetter,
		BidAmount:    proposal.BidAmount,
		Currency:     proposal.Currency, // ✅ Tambahkan currency
		Status:       proposal.Status,
		CreatedAt:    proposal.CreatedAt,
	}

	return &response, nil
}

// ✅ 2. Perusahaan melihat proposal berdasarkan Job ID
func (s *proposalService) GetProposalsByJobID(jobID uint, companyID uint) ([]dto.ProposalResponse, error) {
	// Cek apakah job ada dan dimiliki oleh perusahaan
	job, err := s.jobRepo.GetJobByID(jobID)
	if err != nil || job.CompanyID != companyID {
		return nil, errors.New("unauthorized: you can only view proposals for your own jobs")
	}

	proposals, err := s.proposalRepo.GetProposalsByJobID(jobID)
	if err != nil {
		return nil, err
	}

	// ✅ Pastikan currency ditampilkan di response
	for i := range proposals {
		proposals[i].Currency = job.Currency
	}

	return proposals, nil
}

// ✅ 3. Freelancer melihat proposal mereka
func (s *proposalService) GetProposalsByFreelancerID(freelancerID uint) ([]dto.ProposalResponse, error) {
	proposals, err := s.proposalRepo.GetProposalsByFreelancerID(freelancerID)
	if err != nil {
		return nil, err
	}

	// ✅ Pastikan currency ditampilkan di response
	for i := range proposals {
		job, _ := s.jobRepo.GetJobByID(proposals[i].JobID)
		proposals[i].Currency = job.Currency
	}

	return proposals, nil
}

// ✅ 4. Perusahaan mengubah status proposal (accept/reject)
func (s *proposalService) UpdateProposalStatus(proposalID uint, status string, companyID uint) (*dto.ProposalResponse, error) {
	// ✅ 1. Cek apakah proposal ada
	proposal, err := s.proposalRepo.GetProposalByID(proposalID)
	if err != nil {
		return nil, errors.New("proposal not found")
	}

	// ✅ 2. Cek apakah job dimiliki oleh perusahaan yang login
	job, err := s.jobRepo.GetJobByID(proposal.JobID)
	if err != nil || job.CompanyID != companyID {
		return nil, errors.New("unauthorized: you can only update proposals for your own jobs")
	}

	// ✅ 3. Update status proposal
	err = s.proposalRepo.UpdateProposalStatus(proposalID, status)
	if err != nil {
		return nil, err
	}

	// ✅ 4. Ambil data freelancer untuk response
	freelancer, err := s.userRepo.GetUserByID(proposal.FreelancerID)
	if err != nil {
		return nil, err
	}

	// ✅ 5. Buat response
	response := &dto.ProposalResponse{
		ID:           proposal.ID,
		JobID:        proposal.JobID,
		JobTitle:     job.Title,
		FreelancerID: proposal.FreelancerID,
		Freelancer:   freelancer.FullName,
		CoverLetter:  proposal.CoverLetter,
		BidAmount:    proposal.BidAmount,
		Currency:     proposal.Currency,
		Status:       status,
		CreatedAt:    proposal.CreatedAt,
	}

	return response, nil
}

// ✅ 5. Freelancer menghapus proposal mereka
func (s *proposalService) DeleteProposal(proposalID uint, freelancerID uint) error {
	// Cek apakah proposal ada
	proposals, err := s.proposalRepo.GetProposalsByFreelancerID(freelancerID)
	if err != nil || len(proposals) == 0 {
		return errors.New("proposal not found")
	}

	// Pastikan freelancer hanya bisa menghapus proposal mereka sendiri
	if proposals[0].FreelancerID != freelancerID {
		return errors.New("unauthorized: you can only delete your own proposals")
	}

	return s.proposalRepo.DeleteProposal(proposalID)
}
