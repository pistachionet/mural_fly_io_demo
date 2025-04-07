package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	region := os.Getenv("FLY_REGION")
	if region == "" {
		region = "unknown"
	}
	fmt.Fprintf(w, "Hello from Navid! Whats craka lacken team. You are served from region: %s\n", region)
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}