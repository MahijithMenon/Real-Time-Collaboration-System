package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/config"
	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/logger"
)

func main() {
	cfg := config.Load("8002")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "collab-service OK")
	})

	addr := ":" + cfg.Port
	logger.Info("collab-service starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("collab-service failed: %v", err)
	}
}
