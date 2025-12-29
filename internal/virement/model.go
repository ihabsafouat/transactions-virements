// Package virement contient les structures et fonctions pour gérer les virements bancaires.
package virement

import (
        "errors"
        "time"
)

// Status représente l'état d'un virement.
type Status string

const (
        // StatusPending indique qu'un virement est en attente.
        StatusPending Status = "PENDING"

        // StatusProcessing indique qu'un virement est en cours de traitement.
        StatusProcessing Status = "PROCESSING"

        // StatusCompleted indique qu'un virement a été complété avec succès.
        StatusCompleted Status = "COMPLETED"

        // StatusRejected indique qu'un virement a été rejeté.
        StatusRejected Status = "REJECTED"

        // StatusCancelled indique qu'un virement a été annulé.
        StatusCancelled Status = "CANCELLED"

        // StatusFailed indique qu'un virement a échoué.
        StatusFailed Status = "FAILED"
)

// Virement représente un transfert d'argent entre deux comptes.
type Virement struct {
        ID            string
        FromAccountID string
        ToAccountID   string
        Amount        float64
        Currency      string
        Status        Status
        CreatedAt     time.Time
        UpdatedAt     time.Time
        Reference     string
}

// NewVirement crée un nouveau virement avec les informations fournies.
// Elle retourne une erreur si le montant est inférieur ou égal à 0 ou si la devise est vide.
func NewVirement(from, to string, amount float64, currency, reference string) (*Virement, error) {
        if amount <= 0 {
                return nil, errors.New("amount must be greater than 0")
        }

        if currency == "" {
                return nil, errors.New("currency is required")
        }

        now := time.Now()

        return &Virement{
                FromAccountID: from,
                ToAccountID:   to,
                Amount:        amount,
                Currency:      currency,
                Status:        StatusPending,
                CreatedAt:     now,
                UpdatedAt:     now,
                Reference:     reference,
        }, nil
}

// IsValid vérifie si le virement contient toutes les informations nécessaires.
func (v Virement) IsValid() bool {
        if v.Amount <= 0 {
                return false
        }
        if v.FromAccountID == "" || v.ToAccountID == "" {
                return false
        }
        if v.Currency == "" {
                return false
        }
        return true
}
