package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/config"
	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/logger"
)

func main() {
	cfg := config.Load("8005")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "comments-service OK")
	})

	addr := ":" + cfg.Port
	logger.Info("comments-service starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("comments-service failed: %v", err)
	}
}
