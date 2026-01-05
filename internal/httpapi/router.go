package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/h2c-bd2c/transactions-virements/internal/service"
)

func NewRouter(svc *service.VirementService) http.Handler {
	r := chi.NewRouter()

	h := NewHandlers(svc)

	r.Get("/api/v1/health", h.Health)

	r.Route("/api/v1/virements", func(r chi.Router) {
		r.Post("/", h.CreateVirement)
		r.Get("/", h.ListVirements)
		r.Get("/{id}", h.GetVirement)
	})

	return r
}