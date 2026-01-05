package service

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.com/h2c-bd2c/transactions-virements/internal/domain"
	"gitlab.com/h2c-bd2c/transactions-virements/internal/repo"
)

var ErrInvalidRequest = errors.New("invalid request")

type VirementService struct {
	repo repo.VirementRepository
}

func NewVirementService(r repo.VirementRepository) *VirementService {
	return &VirementService{repo: r}
}

func (s *VirementService) Create(req domain.CreateVirementRequest) (domain.Virement, error) {
	if strings.TrimSpace(req.FromAccountID) == "" ||
		strings.TrimSpace(req.ToAccountID) == "" ||
		req.Amount <= 0 ||
		strings.TrimSpace(req.Currency) == "" {
		return domain.Virement{}, ErrInvalidRequest
	}
	if req.FromAccountID == req.ToAccountID {
		return domain.Virement{}, ErrInvalidRequest
	}

	now := time.Now().UTC()
	v := domain.Virement{
		ID:            uuid.NewString(),
		Reference:     "VIR-" + uuid.NewString()[:8],
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      strings.ToUpper(req.Currency),
		Label:         req.Label,
		Status:        domain.StatusPending,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	if err := s.repo.Save(v); err != nil {
		return domain.Virement{}, err
	}
	return v, nil
}

func (s *VirementService) Get(id string) (domain.Virement, error) {
	return s.repo.FindByID(id)
}

func (s *VirementService) List() ([]domain.Virement, error) {
	return s.repo.List()
}