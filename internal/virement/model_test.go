package virement

import "testing"

func TestNewVirement_StatusIsPending(t *testing.T) {
        v, err := NewVirement("ACC1", "ACC2", 200, "MAD", "REF123")

        if err != nil {
                t.Fatalf("unexpected error: %v", err)
        }

        if v.Status != StatusPending {
                t.Errorf("expected status %s, got %s", StatusPending, v.Status)
        }
}

func TestNewVirement_InvalidAmount(t *testing.T) {
        _, err := NewVirement("ACC1", "ACC2", 0, "MAD", "REF123")

        if err == nil {
                t.Error("expected error for invalid amount")
        }
}

func TestNewVirement_EmptyCurrency(t *testing.T) {
        _, err := NewVirement("ACC1", "ACC2", 100, "", "REF123")

        if err == nil {
                t.Error("expected error for empty currency")
        }
}

func TestNewVirement_DefaultStatusIsPending(t *testing.T) {
        v := Virement{
                FromAccountID: "ACC123",
                ToAccountID:   "ACC456",
                Amount:        100,
                Currency:      "MAD",
        }

        if v.Status != "" {
                t.Errorf("expected empty status, got %s", v.Status)
        }
}

func TestVirementIsValid(t *testing.T) {
        v := Virement{
                FromAccountID: "ACC123",
                ToAccountID:   "ACC999",
                Amount:        100,
                Currency:      "MAD",
        }

        if !v.IsValid() {
                t.Errorf("expected virement to be valid")
        }
}

func TestVirementInvalidAmount(t *testing.T) {
        v := Virement{
                FromAccountID: "A",
                ToAccountID:   "B",
                Amount:        -10,
                Currency:      "MAD",
        }

        if v.IsValid() {
                t.Errorf("expected virement to be invalid due to amount")
        }
}

func TestVirementMissingAccounts(t *testing.T) {
        v := Virement{
                Amount:   50,
                Currency: "MAD",
        }

        if v.IsValid() {
                t.Errorf("expected virement to be invalid due to missing accounts")
        }
}
