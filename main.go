package main

import (
	"github-webhook/src"
	"github-webhook/src/config"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", src.Home)
	http.HandleFunc("/github", src.GitHubWebhook)

	port := config.Port
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
