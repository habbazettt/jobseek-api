package dto

import "time"

// ✅ Request DTO untuk menyimpan pekerjaan
type SaveJobRequest struct {
	JobID uint `json:"job_id" binding:"required"`
}

// ✅ Response DTO untuk pekerjaan yang disimpan
type SavedJobResponse struct {
	ID        uint      `json:"id"`
	JobID     uint      `json:"job_id"`
	JobTitle  string    `json:"job_title"`
	CreatedAt time.Time `json:"created_at"`
}

// ✅ Request DTO untuk menyimpan freelancer
type SaveFreelancerRequest struct {
	FreelancerID uint `json:"freelancer_id" binding:"required"`
}

// ✅ Response DTO untuk freelancer yang disimpan
type SavedFreelancerResponse struct {
	ID           uint      `json:"id"`
	FreelancerID uint      `json:"freelancer_id"`
	FullName     string    `json:"full_name"`
	CreatedAt    time.Time `json:"created_at"`
}
