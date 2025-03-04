package repositories

import (
	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"gorm.io/gorm"
)

type JobRepository interface {
	CreateJob(job *models.Job) error
	GetJobs(filters dto.JobFilterRequest) ([]models.Job, int64, error) // ✅ Perbarui definisi
	GetJobByID(id uint) (*models.Job, error)
	UpdateJob(job *models.Job) error
	DeleteJob(id uint) error
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db}
}

// ✅ Simpan job ke database tanpa perlu manual json.Marshal()
func (r *jobRepository) CreateJob(job *models.Job) error {
	return r.db.Create(job).Error
}

// ✅ Ambil semua jobs tanpa perlu manual json.Unmarshal()
func (r *jobRepository) GetJobs(filters dto.JobFilterRequest) ([]models.Job, int64, error) {
	var jobs []models.Job
	var total int64

	query := r.db.Model(&models.Job{})

	// 🔍 Gunakan LIKE agar pencarian lebih fleksibel
	if filters.SearchQuery != "" {
		query = query.Where("(title LIKE ? OR description LIKE ?)", "%"+filters.SearchQuery+"%", "%"+filters.SearchQuery+"%")
	}
	if filters.Category != "" {
		query = query.Where("category LIKE ?", "%"+filters.Category+"%")
	}
	if filters.Location != "" {
		query = query.Where("location LIKE ?", "%"+filters.Location+"%")
	}
	if filters.ExperienceLevel != "" {
		query = query.Where("experience_level = ?", filters.ExperienceLevel)
	}
	if filters.MinSalary > 0 {
		query = query.Where("salary >= ?", filters.MinSalary)
	}
	if filters.MaxSalary > 0 {
		query = query.Where("salary <= ?", filters.MaxSalary)
	}

	query.Count(&total)

	// 📌 Pagination dengan LIMIT & OFFSET
	offset := (filters.Page - 1) * filters.Limit
	err := query.Limit(filters.Limit).Offset(offset).Find(&jobs).Error

	return jobs, total, err
}

// ✅ Ambil job berdasarkan ID tanpa manual json.Unmarshal()
func (r *jobRepository) GetJobByID(id uint) (*models.Job, error) {
	var job models.Job
	err := r.db.First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// ✅ Update job
func (r *jobRepository) UpdateJob(job *models.Job) error {
	return r.db.Save(job).Error
}

// ✅ Hapus job berdasarkan ID
func (r *jobRepository) DeleteJob(id uint) error {
	return r.db.Model(&models.Job{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}
