package main

import (
	"log"
	"net/http"
	"os"

	"github.com/a3ylf/bloggregator/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("GET /v1/err", handlers.HandleErr)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("Serving on port: 8080\n")
	log.Fatal(server.ListenAndServe())

}
