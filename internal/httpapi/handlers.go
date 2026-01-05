package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/h2c-bd2c/transactions-virements/internal/domain"
	"gitlab.com/h2c-bd2c/transactions-virements/internal/repo"
	"gitlab.com/h2c-bd2c/transactions-virements/internal/service"
)

type Handlers struct {
	svc *service.VirementService
}

func NewHandlers(svc *service.VirementService) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h *Handlers) CreateVirement(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateVirementRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"code":    "BAD_JSON",
			"message": "invalid json",
		})
		return
	}

	v, err := h.svc.Create(req)
	if err != nil {
		if err == service.ErrInvalidRequest {
			writeJSON(w, http.StatusBadRequest, map[string]any{
				"code":    "INVALID_REQUEST",
				"message": "invalid request fields",
			})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"code":    "INTERNAL",
			"message": "server error",
		})
		return
	}

	writeJSON(w, http.StatusCreated, v)
}

func (h *Handlers) GetVirement(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	v, err := h.svc.Get(id)
	if err != nil {
		if err == repo.ErrNotFound {
			writeJSON(w, http.StatusNotFound, map[string]any{
				"code":    "NOT_FOUND",
				"message": "virement not found",
			})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"code":    "INTERNAL",
			"message": "server error",
		})
		return
	}

	writeJSON(w, http.StatusOK, v)
}

func (h *Handlers) ListVirements(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.List()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"code":    "INTERNAL",
			"message": "server error",
		})
		return
	}
	writeJSON(w, http.StatusOK, list)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}