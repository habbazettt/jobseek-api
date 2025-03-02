package services

import (
	"errors"

	"github.com/habbazettt/jobseek-go/dto"
	"github.com/habbazettt/jobseek-go/models"
	"github.com/habbazettt/jobseek-go/repositories"
)

type JobService interface {
	CreateJob(request dto.JobRequest, companyID uint) (*dto.JobResponse, error)
	GetJobs(filters dto.JobFilterRequest) (map[string]interface{}, error)
	GetJobByID(id uint) (*dto.JobResponse, error)
	UpdateJob(id uint, request dto.JobRequest, companyID uint) (*dto.JobResponse, error)
	DeleteJob(id uint, companyID uint, userRole string) error
}

type jobService struct {
	jobRepo repositories.JobRepository
}

func NewJobService(jobRepo repositories.JobRepository) JobService {
	return &jobService{jobRepo}
}

// ✅ CreateJob - Tambahkan pekerjaan
func (s *jobService) CreateJob(request dto.JobRequest, companyID uint) (*dto.JobResponse, error) {
	job := models.Job{
		Title:           request.Title,
		Description:     request.Description,
		CompanyID:       companyID,
		Location:        request.Location,
		Salary:          request.Salary,
		Currency:        request.Currency,
		JobType:         request.JobType,
		Category:        request.Category,
		ExperienceLevel: request.ExperienceLevel,
		Skills:          request.Skills, // ✅ GORM akan menyimpan sebagai JSON otomatis
		Deadline:        request.Deadline,
		Status:          "open",
	}

	err := s.jobRepo.CreateJob(&job)
	if err != nil {
		return nil, err
	}

	response := dto.JobResponse{
		ID:              job.ID,
		Title:           job.Title,
		Description:     job.Description,
		CompanyID:       job.CompanyID,
		Location:        job.Location,
		Salary:          job.Salary,
		Currency:        job.Currency,
		JobType:         job.JobType,
		Category:        job.Category,
		ExperienceLevel: job.ExperienceLevel,
		Skills:          job.Skills, // ✅ GORM akan mengembalikan dalam bentuk []string
		Deadline:        job.Deadline,
		Status:          job.Status,
		CreatedAt:       job.CreatedAt,
	}
	return &response, nil
}

// ✅ GetJobs - Ambil semua pekerjaan
func (s *jobService) GetJobs(filters dto.JobFilterRequest) (map[string]interface{}, error) {
	jobs, total, err := s.jobRepo.GetJobs(filters) // ✅ Tambahkan parameter `filters`
	if err != nil {
		return nil, err
	}

	var jobResponses []dto.JobResponse
	for _, job := range jobs {
		jobResponses = append(jobResponses, dto.JobResponse{
			ID:              job.ID,
			Title:           job.Title,
			Description:     job.Description,
			CompanyID:       job.CompanyID,
			Location:        job.Location,
			Salary:          job.Salary,
			Currency:        job.Currency,
			JobType:         job.JobType,
			Category:        job.Category,
			ExperienceLevel: job.ExperienceLevel,
			Skills:          job.Skills,
			Deadline:        job.Deadline,
			Status:          job.Status,
			CreatedAt:       job.CreatedAt,
		})
	}

	// 🔄 Format response dengan pagination
	response := map[string]interface{}{
		"total":   total,
		"page":    filters.Page,
		"limit":   filters.Limit,
		"results": jobResponses,
	}
	return response, nil
}

// ✅ GetJobByID - Ambil pekerjaan berdasarkan ID
func (s *jobService) GetJobByID(id uint) (*dto.JobResponse, error) {
	job, err := s.jobRepo.GetJobByID(id)
	if err != nil {
		return nil, err
	}

	response := dto.JobResponse{
		ID:              job.ID,
		Title:           job.Title,
		Description:     job.Description,
		CompanyID:       job.CompanyID,
		Location:        job.Location,
		Salary:          job.Salary,
		Currency:        job.Currency,
		JobType:         job.JobType,
		Category:        job.Category,
		ExperienceLevel: job.ExperienceLevel,
		Skills:          job.Skills, // ✅ GORM akan mengembalikan []string
		Deadline:        job.Deadline,
		Status:          job.Status,
		CreatedAt:       job.CreatedAt,
	}
	return &response, nil
}

// ✅ UpdateJob - Perusahaan hanya bisa mengupdate pekerjaannya sendiri
func (s *jobService) UpdateJob(id uint, request dto.JobRequest, companyID uint) (*dto.JobResponse, error) {
	job, err := s.jobRepo.GetJobByID(id)
	if err != nil {
		return nil, err
	}

	if job.CompanyID != companyID {
		return nil, errors.New("unauthorized: you can only update your own jobs")
	}

	// Update data pekerjaan
	job.Title = request.Title
	job.Description = request.Description
	job.Location = request.Location
	job.Salary = request.Salary
	job.Currency = request.Currency
	job.JobType = request.JobType
	job.Category = request.Category
	job.ExperienceLevel = request.ExperienceLevel
	job.Skills = request.Skills
	job.Deadline = request.Deadline

	err = s.jobRepo.UpdateJob(job)
	if err != nil {
		return nil, err
	}

	response := dto.JobResponse{
		ID:              job.ID,
		Title:           job.Title,
		Description:     job.Description,
		CompanyID:       job.CompanyID,
		Location:        job.Location,
		Salary:          job.Salary,
		Currency:        job.Currency,
		JobType:         job.JobType,
		Category:        job.Category,
		ExperienceLevel: job.ExperienceLevel,
		Skills:          job.Skills, // ✅ GORM akan mengembalikan []string
		Deadline:        job.Deadline,
		Status:          job.Status,
		CreatedAt:       job.CreatedAt,
	}
	return &response, nil
}

// ✅ DeleteJob - Hanya perusahaan yang membuat atau admin yang bisa menghapus
func (s *jobService) DeleteJob(id uint, companyID uint, userRole string) error {
	job, err := s.jobRepo.GetJobByID(id)
	if err != nil {
		return err
	}

	if job.CompanyID != companyID && userRole != "admin" {
		return errors.New("unauthorized: only the job owner or admin can delete")
	}

	return s.jobRepo.DeleteJob(id)
}
