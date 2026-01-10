package main

import (
	"log"
	"net/http"
	"time"

	"github.com/harshit/load-balancer/internal/proxy"
)

func main() {
	backendURL := "http://localhost:9001"

	handler := proxy.NewReverseProxy(backendURL)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Reverse proxy listening on :8080")
	log.Fatal(server.ListenAndServe())
}
