package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/config"
	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/logger"
)

func main() {
	cfg := config.Load("8004")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "snapshot-service OK")
	})

	addr := ":" + cfg.Port
	logger.Info("snapshot-service starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("snapshot-service failed: %v", err)
	}
}
