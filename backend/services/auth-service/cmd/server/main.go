package main

import (
	"log"
	"net/http"

	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/config"
	"github.com/MahijithMenon/Real-Time-Collaboration-System/backend/pkg/logger"
	"github.com/go-chi/chi/v5"

	transport "github.com/MahijithMenon/Real-Time-Collaboration-System/backend/services/auth-service/internal/transport/http"

	db "github.com/MahijithMenon/Real-Time-Collaboration-System/backend/services/auth-service/internal/db"
)

func main() {
	cfg := config.Load("8001")

	dbConn := db.Connect(cfg.DBUrl)
	defer dbConn.Close()

	r := chi.NewRouter()

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth-service OK"))
	})

	addr := ":" + cfg.Port
	logger.Info("auth-service running on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}

	r.Post("/auth/magic/request", transport.MagicLinkRequestHandler)

}
