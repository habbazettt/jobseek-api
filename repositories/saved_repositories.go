package repositories

import (
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"gorm.io/gorm"
)

type SavedRepository interface {
	SaveJob(freelancerID, jobID uint) error
	GetSavedJobs(freelancerID uint) ([]dto.SavedJobResponse, error)
	RemoveSavedJob(freelancerID, jobID uint) error

	SaveFreelancer(companyID, freelancerID uint) error
	GetSavedFreelancers(companyID uint) ([]dto.SavedFreelancerResponse, error)
	RemoveSavedFreelancer(companyID, freelancerID uint) error
}

type savedRepository struct {
	db *gorm.DB
}

func NewSavedRepository(db *gorm.DB) SavedRepository {
	return &savedRepository{db}
}

// ✅ Simpan pekerjaan ke daftar favorit freelancer
func (r *savedRepository) SaveJob(freelancerID, jobID uint) error {
	savedJob := models.SavedJob{
		FreelancerID: freelancerID,
		JobID:        jobID,
	}
	return r.db.Create(&savedJob).Error
}

// ✅ Ambil daftar pekerjaan yang disimpan oleh freelancer
func (r *savedRepository) GetSavedJobs(freelancerID uint) ([]dto.SavedJobResponse, error) {
	var savedJobs []dto.SavedJobResponse

	err := r.db.Table("saved_jobs").
		Select("saved_jobs.id, saved_jobs.job_id, jobs.title AS job_title, saved_jobs.created_at").
		Joins("JOIN jobs ON jobs.id = saved_jobs.job_id").
		Where("saved_jobs.freelancer_id = ?", freelancerID).
		Scan(&savedJobs).Error

	return savedJobs, err
}

// ✅ Hapus pekerjaan dari daftar favorit freelancer
func (r *savedRepository) RemoveSavedJob(freelancerID, jobID uint) error {
	return r.db.Where("freelancer_id = ? AND job_id = ?", freelancerID, jobID).
		Delete(&models.SavedJob{}).Error
}

// ✅ Simpan freelancer ke daftar favorit perusahaan
func (r *savedRepository) SaveFreelancer(companyID, freelancerID uint) error {
	savedFreelancer := models.SavedFreelancer{
		CompanyID:    companyID,
		FreelancerID: freelancerID,
	}
	return r.db.Create(&savedFreelancer).Error
}

// ✅ Ambil daftar freelancer yang disimpan oleh perusahaan
func (r *savedRepository) GetSavedFreelancers(companyID uint) ([]dto.SavedFreelancerResponse, error) {
	var savedFreelancers []dto.SavedFreelancerResponse

	err := r.db.Table("saved_freelancers").
		Select("saved_freelancers.id, saved_freelancers.freelancer_id, users.full_name, saved_freelancers.created_at").
		Joins("JOIN users ON users.id = saved_freelancers.freelancer_id").
		Where("saved_freelancers.company_id = ?", companyID).
		Scan(&savedFreelancers).Error

	return savedFreelancers, err
}

// ✅ Hapus freelancer dari daftar favorit perusahaan
func (r *savedRepository) RemoveSavedFreelancer(companyID, freelancerID uint) error {
	return r.db.Where("company_id = ? AND freelancer_id = ?", companyID, freelancerID).
		Delete(&models.SavedFreelancer{}).Error
}
