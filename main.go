package main

import (
	"log"
	"net/http"

	"backend/internal/app"
)

func main() {
    handler := app.NewHandler()
    http.Handle("/", handler)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}