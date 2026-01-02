package repo

import (
	"errors"
	"sync"

	"gitlab.com/h2c-bd2c/transactions-virements/internal/domain"
)

var ErrNotFound = errors.New("not found")

type VirementRepository interface {
	Save(v domain.Virement) error
	FindByID(id string) (domain.Virement, error)
	List() ([]domain.Virement, error)
	Update(v domain.Virement) error
}

type MemoryRepo struct {
	mu   sync.RWMutex
	data map[string]domain.Virement
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{data: make(map[string]domain.Virement)}
}

func (r *MemoryRepo) Save(v domain.Virement) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[v.ID] = v
	return nil
}

func (r *MemoryRepo) FindByID(id string) (domain.Virement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.data[id]
	if !ok {
		return domain.Virement{}, ErrNotFound
	}
	return v, nil
}

func (r *MemoryRepo) List() ([]domain.Virement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]domain.Virement, 0, len(r.data))
	for _, v := range r.data {
		out = append(out, v)
	}
	return out, nil
}

func (r *MemoryRepo) Update(v domain.Virement) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[v.ID]; !ok {
		return ErrNotFound
	}
	r.data[v.ID] = v
	return nil
}