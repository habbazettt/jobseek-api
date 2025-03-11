package dto

import "time"

type SaveJobRequest struct {
	JobID uint `json:"job_id" binding:"required"`
}

type SavedJobResponse struct {
	ID        uint      `json:"id"`
	JobID     uint      `json:"job_id"`
	JobTitle  string    `json:"job_title"`
	CreatedAt time.Time `json:"created_at"`
}

type SaveFreelancerRequest struct {
	FreelancerID uint `json:"freelancer_id" binding:"required"`
}

type SavedFreelancerResponse struct {
	ID           uint      `json:"id"`
	FreelancerID uint      `json:"freelancer_id"`
	FullName     string    `json:"full_name"`
	CreatedAt    time.Time `json:"created_at"`
}
