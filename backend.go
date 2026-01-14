package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run backend.go <port>")
		os.Exit(1)
	}

	port := os.Args[1]

	fmt.Println("Starting backend on port", port)

	http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend on port %s\n", port)
	}))
}
