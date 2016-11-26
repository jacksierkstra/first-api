package main

import (
	"first-api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/api/token", handlers.TokenHandler)
	http.ListenAndServe(":8080", nil)
}
