package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
        "sync"
        "time"

        "gitlab.com/h2c-bd2c/transactions-virements/internal/virement" 
)

// Stockage en mémoire des virements (thread-safe)
var (
        virements  = make(map[string]*virement.Virement)
        virementsM sync.Mutex
)

func main() {
        http.HandleFunc("/api/v1/virements", createVirementHandler)
        http.HandleFunc("/api/v1/virements/", getVirementByIDHandler)

        fmt.Println("Virements service started on port 8080")
        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatalf("Server failed: %v", err)
        }
}

// createVirementHandler crée un virement depuis une requête POST JSON
func createVirementHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
                http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
                return
        }

        var req struct {
                FromAccountID string  `json:"from"`
                ToAccountID   string  `json:"to"`
                Amount        float64 `json:"amount"`
                Currency      string  `json:"currency"`
                Reference     string  `json:"reference"`
        }

        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                http.Error(w, "invalid request body", http.StatusBadRequest)
                return
        }

        v, err := virement.NewVirement(req.FromAccountID, req.ToAccountID, req.Amount, req.Currency, req.Reference)
        if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        // Générer un ID simple basé sur timestamp
        v.ID = fmt.Sprintf("%d", time.Now().UnixNano())

        // Stocker le virement
        virementsM.Lock()
        virements[v.ID] = v
        virementsM.Unlock()

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(v)

        fmt.Printf("Virement created: %+v\n", v)
}

// listVirementsHandler retourne tous les virements
func listVirementsHandler(w http.ResponseWriter, r *http.Request) {
        virementsM.Lock()
        defer virementsM.Unlock()

        list := make([]*virement.Virement, 0, len(virements))
        for _, v := range virements {
                list = append(list, v)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(list)
}

func getVirementByIDHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
                http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
                return
        }

        // Extraire l'ID depuis l'URL
        id := r.URL.Path[len("/api/v1/virements/"):]

        if id == "" {
                http.Error(w, "virement id is required", http.StatusBadRequest)
                return
        }

        virementsM.Lock()
        v, exists := virements[id]
        virementsM.Unlock()

        if !exists {
                http.Error(w, "virement not found", http.StatusNotFound)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(v)
}
