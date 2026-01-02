package domain

import "time"

type VirementStatus string

const (
	StatusPending    VirementStatus = "PENDING"
	StatusProcessing VirementStatus = "PROCESSING"
	StatusCompleted  VirementStatus = "COMPLETED"
	StatusRejected   VirementStatus = "REJECTED"
	StatusCancelled  VirementStatus = "CANCELLED"
	StatusFailed     VirementStatus = "FAILED"
)

type Virement struct {
	ID            string         `json:"id"`
	Reference     string         `json:"reference"`
	FromAccountID string         `json:"fromAccountId"`
	ToAccountID   string         `json:"toAccountId"`
	Amount        float64        `json:"amount"`
	Currency      string         `json:"currency"`
	Label         string         `json:"label"`
	Status        VirementStatus `json:"status"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

type CreateVirementRequest struct {
	FromAccountID string  `json:"fromAccountId"`
	ToAccountID   string  `json:"toAccountId"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Label         string  `json:"label"`
}