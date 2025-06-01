package main

import (
	"github-webhook/src"
	"github-webhook/src/config"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", src.Home)
	mux.HandleFunc("/github", src.GitHubWebhook)

	port := config.Port
	if port == "" {
		port = "3000"
	}

	server := &http.Server{
		Addr:              "0.0.0.0:" + port,
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("🚀 Server running at http://0.0.0.0:%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("❌ Server error: %v", err)
	}
}
