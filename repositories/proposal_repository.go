package repositories

import (
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"gorm.io/gorm"
)

type ProposalRepository interface {
	GetProposalsByCompanyID(companyID uint) ([]dto.ProposalResponse, error)
	CreateProposal(proposal *models.Proposal) error
	GetProposalsByJobID(jobID uint) ([]dto.ProposalResponse, error)
	GetProposalsByFreelancerID(freelancerID uint) ([]dto.ProposalResponse, error)
	UpdateProposalStatus(proposalID uint, status string) error
	DeleteProposal(proposalID uint) error
	GetProposalByID(proposalID uint) (*models.Proposal, error)
}

type proposalRepository struct {
	db *gorm.DB
}

func NewProposalRepository(db *gorm.DB) ProposalRepository {
	return &proposalRepository{db}
}

// ✅ 1. Simpan proposal baru ke database
func (r *proposalRepository) CreateProposal(proposal *models.Proposal) error {
	return r.db.Create(proposal).Error
}

func (r *proposalRepository) GetProposalsByCompanyID(companyID uint) ([]dto.ProposalResponse, error) {
	var proposals []dto.ProposalResponse

	err := r.db.Table("proposals").
		Select("proposals.id, proposals.job_id, jobs.title AS job_title, proposals.freelancer_id, users.full_name AS freelancer, proposals.cover_letter, proposals.bid_amount, jobs.currency, proposals.status, proposals.created_at").
		Joins("JOIN jobs ON jobs.id = proposals.job_id").
		Joins("JOIN users ON users.id = proposals.freelancer_id").
		Where("jobs.company_id = ?", companyID).
		Scan(&proposals).Error

	if err != nil {
		return nil, err
	}

	return proposals, nil
}

// ✅ 2. Ambil semua proposal berdasarkan Job ID (Hanya Perusahaan yang Bisa Melihat)
func (r *proposalRepository) GetProposalsByJobID(jobID uint) ([]dto.ProposalResponse, error) {
	var proposals []dto.ProposalResponse

	err := r.db.Table("proposals").
		Select("proposals.id, proposals.job_id, jobs.title AS job_title, proposals.freelancer_id, users.full_name AS freelancer, proposals.cover_letter, proposals.bid_amount, proposals.currency, proposals.status, proposals.created_at").
		Joins("JOIN jobs ON jobs.id = proposals.job_id").
		Joins("JOIN users ON users.id = proposals.freelancer_id").
		Where("proposals.job_id = ?", jobID).
		Order("proposals.created_at DESC"). // Urutkan dari terbaru
		Scan(&proposals).Error

	if err != nil {
		return nil, err
	}

	return proposals, nil
}

// ✅ 3. Ambil proposal berdasarkan Freelancer ID (Hanya Freelancer yang Bisa Melihat)
func (r *proposalRepository) GetProposalsByFreelancerID(freelancerID uint) ([]dto.ProposalResponse, error) {
	var proposals []dto.ProposalResponse

	err := r.db.Table("proposals").
		Select("proposals.id, proposals.job_id, jobs.title AS job_title, proposals.freelancer_id, users.full_name AS freelancer, proposals.cover_letter, proposals.bid_amount, proposals.currency, proposals.status, proposals.created_at").
		Joins("JOIN jobs ON jobs.id = proposals.job_id").
		Joins("JOIN users ON users.id = proposals.freelancer_id").
		Where("proposals.freelancer_id = ?", freelancerID).
		Order("proposals.created_at DESC"). // Urutkan dari terbaru
		Scan(&proposals).Error

	if err != nil {
		return nil, err
	}

	return proposals, nil
}

// ✅ 4. Perbarui status proposal (Hanya Perusahaan yang Bisa)
func (r *proposalRepository) UpdateProposalStatus(proposalID uint, status string) error {
	return r.db.Model(&models.Proposal{}).
		Where("id = ?", proposalID).
		Update("status", status).Error
}

// ✅ 5. Hapus proposal (Hanya Freelancer yang Bisa Menghapus Proposal Mereka)
func (r *proposalRepository) DeleteProposal(proposalID uint) error {
	return r.db.Delete(&models.Proposal{}, proposalID).Error
}

// ✅ Ambil satu proposal berdasarkan proposal_id
func (r *proposalRepository) GetProposalByID(proposalID uint) (*models.Proposal, error) {
	var proposal models.Proposal
	err := r.db.
		Joins("JOIN jobs ON jobs.id = proposals.job_id").
		Where("proposals.id = ?", proposalID).
		Select("proposals.*, jobs.company_id").
		First(&proposal).Error
	if err != nil {
		return nil, err
	}
	return &proposal, nil
}
