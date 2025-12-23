package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/config"
	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/logger"
)

func main() {
	cfg := config.Load("8000")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "gateway-service OK")
	})

	addr := ":" + cfg.Port
	logger.Info("gateway-service starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("gateway-service failed: %v", err)
	}
}
