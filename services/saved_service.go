package services

import (
	"errors"

	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/repositories"
)

type SavedService interface {
	SaveJob(freelancerID, jobID uint) error
	GetSavedJobs(freelancerID uint) ([]dto.SavedJobResponse, error)
	RemoveSavedJob(freelancerID, jobID uint) error

	SaveFreelancer(companyID, freelancerID uint) error
	GetSavedFreelancers(companyID uint) ([]dto.SavedFreelancerResponse, error)
	RemoveSavedFreelancer(companyID, freelancerID uint) error
}

type savedService struct {
	savedRepo repositories.SavedRepository
	jobRepo   repositories.JobRepository
	userRepo  repositories.UserRepository
}

func NewSavedService(savedRepo repositories.SavedRepository, jobRepo repositories.JobRepository, userRepo repositories.UserRepository) SavedService {
	return &savedService{savedRepo, jobRepo, userRepo}
}

// ✅ Simpan pekerjaan ke daftar favorit freelancer
func (s *savedService) SaveJob(freelancerID, jobID uint) error {
	// Cek apakah job ada
	_, err := s.jobRepo.GetJobByID(jobID)
	if err != nil {
		return errors.New("job not found")
	}

	return s.savedRepo.SaveJob(freelancerID, jobID)
}

// ✅ Ambil daftar pekerjaan yang disimpan oleh freelancer
func (s *savedService) GetSavedJobs(freelancerID uint) ([]dto.SavedJobResponse, error) {
	return s.savedRepo.GetSavedJobs(freelancerID)
}

// ✅ Hapus pekerjaan dari daftar favorit freelancer
func (s *savedService) RemoveSavedJob(freelancerID, jobID uint) error {
	return s.savedRepo.RemoveSavedJob(freelancerID, jobID)
}

// ✅ Simpan freelancer ke daftar favorit perusahaan
func (s *savedService) SaveFreelancer(companyID, freelancerID uint) error {
	// Cek apakah freelancer ada
	user, err := s.userRepo.GetUserByID(freelancerID)
	if err != nil {
		return errors.New("freelancer not found")
	}

	// Pastikan user adalah seorang freelancer
	if user.Role != "freelancer" {
		return errors.New("user is not a freelancer")
	}

	return s.savedRepo.SaveFreelancer(companyID, freelancerID)
}

// ✅ Ambil daftar freelancer yang disimpan oleh perusahaan
func (s *savedService) GetSavedFreelancers(companyID uint) ([]dto.SavedFreelancerResponse, error) {
	return s.savedRepo.GetSavedFreelancers(companyID)
}

// ✅ Hapus freelancer dari daftar favorit perusahaan
func (s *savedService) RemoveSavedFreelancer(companyID, freelancerID uint) error {
	return s.savedRepo.RemoveSavedFreelancer(companyID, freelancerID)
}
