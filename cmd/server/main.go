package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"gitlab.com/h2c-bd2c/transactions-virements/internal/httpapi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	srv := &http.Server{
		Addr:              addr,
		Handler:           httpapi.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("virements service starting on %s", addr)
	log.Fatal(srv.ListenAndServe())
}
