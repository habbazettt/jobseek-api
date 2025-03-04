package dto

import "time"

// ✅ Request untuk mengajukan proposal
type CreateProposalRequest struct {
	JobID       uint   `json:"job_id" binding:"required"`
	CoverLetter string `json:"cover_letter" binding:"required,min=10"`
	BidAmount   int64  `json:"bid_amount" binding:"required,min=0"`
	Currency    string `json:"currency" binding:"required,oneof=IDR USD EUR"`
}

// ✅ Request untuk memperbarui status proposal (hanya perusahaan yang bisa)
type UpdateProposalStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending accepted rejected"`
}

// ✅ Response untuk menampilkan data proposal
type ProposalResponse struct {
	ID           uint      `json:"id"`
	JobID        uint      `json:"job_id"`
	JobTitle     string    `json:"job_title"`
	FreelancerID uint      `json:"freelancer_id"`
	Freelancer   string    `json:"freelancer"`
	CoverLetter  string    `json:"cover_letter"`
	BidAmount    int64     `json:"bid_amount"`
	Currency     string    `json:"currency"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}
