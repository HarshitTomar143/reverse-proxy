package main

import (
	"log"
	"net/http"
	"time"

	"github.com/harshit/load-balancer/internal/balancer"
    "github.com/harshit/load-balancer/internal/proxy"
)

func main() {
	backends := []string{
		"http://localhost:9001",
		"http://localhost:9002",
		"http://localhost:9003",
	}

	pool, err := balancer.NewBackendPool(backends)
	if err != nil {
		log.Fatal(err)
	}

	handler := proxy.NewReverseProxy(pool)

	balancer.StartHealthCheck(pool, 5*time.Second)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Load-balancing proxy running on :8080")
	log.Fatal(server.ListenAndServe())
}
